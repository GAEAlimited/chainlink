package web

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/forwarders"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger/audit"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/core/utils/stringutils"
	"github.com/smartcontractkit/chainlink/v2/core/web/presenters"

	"github.com/gin-gonic/gin"
)

// EVMForwardersController manages EVM forwarders.
type EVMForwardersController struct {
	App chainlink.Application
}

// Index lists EVM forwarders.
func (cc *EVMForwardersController) Index(c *gin.Context, size, page, offset int) {
	orm := forwarders.NewORM(cc.App.GetSqlxDB(), cc.App.GetLogger(), cc.App.GetConfig().Database())
	fwds, count, err := orm.FindForwarders(0, size)

	if err != nil {
		jsonAPIError(c, http.StatusBadRequest, err)
		return
	}

	var resources []presenters.EVMForwarderResource
	for _, fwd := range fwds {
		resources = append(resources, presenters.NewEVMForwarderResource(fwd))
	}

	paginatedResponse(c, "forwarder", size, page, resources, count, err)
}

// TrackEVMForwarderRequest is a JSONAPI request for creating an EVM forwarder.
type TrackEVMForwarderRequest struct {
	EVMChainID *utils.Big     `json:"evmChainId"`
	Address    common.Address `json:"address"`
}

// Track adds a new EVM forwarder.
func (cc *EVMForwardersController) Track(c *gin.Context) {
	var fwd forwarders.Forwarder
	request := &TrackEVMForwarderRequest{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, err)
		return
	}

	db, lggr, cfg := cc.App.GetSqlxDB(), cc.App.GetLogger(), cc.App.GetConfig().Database()
	orm := forwarders.NewORM(db, lggr, cfg)
	q := pg.NewQ(db, lggr, cfg)
	err = q.Transaction(func(tx pg.Queryer) error {
		fwd, err = orm.CreateForwarder(request.Address, *request.EVMChainID, pg.WithQueryer(tx))
		if err != nil {
			return err
		}
		chain, err2 := cc.App.GetRelayers().LegacyEVMChains().Get(request.EVMChainID.String())
		if err2 != nil {
			return err2
		}

		if chain.LogPoller() == logpoller.LogPollerDisabled {
			return errors.New("Log poller must be enabled to add or use forwarders")
		}
		return chain.LogPoller().RegisterFilter(forwarders.NewLogFilter(fwd.Address), pg.WithQueryer(tx))
	})

	if err != nil {
		jsonAPIError(c, http.StatusBadRequest, err)
		return
	}

	cc.App.GetAuditLogger().Audit(audit.ForwarderCreated, map[string]interface{}{
		"forwarderID":         fwd.ID,
		"forwarderAddress":    fwd.Address,
		"forwarderEVMChainID": fwd.EVMChainID,
	})
	jsonAPIResponseWithStatus(c, presenters.NewEVMForwarderResource(fwd), "forwarder", http.StatusCreated)
}

// Delete removes an EVM Forwarder.
func (cc *EVMForwardersController) Delete(c *gin.Context) {
	id, err := stringutils.ToInt64(c.Param("fwdID"))
	if err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, err)
		return
	}

	filterCleanup := func(tx pg.Queryer, evmChainID int64, addr common.Address) error {
		chain, err2 := cc.App.GetRelayers().LegacyEVMChains().Get(big.NewInt(evmChainID).String())
		if err2 != nil {
			// If the chain id doesn't even exist, or logpoller is disabled, then there isn't any filter to clean up.  Returning an error
			// here could be dangerous as it would make it impossible to delete a forwarder with an invalid chain id
			return nil
		}

		if chain.LogPoller() == logpoller.LogPollerDisabled {
			// handle same as non-existent chain id
			return nil
		}

		filter := forwarders.NewLogFilter(addr)
		return chain.LogPoller().UnregisterFilter(filter.Name, pg.WithQueryer(tx))
	}

	orm := forwarders.NewORM(cc.App.GetSqlxDB(), cc.App.GetLogger(), cc.App.GetConfig().Database())
	err = orm.DeleteForwarder(id, filterCleanup)

	if err != nil {
		jsonAPIError(c, http.StatusInternalServerError, err)
		return
	}

	cc.App.GetAuditLogger().Audit(audit.ForwarderDeleted, map[string]interface{}{"id": id})
	jsonAPIResponseWithStatus(c, nil, "forwarder", http.StatusNoContent)
}
