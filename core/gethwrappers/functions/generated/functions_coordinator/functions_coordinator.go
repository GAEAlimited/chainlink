// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package functions_coordinator

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type FunctionsBillingCommitment struct {
	SubscriptionId          uint64
	Client                  common.Address
	CallbackGasLimit        uint32
	Don                     common.Address
	AdminFee                *big.Int
	EstimatedTotalCostJuels *big.Int
	DonFee                  *big.Int
	Timestamp               uint32
	GasOverhead             *big.Int
	ExpectedGasPrice        *big.Int
}

type IFunctionsBillingRequestBilling struct {
	SubscriptionId   uint64
	Client           common.Address
	CallbackGasLimit uint32
	ExpectedGasPrice *big.Int
}

type IFunctionsCoordinatorRequest struct {
	RequestingContract common.Address
	SubscriptionOwner  common.Address
	Data               []byte
	SubscriptionId     uint64
	DataVersion        uint16
	Flags              [32]byte
	CallbackGasLimit   uint32
}

var FunctionsCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"linkToNativeFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyRequestData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTransmittersSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByRouterOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustBeSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedPublicKeyChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedRequestDataVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"signerPayment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"transmitterPayment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"totalCost\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"enumFulfillResult\",\"name\":\"result\",\"type\":\"uint8\"}],\"name\":\"BillingEnd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"don\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"adminFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"estimatedTotalCostJuels\",\"type\":\"uint96\"},{\"internalType\":\"uint80\",\"name\":\"donFee\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"gasOverhead\",\"type\":\"uint40\"},{\"internalType\":\"uint256\",\"name\":\"expectedGasPrice\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structFunctionsBilling.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"BillingStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxCallbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"feedStalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasOverheadBeforeCallback\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasOverheadAfterCallback\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackNativePerUnitLink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"donFee\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"maxSupportedRequestDataVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fulfillmentGasPriceOverEstimationBP\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"CostExceedsCommitment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"InsufficientGasProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"InsufficientSubscriptionBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"InvalidRequestID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestInitiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptionOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dataVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"flags\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"callbackGasLimit\",\"type\":\"uint64\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"OracleResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"RequestTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"deleteCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"deleteNodePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateCost\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"expectedGasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIFunctionsBilling.RequestBilling\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getAdminFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNodePublicKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maxCallbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feedStalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasOverheadAfterCallback\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"fallbackNativePerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"gasOverheadBeforeCallback\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"linkPriceFeed\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxSupportedRequestDataVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"fulfillmentGasPriceOverEstimationBP\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"config\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"expectedGasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIFunctionsBilling.RequestBilling\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getDONFee\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDONPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedData\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThresholdPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subscriptionOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"dataVersion\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"flags\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structIFunctionsCoordinator.Request\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"sendRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"estimatedCost\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTimeoutSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"donPublicKey\",\"type\":\"bytes\"}],\"name\":\"setDONPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"setNodePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"thresholdPublicKey\",\"type\":\"bytes\"}],\"name\":\"setThresholdPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"name\":\"updateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162005ca638038062005ca683398101604081905262000034916200049c565b828282828260013380600081620000925760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c557620000c58162000158565b50505015156080526001600160a01b038216620000f557604051632530e88560e11b815260040160405180910390fd5b600880546001600160a01b0319166001600160a01b0384161790556200011b8162000203565b5050600e80546001600160a01b039092166c01000000000000000000000000026001600160601b03909216919091179055506200066e9350505050565b336001600160a01b03821603620001b25760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000089565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6008546001600160a01b031633146200022f5760405163c41a5b0960e01b815260040160405180910390fd5b6200023a8162000247565b8051602090910120600755565b6000806000806000806000806000898060200190518101906200026b9190620005ac565b98509850985098509850985098509850985060008513620002a3576040516321ea67b360e11b81526004810186905260240162000089565b604080516101208101825263ffffffff808c168083528b8216602084018190528a83168486018190528c841660608601819052938a16608086018190526001600160501b038a1660a0870181905261ffff8a1660c0880181905260e088018a90526101009097018d9052600a8054600160f01b9098026001600160f01b03600160a01b909302600160a01b600160f01b0319600160801b90950294909416600160801b600160f01b03196c0100000000000000000000000090990263ffffffff60601b196801000000000000000090970296909616600160401b600160801b03196401000000009098026001600160401b0319909b169098179990991795909516959095179290921794909416949094179290921792909216179055600b829055600c869055517f9136fb3a39fbe82d27ef8048354b8fcc93cf828f72e5c877d8135fde8dc7fd739062000455908b908b908a908c908b908a908a908a9063ffffffff98891681529688166020880152948716604087015292909516606085015260808401526001600160501b039390931660a083015261ffff9290921660c082015260e08101919091526101000190565b60405180910390a150505050505050505050565b80516001600160a01b03811681146200048157600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600080600060608486031215620004b257600080fd5b620004bd8462000469565b602085810151919450906001600160401b0380821115620004dd57600080fd5b818701915087601f830112620004f257600080fd5b81518181111562000507576200050762000486565b604051601f8201601f19908116603f0116810190838211818310171562000532576200053262000486565b816040528281528a868487010111156200054b57600080fd5b600093505b828410156200056f578484018601518185018701529285019262000550565b60008684830101528097505050505050506200058e6040850162000469565b90509250925092565b805163ffffffff811681146200048157600080fd5b60008060008060008060008060006101208a8c031215620005cc57600080fd5b620005d78a62000597565b9850620005e760208b0162000597565b9750620005f760408b0162000597565b96506200060760608b0162000597565b955060808a015194506200061e60a08b0162000597565b60c08b01519094506001600160501b03811681146200063c57600080fd5b60e08b015190935061ffff811681146200065557600080fd5b809250506101008a015190509295985092959850929598565b60805161561c6200068a6000396000611555015261561c6000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80638cc6acce116100ee578063b1dc65a411610097578063d328a91e11610071578063d328a91e1461053f578063de9dfa4614610547578063e3d0e7121461055a578063f2fde38b1461056d57600080fd5b8063b1dc65a41461040d578063c3f909d414610420578063d227d2451461050f57600080fd5b8063aecc12c5116100c8578063aecc12c5146103a5578063af1296d3146103e5578063afcb95d7146103ed57600080fd5b80638cc6acce146103585780638da5cb5b1461036b5780639883c10d1461039357600080fd5b806379ba50971161015b578063814118341161013557806381411834146102ab57806381f1b938146102c057806381ff7048146102c857806385b214cf1461033557600080fd5b806379ba50971461027d5780637f15e16614610285578063807560311461029857600080fd5b8063466557181161018c5780634665571814610226578063533989871461025457806366316d8d1461026a57600080fd5b8063083a5466146101b3578063181f5a77146101c857806326ceabac14610213575b600080fd5b6101c66101c13660046140bc565b610580565b005b60408051808201909152601c81527f46756e6374696f6e7320436f6f7264696e61746f722076312e302e300000000060208201525b60405161020a9190614162565b60405180910390f35b6101c6610221366004614197565b6105d5565b6102396102343660046142f3565b610677565b60405169ffffffffffffffffffff909116815260200161020a565b61025c6106a5565b60405161020a92919061441a565b6101c66102783660046144c4565b610929565b6101c6610adb565b6101c66102933660046140bc565b610bdd565b6101c66102a63660046144fd565b610c2d565b6102b3610ce4565b60405161020a9190614552565b6101fd610d53565b61031260015460025463ffffffff74010000000000000000000000000000000000000000830481169378010000000000000000000000000000000000000000000000009093041691565b6040805163ffffffff94851681529390921660208401529082015260600161020a565b610348610343366004614565565b610e24565b604051901515815260200161020a565b6101c661036636600461457e565b611046565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161020a565b6007545b60405190815260200161020a565b6103b86103b33660046145bb565b6110ad565b604080519485526bffffffffffffffffffffffff909316602085015291830152606082015260800161020a565b6103976112ef565b60408051600181526000602082018190529181019190915260600161020a565b6101c661041b36600461463b565b6113e2565b6104ab600a54600c54600e54600b5463ffffffff8085169564010000000086048216956c010000000000000000000000008082048416969568010000000000000000830490941694930473ffffffffffffffffffffffffffffffffffffffff16927e0100000000000000000000000000000000000000000000000000000000000090910461ffff1691565b6040805163ffffffff998a168152978916602089015287019590955260608601939093529416608084015273ffffffffffffffffffffffffffffffffffffffff90931660a083015261ffff90921660c082015260e08101919091526101000161020a565b61052261051d3660046146f2565b611b13565b6040516bffffffffffffffffffffffff909116815260200161020a565b6101fd611d03565b6105226105553660046142f3565b611d5a565b6101c6610568366004614814565b611df0565b6101c661057b366004614197565b612814565b610588612825565b60008190036105c3576040517f4f42be3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60116105d0828483614982565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633148061061057503373ffffffffffffffffffffffffffffffffffffffff8216145b610646576040517fed6dd19b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116600090815260106020526040812061067491614001565b50565b600a5474010000000000000000000000000000000000000000900469ffffffffffffffffffff165b92915050565b6060806000600680548060200260200160405190810160405280929190818152602001828054801561070d57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116106e2575b505050505090506000815167ffffffffffffffff811115610730576107306141b4565b60405190808252806020026020018201604052801561076357816020015b606081526020019060019003908161074e5790505b50905060005b825181101561091f576010600084838151811061078857610788614a9d565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546107d5906148e1565b9050600003610810576040517f4f42be3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6010600084838151811061082657610826614a9d565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054610873906148e1565b80601f016020809104026020016040519081016040528092919081815260200182805461089f906148e1565b80156108ec5780601f106108c1576101008083540402835291602001916108ec565b820191906000526020600020905b8154815290600101906020018083116108cf57829003601f168201915b505050505082828151811061090357610903614a9d565b60200260200101819052508061091890614afb565b9050610769565b5090939092509050565b6109316128a8565b806bffffffffffffffffffffffff166000036109675750336000908152600d60205260409020546bffffffffffffffffffffffff165b336000908152600d60205260409020546bffffffffffffffffffffffff808316911610156109c1576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000908152600d6020526040812080548392906109ee9084906bffffffffffffffffffffffff16614b33565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055506000610a3d60085473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f66316d8d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526bffffffffffffffffffffffff85166024830152919250908216906366316d8d90604401600060405180830381600087803b158015610abe57600080fd5b505af1158015610ad2573d6000803e3d6000fd5b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610b61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610be5612825565b6000819003610c20576040517f4f42be3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f6105d0828483614982565b60005473ffffffffffffffffffffffffffffffffffffffff16331480610c785750610c5733612a3f565b8015610c7857503373ffffffffffffffffffffffffffffffffffffffff8416145b610cae576040517fed6dd19b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83166000908152601060205260409020610cde828483614982565b50505050565b60606006805480602002602001604051908101604052809291908181526020018280548015610d4957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d1e575b5050505050905090565b606060118054610d62906148e1565b9050600003610d9d576040517f4f42be3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60118054610daa906148e1565b80601f0160208091040260200160405190810160405280929190818152602001828054610dd6906148e1565b8015610d495780601f10610df857610100808354040283529160200191610d49565b820191906000526020600020905b815481529060010190602001808311610e0657509395945050505050565b60085460009073ffffffffffffffffffffffffffffffffffffffff163314610e78576040517fc41a5b0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260096020908152604091829020825161014081018452815467ffffffffffffffff8116825268010000000000000000810473ffffffffffffffffffffffffffffffffffffffff908116948301949094527c0100000000000000000000000000000000000000000000000000000000900463ffffffff90811694820194909452600182015492831660608201819052740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff9081166080830152600283015490811660a08301526c01000000000000000000000000810469ffffffffffffffffffff1660c0830152760100000000000000000000000000000000000000000000810490941660e08201527a01000000000000000000000000000000000000000000000000000090930464ffffffffff1661010084015260030154610120830152610fcc5750600092915050565b600083815260096020526040808220828155600181018390556002810180547fff000000000000000000000000000000000000000000000000000000000000001690556003018290555184917ff1ca1e9147be737b04a2b018a79405f687a97de8dd8a2559bbe62357343af41491a260019150505b919050565b60085473ffffffffffffffffffffffffffffffffffffffff163314611097576040517fc41a5b0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110a081612b28565b8051602090910120600755565b60085460009081908190819073ffffffffffffffffffffffffffffffffffffffff163314611107576040517fc41a5b0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111146040860186614b5f565b905060000361114e576040517ec1cfc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604051806080016040528087606001602081019061116e9190614bc4565b67ffffffffffffffff16815260209081019061118c90890189614197565b73ffffffffffffffffffffffffffffffffffffffff1681526020016111b760e0890160c08a01614be1565b63ffffffff1681523a60209091015290506112236111d86040880188614b5f565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061121d9250505060a0890160808a01614c0e565b83612e23565b9297509095509350915061123a6020870187614197565b73ffffffffffffffffffffffffffffffffffffffff16857fd6bf2929a5458fc22a90821275e4f54e2566c579e4bb3d38184880bf1701762c3261128360808b0160608c01614bc4565b61129360408c0160208d01614197565b6112a060408d018d614b5f565b8d60800160208101906112b39190614c0e565b8e60a001358f60c00160208101906112cb9190614be1565b6040516112df989796959493929190614c2b565b60405180910390a3509193509193565b600a54600e54604080517ffeaf968c0000000000000000000000000000000000000000000000000000000081529051600093640100000000900463ffffffff169283151592859283926c01000000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff169163feaf968c9160048083019260a09291908290030181865afa158015611389573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ad9190614ceb565b509350509250508280156113cf57506113c68142614d3b565b8463ffffffff16105b156113da57600c5491505b509392505050565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161143891849163ffffffff851691908e908e90819084018382808284376000920191909152506134cf92505050565b61146e576040517f0be3632800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260025480825260035460ff80821660208501526101009091041692820192909252908314611543576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610b58565b6115518b8b8b8b8b8b6134d8565b60007f0000000000000000000000000000000000000000000000000000000000000000156115ae5760028260200151836040015161158f9190614d4e565b6115999190614d96565b6115a4906001614d4e565b60ff1690506115c4565b60208201516115be906001614d4e565b60ff1690505b88811461162d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610b58565b888714611696576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610b58565b3360009081526004602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156116d9576116d9614db8565b60028111156116ea576116ea614db8565b905250905060028160200151600281111561170757611707614db8565b14801561174e57506006816000015160ff168154811061172957611729614a9d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6117b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610b58565b50505050506117c161403b565b6000808a8a6040516117d4929190614de7565b6040519081900381206117eb918e90602001614df7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120838301909252600080845290830152915060005b89811015611af557600060018489846020811061185457611854614a9d565b61186191901a601b614d4e565b8e8e8681811061187357611873614a9d565b905060200201358d8d8781811061188c5761188c614a9d565b90506020020135604051600081526020016040526040516118c9949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156118eb573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526004602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561196b5761196b614db8565b600281111561197c5761197c614db8565b905250925060018360200151600281111561199957611999614db8565b14611a00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610b58565b8251600090879060ff16601f8110611a1a57611a1a614a9d565b602002015173ffffffffffffffffffffffffffffffffffffffff1614611a9c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610b58565b8086846000015160ff16601f8110611ab657611ab6614a9d565b73ffffffffffffffffffffffffffffffffffffffff9092166020929092020152611ae1600186614d4e565b94505080611aee90614afb565b9050611835565b505050611b06833383858e8e613586565b5050505050505050505050565b600080611b3560085473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f10fc49c100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8916600482015263ffffffff8616602482015290915073ffffffffffffffffffffffffffffffffffffffff8216906310fc49c19060440160006040518083038186803b158015611bb357600080fd5b505afa158015611bc7573d6000803e3d6000fd5b50505050620f4240831115611c08576040517f8129bbcd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060405180608001604052808967ffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018663ffffffff1681526020018581525090506000611c9688888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250610677915050565b69ffffffffffffffffffff1690506000611ce789898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250611d5a915050565b9050611cf58787848461388d565b9a9950505050505050505050565b6060600f8054611d12906148e1565b9050600003611d4d576040517f4f42be3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f8054610daa906148e1565b6000611d7b60085473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16632a905ccc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611dc5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de99190614e0b565b9392505050565b855185518560ff16601f831115611e63576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610b58565b80600003611ecd576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610b58565b818314611f5b576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610b58565b611f66816003614e28565b8311611fce576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610b58565b611fd6612825565b6040805160c0810182528a8152602081018a905260ff89169181018290526060810188905267ffffffffffffffff8716608082015260a081018690529061201d90886139ef565b600554156121d25760055460009061203790600190614d3b565b905060006005828154811061204e5761204e614a9d565b60009182526020822001546006805473ffffffffffffffffffffffffffffffffffffffff9092169350908490811061208857612088614a9d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff85811684526004909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009081169091559290911680845292208054909116905560058054919250908061210857612108614e3f565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600680548061217157612171614e3f565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190555061201d915050565b60005b81515181101561263957600060046000846000015184815181106121fb576121fb614a9d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561224557612245614db8565b146122ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610b58565b6040805180820190915260ff821681526001602082015282518051600491600091859081106122dd576122dd614a9d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561237e5761237e614db8565b02179055506000915061238e9050565b60046000846020015184815181106123a8576123a8614a9d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156123f2576123f2614db8565b14612459576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610b58565b6040805180820190915260ff82168152602081016002815250600460008460200151848151811061248c5761248c614a9d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561252d5761252d614db8565b02179055505082518051600592508390811061254b5761254b614a9d565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90931692909217909155820151805160069190839081106125c7576125c7614a9d565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061263181614afb565b9150506121d5565b506040810151600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600180547fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff8116780100000000000000000000000000000000000000000000000063ffffffff43811682029290921780855592048116929182916014916126f191849174010000000000000000000000000000000000000000900416614e6e565b92506101000a81548163ffffffff021916908363ffffffff1602179055506127504630600160149054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613a0c565b600281905582518051600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff9093169290920291909117905560015460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612807988b9891977401000000000000000000000000000000000000000090920463ffffffff16969095919491939192614e8b565b60405180910390a1611b06565b61281c612825565b61067481613ab7565b60005473ffffffffffffffffffffffffffffffffffffffff1633146128a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610b58565b565b60006128b2610ce4565b905080516000036128ef576040517f30274b3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e5460009161290e916bffffffffffffffffffffffff16614f11565b905060005b82518160ff1610156129e05781600d6000858460ff168151811061293957612939614a9d565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a90046bffffffffffffffffffffffff166129a19190614f3c565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550806129d990614f61565b9050612913565b5081516129ed9082614f80565b600e8054600090612a0d9084906bffffffffffffffffffffffff16614b33565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505050565b6000806006805480602002602001604051908101604052809291908181526020018280548015612aa557602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311612a7a575b5050505050905060005b8151811015612b1e578373ffffffffffffffffffffffffffffffffffffffff16828281518110612ae157612ae1614a9d565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603612b0e575060019392505050565b612b1781614afb565b9050612aaf565b5060009392505050565b600080600080600080600080600089806020019051810190612b4a9190614fb0565b98509850985098509850985098509850985060008513612b99576040517f43d4cf6600000000000000000000000000000000000000000000000000000000815260048101869052602401610b58565b604080516101208101825263ffffffff808c168083528b8216602084018190528a83168486018190528c841660608601819052938a166080860181905269ffffffffffffffffffff8a1660a0870181905261ffff8a1660c0880181905260e088018a90526101009097018d9052600a80547e010000000000000000000000000000000000000000000000000000000000009098027dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff740100000000000000000000000000000000000000009093027fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff700100000000000000000000000000000000909502949094167fffff0000000000000000000000000000ffffffffffffffffffffffffffffffff6c010000000000000000000000009099027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff68010000000000000000909702969096167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff6401000000009098027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909b169098179990991795909516959095179290921794909416949094179290921792909216179055600b829055600c869055517f9136fb3a39fbe82d27ef8048354b8fcc93cf828f72e5c877d8135fde8dc7fd7390612e0f908b908b908a908c908b908a908a908a9063ffffffff988916815296881660208801529487166040870152929095166060850152608084015269ffffffffffffffffffff9390931660a083015261ffff9290921660c082015260e08101919091526101000190565b60405180910390a150505050505050505050565b600a5460009081908190819061ffff7e0100000000000000000000000000000000000000000000000000000000000090910481169087161115612e92576040517fdada758700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000612e9e8887610677565b90506000612eac8988611d5a565b90506000612ed0886040015189606001518569ffffffffffffffffffff168561388d565b90506000612ef360085473ffffffffffffffffffffffffffffffffffffffff1690565b89516040517fa47c769600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152909150600090819073ffffffffffffffffffffffffffffffffffffffff84169063a47c769690602401600060405180830381865afa158015612f72573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612fb8919081019061505c565b50505060208d01518d516040517f674603d000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff928316600482015267ffffffffffffffff90911660248201529294509092506000919085169063674603d090604401606060405180830381865afa158015613049573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061306d9190615148565b509150506bffffffffffffffffffffffff851661308a8385614b33565b6bffffffffffffffffffffffff1610156130d0576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061314f308e602001518f600001518560016130ed919061519a565b6040805173ffffffffffffffffffffffffffffffffffffffff958616602080830191909152949095168582015267ffffffffffffffff928316606086015291166080808501919091528151808503909101815260a09093019052815191012090565b905060006040518061014001604052808f6000015167ffffffffffffffff1681526020018f6020015173ffffffffffffffffffffffffffffffffffffffff1681526020018f6040015163ffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff168152602001896bffffffffffffffffffffffff168152602001886bffffffffffffffffffffffff1681526020018a69ffffffffffffffffffff1681526020014263ffffffff168152602001600a600001600c9054906101000a900463ffffffff16600a60000160089054906101000a900463ffffffff1661323b9190614e6e565b63ffffffff1664ffffffffff1681526020018f606001518152509050806009600084815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060608201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160010160146101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555060a08201518160020160006101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555060c082015181600201600c6101000a81548169ffffffffffffffffffff021916908369ffffffffffffffffffff16021790555060e08201518160020160166101000a81548163ffffffff021916908363ffffffff16021790555061010082015181600201601a6101000a81548164ffffffffff021916908364ffffffffff1602179055506101208201518160030155905050817fe2cc993cebc49ce369b0174671d90dbae2201decdae368805de999b437336dcd8260405161347f91906151bb565b60405180910390a250600a54909f959e5063ffffffff6c01000000000000000000000000820481169e50700100000000000000000000000000000000909104169b50939950505050505050505050565b60019392505050565b60006134e5826020614e28565b6134f0856020614e28565b6134fc886101446152d0565b61350691906152d0565b61351091906152d0565b61351b9060006152d0565b9050368114610ad2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610b58565b6060808061359684860186615363565b8251929550909350915015806135ae57508151835114155b806135bb57508051835114155b156135f2576040517f0be3632800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b835181101561388157600061365685838151811061361557613615614a9d565b602002602001015185848151811061362f5761362f614a9d565b602002602001015185858151811061364957613649614a9d565b6020026020010151613bac565b9050600081600681111561366c5761366c614db8565b14806136895750600181600681111561368757613687614db8565b145b156136e4578482815181106136a0576136a0614a9d565b60209081029190910181015160405133815290917fc708e0440951fd63499c0f7a73819b469ee5dd3ecc356c0ab4eb7f18389009d9910160405180910390a2613870565b60028160068111156136f8576136f8614db8565b036137485784828151811061370f5761370f614a9d565b60200260200101517fa1c120e327c9ad8b075793878c88d59b8934b97ae37117faa3bb21616237f7be60405160405180910390a2613870565b600381600681111561375c5761375c614db8565b036137ac5784828151811061377357613773614a9d565b60200260200101517f357a60574f143e144d62ddc89636fee45ba5e2f8c9a30b9f91e28f8dfe5a56e060405160405180910390a2613870565b60058160068111156137c0576137c0614db8565b03613810578482815181106137d7576137d7614a9d565b60200260200101517f689f49eacf6db85082542f4b7873f4a519744a4e093fd002310545bfb4354cc860405160405180910390a2613870565b600481600681111561382457613824614db8565b036138705784828151811061383b5761383b614a9d565b60200260200101517fb83d9204606ee3780daa106ad026c4c2853f5af126edbb8fe15e813e62dbd94d60405160405180910390a25b5061387a81614afb565b90506135f5565b50505050505050505050565b6000806138986112ef565b9050600081136138d7576040517f43d4cf6600000000000000000000000000000000000000000000000000000000815260048101829052602401610b58565b600a54600090879061390d9063ffffffff6c01000000000000000000000000820481169168010000000000000000900416614e6e565b6139179190614e6e565b63ffffffff1690506000612710600a60010154886139359190614e28565b61393f9190615440565b61394990886152d0565b90506000838361396184670de0b6b3a7640000614e28565b61396b9190614e28565b6139759190615440565b905060006139946bffffffffffffffffffffffff808916908a166152d0565b90506139ac816b033b2e3c9fd0803ce8000000614d3b565b8211156139e5576040517fe80fa38100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611cf581836152d0565b60006139f9610ce4565b511115613a0857613a086128a8565b5050565b6000808a8a8a8a8a8a8a8a8a604051602001613a3099989796959493929190615454565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613b36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610b58565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000838152600960209081526040808320815161014081018352815467ffffffffffffffff8116825273ffffffffffffffffffffffffffffffffffffffff68010000000000000000820481169583019590955263ffffffff7c01000000000000000000000000000000000000000000000000000000009091048116938201939093526001820154938416606082018190526bffffffffffffffffffffffff7401000000000000000000000000000000000000000090950485166080830152600283015494851660a083015269ffffffffffffffffffff6c0100000000000000000000000086041660c0830152760100000000000000000000000000000000000000000000850490931660e082015264ffffffffff7a0100000000000000000000000000000000000000000000000000009094049390931661010084015260030154610120830152613d01576002915050611de9565b6000858152600960205260408120818155600181018290556002810180547fff00000000000000000000000000000000000000000000000000000000000000169055600301819055613d516112ef565b905060008113613d90576040517f43d4cf6600000000000000000000000000000000000000000000000000000000815260048101829052602401610b58565b600081613da53a670de0b6b3a7640000614e28565b613daf9190615440565b9050600083610100015164ffffffffff1682613dcb9190614e28565b905060008460c0015169ffffffffffffffffffff1682613deb9190614f3c565b90506000613e0e60085473ffffffffffffffffffffffffffffffffffffffff1690565b90506000808273ffffffffffffffffffffffffffffffffffffffff16633fd67e058d8d8d8a89336040518763ffffffff1660e01b8152600401613e56969594939291906154e9565b60408051808303816000875af1158015613e74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e98919061555a565b9092509050613ea78186614f3c565b336000908152600d6020526040812080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff93841617905560c08a0151600e805469ffffffffffffffffffff90921693909291613f1491859116614f3c565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055508b7f30e1e299f390e4f3fa75e4489fed477e86183937a4a21a0ae7b137a4dff0353589600001518a60c001518489613f7a9190614f3c565b60808d015160c08e015169ffffffffffffffffffff16613f9a888d614f3c565b613fa49190614f3c565b613fae9190614f3c565b8760ff166006811115613fc357613fc3614db8565b604051613fd4959493929190615589565b60405180910390a28160ff166006811115613ff157613ff1614db8565b9c9b505050505050505050505050565b50805461400d906148e1565b6000825580601f1061401d575050565b601f016020900490600052602060002090810190610674919061405a565b604051806103e00160405280601f906020820280368337509192915050565b5b8082111561406f576000815560010161405b565b5090565b60008083601f84011261408557600080fd5b50813567ffffffffffffffff81111561409d57600080fd5b6020830191508360208285010111156140b557600080fd5b9250929050565b600080602083850312156140cf57600080fd5b823567ffffffffffffffff8111156140e657600080fd5b6140f285828601614073565b90969095509350505050565b6000815180845260005b8181101561412457602081850181015186830182015201614108565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000611de960208301846140fe565b73ffffffffffffffffffffffffffffffffffffffff8116811461067457600080fd5b6000602082840312156141a957600080fd5b8135611de981614175565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561422a5761422a6141b4565b604052919050565b600082601f83011261424357600080fd5b813567ffffffffffffffff81111561425d5761425d6141b4565b61428e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016141e3565b8181528460208386010111156142a357600080fd5b816020850160208301376000918101602001919091529392505050565b67ffffffffffffffff8116811461067457600080fd5b8035611041816142c0565b63ffffffff8116811461067457600080fd5b60008082840360a081121561430757600080fd5b833567ffffffffffffffff8082111561431f57600080fd5b61432b87838801614232565b945060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08401121561435d57600080fd5b6040519250608083019150828210818311171561437c5761437c6141b4565b50604052602084013561438e816142c0565b8152604084013561439e81614175565b602082015260608401356143b1816142e1565b60408201526080939093013560608401525092909150565b600081518084526020808501945080840160005b8381101561440f57815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016143dd565b509495945050505050565b60408152600061442d60408301856143c9565b6020838203818501528185518084528284019150828160051b85010183880160005b8381101561449b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030185526144898383516140fe565b9486019492509085019060010161444f565b50909998505050505050505050565b6bffffffffffffffffffffffff8116811461067457600080fd5b600080604083850312156144d757600080fd5b82356144e281614175565b915060208301356144f2816144aa565b809150509250929050565b60008060006040848603121561451257600080fd5b833561451d81614175565b9250602084013567ffffffffffffffff81111561453957600080fd5b61454586828701614073565b9497909650939450505050565b602081526000611de960208301846143c9565b60006020828403121561457757600080fd5b5035919050565b60006020828403121561459057600080fd5b813567ffffffffffffffff8111156145a757600080fd5b6145b384828501614232565b949350505050565b6000602082840312156145cd57600080fd5b813567ffffffffffffffff8111156145e457600080fd5b820160e08185031215611de957600080fd5b60008083601f84011261460857600080fd5b50813567ffffffffffffffff81111561462057600080fd5b6020830191508360208260051b85010111156140b557600080fd5b60008060008060008060008060e0898b03121561465757600080fd5b606089018a81111561466857600080fd5b8998503567ffffffffffffffff8082111561468257600080fd5b61468e8c838d01614073565b909950975060808b01359150808211156146a757600080fd5b6146b38c838d016145f6565b909750955060a08b01359150808211156146cc57600080fd5b506146d98b828c016145f6565b999c989b50969995989497949560c00135949350505050565b60008060008060006080868803121561470a57600080fd5b8535614715816142c0565b9450602086013567ffffffffffffffff81111561473157600080fd5b61473d88828901614073565b9095509350506040860135614751816142e1565b949793965091946060013592915050565b600067ffffffffffffffff82111561477c5761477c6141b4565b5060051b60200190565b600082601f83011261479757600080fd5b813560206147ac6147a783614762565b6141e3565b82815260059290921b840181019181810190868411156147cb57600080fd5b8286015b848110156147ef5780356147e281614175565b83529183019183016147cf565b509695505050505050565b60ff8116811461067457600080fd5b8035611041816147fa565b60008060008060008060c0878903121561482d57600080fd5b863567ffffffffffffffff8082111561484557600080fd5b6148518a838b01614786565b9750602089013591508082111561486757600080fd5b6148738a838b01614786565b965061488160408a01614809565b9550606089013591508082111561489757600080fd5b6148a38a838b01614232565b94506148b160808a016142d6565b935060a08901359150808211156148c757600080fd5b506148d489828a01614232565b9150509295509295509295565b600181811c908216806148f557607f821691505b60208210810361492e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f8211156105d057600081815260208120601f850160051c8101602086101561495b5750805b601f850160051c820191505b8181101561497a57828155600101614967565b505050505050565b67ffffffffffffffff83111561499a5761499a6141b4565b6149ae836149a883546148e1565b83614934565b6000601f841160018114614a0057600085156149ca5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355614a96565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b82811015614a4f5786850135825560209485019460019092019101614a2f565b5086821015614a8a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614b2c57614b2c614acc565b5060010190565b6bffffffffffffffffffffffff828116828216039080821115614b5857614b58614acc565b5092915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614b9457600080fd5b83018035915067ffffffffffffffff821115614baf57600080fd5b6020019150368190038213156140b557600080fd5b600060208284031215614bd657600080fd5b8135611de9816142c0565b600060208284031215614bf357600080fd5b8135611de9816142e1565b61ffff8116811461067457600080fd5b600060208284031215614c2057600080fd5b8135611de981614bfe565b600073ffffffffffffffffffffffffffffffffffffffff808b16835267ffffffffffffffff8a16602084015280891660408401525060e060608301528560e0830152610100868882850137600083880182015261ffff9590951660808301525060a081019290925263ffffffff1660c0820152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910101949350505050565b805169ffffffffffffffffffff8116811461104157600080fd5b600080600080600060a08688031215614d0357600080fd5b614d0c86614cd1565b9450602086015193506040860151925060608601519150614d2f60808701614cd1565b90509295509295909350565b8181038181111561069f5761069f614acc565b60ff818116838216019081111561069f5761069f614acc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff831680614da957614da9614d67565b8060ff84160491505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8183823760009101908152919050565b828152606082602083013760800192915050565b600060208284031215614e1d57600080fd5b8151611de9816144aa565b808202811582820484141761069f5761069f614acc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b63ffffffff818116838216019080821115614b5857614b58614acc565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614ebb8184018a6143c9565b90508281036080840152614ecf81896143c9565b905060ff871660a084015282810360c0840152614eec81876140fe565b905067ffffffffffffffff851660e0840152828103610100840152613ff181856140fe565b60006bffffffffffffffffffffffff80841680614f3057614f30614d67565b92169190910492915050565b6bffffffffffffffffffffffff818116838216019080821115614b5857614b58614acc565b600060ff821660ff8103614f7757614f77614acc565b60010192915050565b6bffffffffffffffffffffffff818116838216028082169190828114614fa857614fa8614acc565b505092915050565b60008060008060008060008060006101208a8c031215614fcf57600080fd5b8951614fda816142e1565b60208b0151909950614feb816142e1565b60408b0151909850614ffc816142e1565b60608b015190975061500d816142e1565b60808b015160a08c01519197509550615025816142e1565b935061503360c08b01614cd1565b925060e08a015161504381614bfe565b809250506101008a015190509295985092959850929598565b600080600080600060a0868803121561507457600080fd5b855161507f816144aa565b80955050602080870151615092816144aa565b60408801519095506150a381614175565b60608801519094506150b481614175565b608088015190935067ffffffffffffffff8111156150d157600080fd5b8701601f810189136150e257600080fd5b80516150f06147a782614762565b81815260059190911b8201830190838101908b83111561510f57600080fd5b928401925b8284101561513657835161512781614175565b82529284019290840190615114565b80955050505050509295509295909350565b60008060006060848603121561515d57600080fd5b8351801515811461516d57600080fd5b602085015190935061517e816142c0565b604085015190925061518f816142c0565b809150509250925092565b67ffffffffffffffff818116838216019080821115614b5857614b58614acc565b815167ffffffffffffffff168152610140810160208301516151f5602084018273ffffffffffffffffffffffffffffffffffffffff169052565b50604083015161520d604084018263ffffffff169052565b506060830151615235606084018273ffffffffffffffffffffffffffffffffffffffff169052565b50608083015161525560808401826bffffffffffffffffffffffff169052565b5060a083015161527560a08401826bffffffffffffffffffffffff169052565b5060c083015161529360c084018269ffffffffffffffffffff169052565b5060e08301516152ab60e084018263ffffffff169052565b506101008381015164ffffffffff811684830152505061012092830151919092015290565b8082018082111561069f5761069f614acc565b600082601f8301126152f457600080fd5b813560206153046147a783614762565b82815260059290921b8401810191818101908684111561532357600080fd5b8286015b848110156147ef57803567ffffffffffffffff8111156153475760008081fd5b6153558986838b0101614232565b845250918301918301615327565b60008060006060848603121561537857600080fd5b833567ffffffffffffffff8082111561539057600080fd5b818601915086601f8301126153a457600080fd5b813560206153b46147a783614762565b82815260059290921b8401810191818101908a8411156153d357600080fd5b948201945b838610156153f1578535825294820194908201906153d8565b9750508701359250508082111561540757600080fd5b615413878388016152e3565b9350604086013591508082111561542957600080fd5b50615436868287016152e3565b9150509250925092565b60008261544f5761544f614d67565b500490565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b16604085015281606085015261549b8285018b6143c9565b915083820360808501526154af828a6143c9565b915060ff881660a085015283820360c08501526154cc82886140fe565b90861660e08501528381036101008501529050613ff181856140fe565b86815260c06020820152600061550260c08301886140fe565b828103604084015261551481886140fe565b6bffffffffffffffffffffffff96871660608501529490951660808301525073ffffffffffffffffffffffffffffffffffffffff9190911660a090910152949350505050565b6000806040838503121561556d57600080fd5b8251615578816147fa565b60208401519092506144f2816144aa565b67ffffffffffffffff8616815269ffffffffffffffffffff851660208201526bffffffffffffffffffffffff84811660408301528316606082015260a08101600783106155ff577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826080830152969550505050505056fea164736f6c6343000813000a",
}

var FunctionsCoordinatorABI = FunctionsCoordinatorMetaData.ABI

var FunctionsCoordinatorBin = FunctionsCoordinatorMetaData.Bin

func DeployFunctionsCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, config []byte, linkToNativeFeed common.Address) (common.Address, *types.Transaction, *FunctionsCoordinator, error) {
	parsed, err := FunctionsCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FunctionsCoordinatorBin), backend, router, config, linkToNativeFeed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FunctionsCoordinator{FunctionsCoordinatorCaller: FunctionsCoordinatorCaller{contract: contract}, FunctionsCoordinatorTransactor: FunctionsCoordinatorTransactor{contract: contract}, FunctionsCoordinatorFilterer: FunctionsCoordinatorFilterer{contract: contract}}, nil
}

type FunctionsCoordinator struct {
	address common.Address
	abi     abi.ABI
	FunctionsCoordinatorCaller
	FunctionsCoordinatorTransactor
	FunctionsCoordinatorFilterer
}

type FunctionsCoordinatorCaller struct {
	contract *bind.BoundContract
}

type FunctionsCoordinatorTransactor struct {
	contract *bind.BoundContract
}

type FunctionsCoordinatorFilterer struct {
	contract *bind.BoundContract
}

type FunctionsCoordinatorSession struct {
	Contract     *FunctionsCoordinator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FunctionsCoordinatorCallerSession struct {
	Contract *FunctionsCoordinatorCaller
	CallOpts bind.CallOpts
}

type FunctionsCoordinatorTransactorSession struct {
	Contract     *FunctionsCoordinatorTransactor
	TransactOpts bind.TransactOpts
}

type FunctionsCoordinatorRaw struct {
	Contract *FunctionsCoordinator
}

type FunctionsCoordinatorCallerRaw struct {
	Contract *FunctionsCoordinatorCaller
}

type FunctionsCoordinatorTransactorRaw struct {
	Contract *FunctionsCoordinatorTransactor
}

func NewFunctionsCoordinator(address common.Address, backend bind.ContractBackend) (*FunctionsCoordinator, error) {
	abi, err := abi.JSON(strings.NewReader(FunctionsCoordinatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFunctionsCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinator{address: address, abi: abi, FunctionsCoordinatorCaller: FunctionsCoordinatorCaller{contract: contract}, FunctionsCoordinatorTransactor: FunctionsCoordinatorTransactor{contract: contract}, FunctionsCoordinatorFilterer: FunctionsCoordinatorFilterer{contract: contract}}, nil
}

func NewFunctionsCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*FunctionsCoordinatorCaller, error) {
	contract, err := bindFunctionsCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorCaller{contract: contract}, nil
}

func NewFunctionsCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*FunctionsCoordinatorTransactor, error) {
	contract, err := bindFunctionsCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorTransactor{contract: contract}, nil
}

func NewFunctionsCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*FunctionsCoordinatorFilterer, error) {
	contract, err := bindFunctionsCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorFilterer{contract: contract}, nil
}

func bindFunctionsCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FunctionsCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FunctionsCoordinator.Contract.FunctionsCoordinatorCaller.contract.Call(opts, result, method, params...)
}

func (_FunctionsCoordinator *FunctionsCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.FunctionsCoordinatorTransactor.contract.Transfer(opts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.FunctionsCoordinatorTransactor.contract.Transact(opts, method, params...)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FunctionsCoordinator.Contract.contract.Call(opts, result, method, params...)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.contract.Transfer(opts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.contract.Transact(opts, method, params...)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) EstimateCost(opts *bind.CallOpts, subscriptionId uint64, data []byte, callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "estimateCost", subscriptionId, data, callbackGasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) EstimateCost(subscriptionId uint64, data []byte, callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _FunctionsCoordinator.Contract.EstimateCost(&_FunctionsCoordinator.CallOpts, subscriptionId, data, callbackGasLimit, gasPrice)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) EstimateCost(subscriptionId uint64, data []byte, callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _FunctionsCoordinator.Contract.EstimateCost(&_FunctionsCoordinator.CallOpts, subscriptionId, data, callbackGasLimit, gasPrice)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetAdminFee(opts *bind.CallOpts, arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getAdminFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetAdminFee(arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error) {
	return _FunctionsCoordinator.Contract.GetAdminFee(&_FunctionsCoordinator.CallOpts, arg0, arg1)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetAdminFee(arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error) {
	return _FunctionsCoordinator.Contract.GetAdminFee(&_FunctionsCoordinator.CallOpts, arg0, arg1)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetAllNodePublicKeys(opts *bind.CallOpts) ([]common.Address, [][]byte, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getAllNodePublicKeys")

	if err != nil {
		return *new([]common.Address), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return out0, out1, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetAllNodePublicKeys() ([]common.Address, [][]byte, error) {
	return _FunctionsCoordinator.Contract.GetAllNodePublicKeys(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetAllNodePublicKeys() ([]common.Address, [][]byte, error) {
	return _FunctionsCoordinator.Contract.GetAllNodePublicKeys(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetConfig(opts *bind.CallOpts) (GetConfig,

	error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getConfig")

	outstruct := new(GetConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxCallbackGasLimit = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.FeedStalenessSeconds = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.GasOverheadAfterCallback = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FallbackNativePerUnitLink = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.GasOverheadBeforeCallback = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.LinkPriceFeed = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.MaxSupportedRequestDataVersion = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.FulfillmentGasPriceOverEstimationBP = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetConfig() (GetConfig,

	error) {
	return _FunctionsCoordinator.Contract.GetConfig(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetConfig() (GetConfig,

	error) {
	return _FunctionsCoordinator.Contract.GetConfig(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetConfigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getConfigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetConfigHash() ([32]byte, error) {
	return _FunctionsCoordinator.Contract.GetConfigHash(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetConfigHash() ([32]byte, error) {
	return _FunctionsCoordinator.Contract.GetConfigHash(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetDONFee(opts *bind.CallOpts, arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getDONFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetDONFee(arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error) {
	return _FunctionsCoordinator.Contract.GetDONFee(&_FunctionsCoordinator.CallOpts, arg0, arg1)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetDONFee(arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error) {
	return _FunctionsCoordinator.Contract.GetDONFee(&_FunctionsCoordinator.CallOpts, arg0, arg1)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetDONPublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getDONPublicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetDONPublicKey() ([]byte, error) {
	return _FunctionsCoordinator.Contract.GetDONPublicKey(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetDONPublicKey() ([]byte, error) {
	return _FunctionsCoordinator.Contract.GetDONPublicKey(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetFeedData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getFeedData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetFeedData() (*big.Int, error) {
	return _FunctionsCoordinator.Contract.GetFeedData(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetFeedData() (*big.Int, error) {
	return _FunctionsCoordinator.Contract.GetFeedData(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) GetThresholdPublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "getThresholdPublicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) GetThresholdPublicKey() ([]byte, error) {
	return _FunctionsCoordinator.Contract.GetThresholdPublicKey(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) GetThresholdPublicKey() ([]byte, error) {
	return _FunctionsCoordinator.Contract.GetThresholdPublicKey(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _FunctionsCoordinator.Contract.LatestConfigDetails(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _FunctionsCoordinator.Contract.LatestConfigDetails(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _FunctionsCoordinator.Contract.LatestConfigDigestAndEpoch(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _FunctionsCoordinator.Contract.LatestConfigDigestAndEpoch(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) Owner() (common.Address, error) {
	return _FunctionsCoordinator.Contract.Owner(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) Owner() (common.Address, error) {
	return _FunctionsCoordinator.Contract.Owner(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) Transmitters() ([]common.Address, error) {
	return _FunctionsCoordinator.Contract.Transmitters(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) Transmitters() ([]common.Address, error) {
	return _FunctionsCoordinator.Contract.Transmitters(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FunctionsCoordinator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) TypeAndVersion() (string, error) {
	return _FunctionsCoordinator.Contract.TypeAndVersion(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorCallerSession) TypeAndVersion() (string, error) {
	return _FunctionsCoordinator.Contract.TypeAndVersion(&_FunctionsCoordinator.CallOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "acceptOwnership")
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.AcceptOwnership(&_FunctionsCoordinator.TransactOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.AcceptOwnership(&_FunctionsCoordinator.TransactOpts)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) DeleteCommitment(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "deleteCommitment", requestId)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) DeleteCommitment(requestId [32]byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.DeleteCommitment(&_FunctionsCoordinator.TransactOpts, requestId)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) DeleteCommitment(requestId [32]byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.DeleteCommitment(&_FunctionsCoordinator.TransactOpts, requestId)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) DeleteNodePublicKey(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "deleteNodePublicKey", node)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) DeleteNodePublicKey(node common.Address) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.DeleteNodePublicKey(&_FunctionsCoordinator.TransactOpts, node)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) DeleteNodePublicKey(node common.Address) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.DeleteNodePublicKey(&_FunctionsCoordinator.TransactOpts, node)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.OracleWithdraw(&_FunctionsCoordinator.TransactOpts, recipient, amount)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.OracleWithdraw(&_FunctionsCoordinator.TransactOpts, recipient, amount)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) SendRequest(opts *bind.TransactOpts, request IFunctionsCoordinatorRequest) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "sendRequest", request)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) SendRequest(request IFunctionsCoordinatorRequest) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SendRequest(&_FunctionsCoordinator.TransactOpts, request)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) SendRequest(request IFunctionsCoordinatorRequest) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SendRequest(&_FunctionsCoordinator.TransactOpts, request)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetConfig(&_FunctionsCoordinator.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetConfig(&_FunctionsCoordinator.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) SetDONPublicKey(opts *bind.TransactOpts, donPublicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "setDONPublicKey", donPublicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) SetDONPublicKey(donPublicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetDONPublicKey(&_FunctionsCoordinator.TransactOpts, donPublicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) SetDONPublicKey(donPublicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetDONPublicKey(&_FunctionsCoordinator.TransactOpts, donPublicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) SetNodePublicKey(opts *bind.TransactOpts, node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "setNodePublicKey", node, publicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) SetNodePublicKey(node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetNodePublicKey(&_FunctionsCoordinator.TransactOpts, node, publicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) SetNodePublicKey(node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetNodePublicKey(&_FunctionsCoordinator.TransactOpts, node, publicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) SetThresholdPublicKey(opts *bind.TransactOpts, thresholdPublicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "setThresholdPublicKey", thresholdPublicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) SetThresholdPublicKey(thresholdPublicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetThresholdPublicKey(&_FunctionsCoordinator.TransactOpts, thresholdPublicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) SetThresholdPublicKey(thresholdPublicKey []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.SetThresholdPublicKey(&_FunctionsCoordinator.TransactOpts, thresholdPublicKey)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "transferOwnership", to)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.TransferOwnership(&_FunctionsCoordinator.TransactOpts, to)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.TransferOwnership(&_FunctionsCoordinator.TransactOpts, to)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.Transmit(&_FunctionsCoordinator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.Transmit(&_FunctionsCoordinator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactor) UpdateConfig(opts *bind.TransactOpts, config []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.contract.Transact(opts, "updateConfig", config)
}

func (_FunctionsCoordinator *FunctionsCoordinatorSession) UpdateConfig(config []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.UpdateConfig(&_FunctionsCoordinator.TransactOpts, config)
}

func (_FunctionsCoordinator *FunctionsCoordinatorTransactorSession) UpdateConfig(config []byte) (*types.Transaction, error) {
	return _FunctionsCoordinator.Contract.UpdateConfig(&_FunctionsCoordinator.TransactOpts, config)
}

type FunctionsCoordinatorBillingEndIterator struct {
	Event *FunctionsCoordinatorBillingEnd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorBillingEndIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorBillingEnd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorBillingEnd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorBillingEndIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorBillingEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorBillingEnd struct {
	RequestId          [32]byte
	SubscriptionId     uint64
	SignerPayment      *big.Int
	TransmitterPayment *big.Int
	TotalCost          *big.Int
	Result             uint8
	Raw                types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterBillingEnd(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorBillingEndIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "BillingEnd", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorBillingEndIterator{contract: _FunctionsCoordinator.contract, event: "BillingEnd", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchBillingEnd(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorBillingEnd, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "BillingEnd", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorBillingEnd)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "BillingEnd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseBillingEnd(log types.Log) (*FunctionsCoordinatorBillingEnd, error) {
	event := new(FunctionsCoordinatorBillingEnd)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "BillingEnd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorBillingStartIterator struct {
	Event *FunctionsCoordinatorBillingStart

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorBillingStartIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorBillingStart)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorBillingStart)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorBillingStartIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorBillingStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorBillingStart struct {
	RequestId  [32]byte
	Commitment FunctionsBillingCommitment
	Raw        types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterBillingStart(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorBillingStartIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "BillingStart", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorBillingStartIterator{contract: _FunctionsCoordinator.contract, event: "BillingStart", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchBillingStart(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorBillingStart, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "BillingStart", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorBillingStart)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "BillingStart", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseBillingStart(log types.Log) (*FunctionsCoordinatorBillingStart, error) {
	event := new(FunctionsCoordinatorBillingStart)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "BillingStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorConfigChangedIterator struct {
	Event *FunctionsCoordinatorConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorConfigChangedIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorConfigChanged struct {
	MaxCallbackGasLimit                 uint32
	FeedStalenessSeconds                uint32
	GasOverheadBeforeCallback           uint32
	GasOverheadAfterCallback            uint32
	FallbackNativePerUnitLink           *big.Int
	DonFee                              *big.Int
	MaxSupportedRequestDataVersion      uint16
	FulfillmentGasPriceOverEstimationBP *big.Int
	Raw                                 types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*FunctionsCoordinatorConfigChangedIterator, error) {

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorConfigChangedIterator{contract: _FunctionsCoordinator.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorConfigChanged) (event.Subscription, error) {

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorConfigChanged)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseConfigChanged(log types.Log) (*FunctionsCoordinatorConfigChanged, error) {
	event := new(FunctionsCoordinatorConfigChanged)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorConfigSetIterator struct {
	Event *FunctionsCoordinatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*FunctionsCoordinatorConfigSetIterator, error) {

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorConfigSetIterator{contract: _FunctionsCoordinator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorConfigSet)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseConfigSet(log types.Log) (*FunctionsCoordinatorConfigSet, error) {
	event := new(FunctionsCoordinatorConfigSet)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorCostExceedsCommitmentIterator struct {
	Event *FunctionsCoordinatorCostExceedsCommitment

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorCostExceedsCommitmentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorCostExceedsCommitment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorCostExceedsCommitment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorCostExceedsCommitmentIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorCostExceedsCommitmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorCostExceedsCommitment struct {
	RequestId [32]byte
	Raw       types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterCostExceedsCommitment(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorCostExceedsCommitmentIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "CostExceedsCommitment", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorCostExceedsCommitmentIterator{contract: _FunctionsCoordinator.contract, event: "CostExceedsCommitment", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchCostExceedsCommitment(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorCostExceedsCommitment, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "CostExceedsCommitment", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorCostExceedsCommitment)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "CostExceedsCommitment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseCostExceedsCommitment(log types.Log) (*FunctionsCoordinatorCostExceedsCommitment, error) {
	event := new(FunctionsCoordinatorCostExceedsCommitment)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "CostExceedsCommitment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorInsufficientGasProvidedIterator struct {
	Event *FunctionsCoordinatorInsufficientGasProvided

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorInsufficientGasProvidedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorInsufficientGasProvided)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorInsufficientGasProvided)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorInsufficientGasProvidedIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorInsufficientGasProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorInsufficientGasProvided struct {
	RequestId [32]byte
	Raw       types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterInsufficientGasProvided(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorInsufficientGasProvidedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "InsufficientGasProvided", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorInsufficientGasProvidedIterator{contract: _FunctionsCoordinator.contract, event: "InsufficientGasProvided", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchInsufficientGasProvided(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorInsufficientGasProvided, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "InsufficientGasProvided", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorInsufficientGasProvided)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "InsufficientGasProvided", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseInsufficientGasProvided(log types.Log) (*FunctionsCoordinatorInsufficientGasProvided, error) {
	event := new(FunctionsCoordinatorInsufficientGasProvided)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "InsufficientGasProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorInsufficientSubscriptionBalanceIterator struct {
	Event *FunctionsCoordinatorInsufficientSubscriptionBalance

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorInsufficientSubscriptionBalanceIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorInsufficientSubscriptionBalance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorInsufficientSubscriptionBalance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorInsufficientSubscriptionBalanceIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorInsufficientSubscriptionBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorInsufficientSubscriptionBalance struct {
	RequestId [32]byte
	Raw       types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterInsufficientSubscriptionBalance(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorInsufficientSubscriptionBalanceIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "InsufficientSubscriptionBalance", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorInsufficientSubscriptionBalanceIterator{contract: _FunctionsCoordinator.contract, event: "InsufficientSubscriptionBalance", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchInsufficientSubscriptionBalance(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorInsufficientSubscriptionBalance, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "InsufficientSubscriptionBalance", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorInsufficientSubscriptionBalance)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "InsufficientSubscriptionBalance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseInsufficientSubscriptionBalance(log types.Log) (*FunctionsCoordinatorInsufficientSubscriptionBalance, error) {
	event := new(FunctionsCoordinatorInsufficientSubscriptionBalance)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "InsufficientSubscriptionBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorInvalidRequestIDIterator struct {
	Event *FunctionsCoordinatorInvalidRequestID

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorInvalidRequestIDIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorInvalidRequestID)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorInvalidRequestID)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorInvalidRequestIDIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorInvalidRequestIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorInvalidRequestID struct {
	RequestId [32]byte
	Raw       types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterInvalidRequestID(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorInvalidRequestIDIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "InvalidRequestID", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorInvalidRequestIDIterator{contract: _FunctionsCoordinator.contract, event: "InvalidRequestID", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchInvalidRequestID(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorInvalidRequestID, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "InvalidRequestID", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorInvalidRequestID)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "InvalidRequestID", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseInvalidRequestID(log types.Log) (*FunctionsCoordinatorInvalidRequestID, error) {
	event := new(FunctionsCoordinatorInvalidRequestID)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "InvalidRequestID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorOracleRequestIterator struct {
	Event *FunctionsCoordinatorOracleRequest

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorOracleRequestIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorOracleRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorOracleRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorOracleRequestIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorOracleRequest struct {
	RequestId          [32]byte
	RequestingContract common.Address
	RequestInitiator   common.Address
	SubscriptionId     uint64
	SubscriptionOwner  common.Address
	Data               []byte
	DataVersion        uint16
	Flags              [32]byte
	CallbackGasLimit   uint64
	Raw                types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterOracleRequest(opts *bind.FilterOpts, requestId [][32]byte, requestingContract []common.Address) (*FunctionsCoordinatorOracleRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestingContractRule []interface{}
	for _, requestingContractItem := range requestingContract {
		requestingContractRule = append(requestingContractRule, requestingContractItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "OracleRequest", requestIdRule, requestingContractRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorOracleRequestIterator{contract: _FunctionsCoordinator.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOracleRequest, requestId [][32]byte, requestingContract []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestingContractRule []interface{}
	for _, requestingContractItem := range requestingContract {
		requestingContractRule = append(requestingContractRule, requestingContractItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "OracleRequest", requestIdRule, requestingContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorOracleRequest)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "OracleRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseOracleRequest(log types.Log) (*FunctionsCoordinatorOracleRequest, error) {
	event := new(FunctionsCoordinatorOracleRequest)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorOracleResponseIterator struct {
	Event *FunctionsCoordinatorOracleResponse

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorOracleResponseIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorOracleResponse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorOracleResponse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorOracleResponseIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorOracleResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorOracleResponse struct {
	RequestId   [32]byte
	Transmitter common.Address
	Raw         types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterOracleResponse(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorOracleResponseIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorOracleResponseIterator{contract: _FunctionsCoordinator.contract, event: "OracleResponse", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOracleResponse, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorOracleResponse)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "OracleResponse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseOracleResponse(log types.Log) (*FunctionsCoordinatorOracleResponse, error) {
	event := new(FunctionsCoordinatorOracleResponse)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "OracleResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorOwnershipTransferRequestedIterator struct {
	Event *FunctionsCoordinatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionsCoordinatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorOwnershipTransferRequestedIterator{contract: _FunctionsCoordinator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorOwnershipTransferRequested)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*FunctionsCoordinatorOwnershipTransferRequested, error) {
	event := new(FunctionsCoordinatorOwnershipTransferRequested)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorOwnershipTransferredIterator struct {
	Event *FunctionsCoordinatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionsCoordinatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorOwnershipTransferredIterator{contract: _FunctionsCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorOwnershipTransferred)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*FunctionsCoordinatorOwnershipTransferred, error) {
	event := new(FunctionsCoordinatorOwnershipTransferred)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorRequestTimedOutIterator struct {
	Event *FunctionsCoordinatorRequestTimedOut

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorRequestTimedOutIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorRequestTimedOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorRequestTimedOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorRequestTimedOutIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorRequestTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorRequestTimedOut struct {
	RequestId [32]byte
	Raw       types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterRequestTimedOut(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorRequestTimedOutIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "RequestTimedOut", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorRequestTimedOutIterator{contract: _FunctionsCoordinator.contract, event: "RequestTimedOut", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchRequestTimedOut(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorRequestTimedOut, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "RequestTimedOut", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorRequestTimedOut)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "RequestTimedOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseRequestTimedOut(log types.Log) (*FunctionsCoordinatorRequestTimedOut, error) {
	event := new(FunctionsCoordinatorRequestTimedOut)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "RequestTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FunctionsCoordinatorTransmittedIterator struct {
	Event *FunctionsCoordinatorTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FunctionsCoordinatorTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsCoordinatorTransmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FunctionsCoordinatorTransmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FunctionsCoordinatorTransmittedIterator) Error() error {
	return it.fail
}

func (it *FunctionsCoordinatorTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FunctionsCoordinatorTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) FilterTransmitted(opts *bind.FilterOpts) (*FunctionsCoordinatorTransmittedIterator, error) {

	logs, sub, err := _FunctionsCoordinator.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &FunctionsCoordinatorTransmittedIterator{contract: _FunctionsCoordinator.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorTransmitted) (event.Subscription, error) {

	logs, sub, err := _FunctionsCoordinator.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FunctionsCoordinatorTransmitted)
				if err := _FunctionsCoordinator.contract.UnpackLog(event, "Transmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FunctionsCoordinator *FunctionsCoordinatorFilterer) ParseTransmitted(log types.Log) (*FunctionsCoordinatorTransmitted, error) {
	event := new(FunctionsCoordinatorTransmitted)
	if err := _FunctionsCoordinator.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetConfig struct {
	MaxCallbackGasLimit                 uint32
	FeedStalenessSeconds                uint32
	GasOverheadAfterCallback            *big.Int
	FallbackNativePerUnitLink           *big.Int
	GasOverheadBeforeCallback           uint32
	LinkPriceFeed                       common.Address
	MaxSupportedRequestDataVersion      uint16
	FulfillmentGasPriceOverEstimationBP *big.Int
}
type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_FunctionsCoordinator *FunctionsCoordinator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _FunctionsCoordinator.abi.Events["BillingEnd"].ID:
		return _FunctionsCoordinator.ParseBillingEnd(log)
	case _FunctionsCoordinator.abi.Events["BillingStart"].ID:
		return _FunctionsCoordinator.ParseBillingStart(log)
	case _FunctionsCoordinator.abi.Events["ConfigChanged"].ID:
		return _FunctionsCoordinator.ParseConfigChanged(log)
	case _FunctionsCoordinator.abi.Events["ConfigSet"].ID:
		return _FunctionsCoordinator.ParseConfigSet(log)
	case _FunctionsCoordinator.abi.Events["CostExceedsCommitment"].ID:
		return _FunctionsCoordinator.ParseCostExceedsCommitment(log)
	case _FunctionsCoordinator.abi.Events["InsufficientGasProvided"].ID:
		return _FunctionsCoordinator.ParseInsufficientGasProvided(log)
	case _FunctionsCoordinator.abi.Events["InsufficientSubscriptionBalance"].ID:
		return _FunctionsCoordinator.ParseInsufficientSubscriptionBalance(log)
	case _FunctionsCoordinator.abi.Events["InvalidRequestID"].ID:
		return _FunctionsCoordinator.ParseInvalidRequestID(log)
	case _FunctionsCoordinator.abi.Events["OracleRequest"].ID:
		return _FunctionsCoordinator.ParseOracleRequest(log)
	case _FunctionsCoordinator.abi.Events["OracleResponse"].ID:
		return _FunctionsCoordinator.ParseOracleResponse(log)
	case _FunctionsCoordinator.abi.Events["OwnershipTransferRequested"].ID:
		return _FunctionsCoordinator.ParseOwnershipTransferRequested(log)
	case _FunctionsCoordinator.abi.Events["OwnershipTransferred"].ID:
		return _FunctionsCoordinator.ParseOwnershipTransferred(log)
	case _FunctionsCoordinator.abi.Events["RequestTimedOut"].ID:
		return _FunctionsCoordinator.ParseRequestTimedOut(log)
	case _FunctionsCoordinator.abi.Events["Transmitted"].ID:
		return _FunctionsCoordinator.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (FunctionsCoordinatorBillingEnd) Topic() common.Hash {
	return common.HexToHash("0x30e1e299f390e4f3fa75e4489fed477e86183937a4a21a0ae7b137a4dff03535")
}

func (FunctionsCoordinatorBillingStart) Topic() common.Hash {
	return common.HexToHash("0xe2cc993cebc49ce369b0174671d90dbae2201decdae368805de999b437336dcd")
}

func (FunctionsCoordinatorConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9136fb3a39fbe82d27ef8048354b8fcc93cf828f72e5c877d8135fde8dc7fd73")
}

func (FunctionsCoordinatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (FunctionsCoordinatorCostExceedsCommitment) Topic() common.Hash {
	return common.HexToHash("0x689f49eacf6db85082542f4b7873f4a519744a4e093fd002310545bfb4354cc8")
}

func (FunctionsCoordinatorInsufficientGasProvided) Topic() common.Hash {
	return common.HexToHash("0x357a60574f143e144d62ddc89636fee45ba5e2f8c9a30b9f91e28f8dfe5a56e0")
}

func (FunctionsCoordinatorInsufficientSubscriptionBalance) Topic() common.Hash {
	return common.HexToHash("0xb83d9204606ee3780daa106ad026c4c2853f5af126edbb8fe15e813e62dbd94d")
}

func (FunctionsCoordinatorInvalidRequestID) Topic() common.Hash {
	return common.HexToHash("0xa1c120e327c9ad8b075793878c88d59b8934b97ae37117faa3bb21616237f7be")
}

func (FunctionsCoordinatorOracleRequest) Topic() common.Hash {
	return common.HexToHash("0xd6bf2929a5458fc22a90821275e4f54e2566c579e4bb3d38184880bf1701762c")
}

func (FunctionsCoordinatorOracleResponse) Topic() common.Hash {
	return common.HexToHash("0xc708e0440951fd63499c0f7a73819b469ee5dd3ecc356c0ab4eb7f18389009d9")
}

func (FunctionsCoordinatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FunctionsCoordinatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (FunctionsCoordinatorRequestTimedOut) Topic() common.Hash {
	return common.HexToHash("0xf1ca1e9147be737b04a2b018a79405f687a97de8dd8a2559bbe62357343af414")
}

func (FunctionsCoordinatorTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_FunctionsCoordinator *FunctionsCoordinator) Address() common.Address {
	return _FunctionsCoordinator.address
}

type FunctionsCoordinatorInterface interface {
	EstimateCost(opts *bind.CallOpts, subscriptionId uint64, data []byte, callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error)

	GetAdminFee(opts *bind.CallOpts, arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error)

	GetAllNodePublicKeys(opts *bind.CallOpts) ([]common.Address, [][]byte, error)

	GetConfig(opts *bind.CallOpts) (GetConfig,

		error)

	GetConfigHash(opts *bind.CallOpts) ([32]byte, error)

	GetDONFee(opts *bind.CallOpts, arg0 []byte, arg1 IFunctionsBillingRequestBilling) (*big.Int, error)

	GetDONPublicKey(opts *bind.CallOpts) ([]byte, error)

	GetFeedData(opts *bind.CallOpts) (*big.Int, error)

	GetThresholdPublicKey(opts *bind.CallOpts) ([]byte, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	DeleteCommitment(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error)

	DeleteNodePublicKey(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error)

	OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	SendRequest(opts *bind.TransactOpts, request IFunctionsCoordinatorRequest) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetDONPublicKey(opts *bind.TransactOpts, donPublicKey []byte) (*types.Transaction, error)

	SetNodePublicKey(opts *bind.TransactOpts, node common.Address, publicKey []byte) (*types.Transaction, error)

	SetThresholdPublicKey(opts *bind.TransactOpts, thresholdPublicKey []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	UpdateConfig(opts *bind.TransactOpts, config []byte) (*types.Transaction, error)

	FilterBillingEnd(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorBillingEndIterator, error)

	WatchBillingEnd(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorBillingEnd, requestId [][32]byte) (event.Subscription, error)

	ParseBillingEnd(log types.Log) (*FunctionsCoordinatorBillingEnd, error)

	FilterBillingStart(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorBillingStartIterator, error)

	WatchBillingStart(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorBillingStart, requestId [][32]byte) (event.Subscription, error)

	ParseBillingStart(log types.Log) (*FunctionsCoordinatorBillingStart, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*FunctionsCoordinatorConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*FunctionsCoordinatorConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*FunctionsCoordinatorConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*FunctionsCoordinatorConfigSet, error)

	FilterCostExceedsCommitment(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorCostExceedsCommitmentIterator, error)

	WatchCostExceedsCommitment(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorCostExceedsCommitment, requestId [][32]byte) (event.Subscription, error)

	ParseCostExceedsCommitment(log types.Log) (*FunctionsCoordinatorCostExceedsCommitment, error)

	FilterInsufficientGasProvided(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorInsufficientGasProvidedIterator, error)

	WatchInsufficientGasProvided(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorInsufficientGasProvided, requestId [][32]byte) (event.Subscription, error)

	ParseInsufficientGasProvided(log types.Log) (*FunctionsCoordinatorInsufficientGasProvided, error)

	FilterInsufficientSubscriptionBalance(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorInsufficientSubscriptionBalanceIterator, error)

	WatchInsufficientSubscriptionBalance(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorInsufficientSubscriptionBalance, requestId [][32]byte) (event.Subscription, error)

	ParseInsufficientSubscriptionBalance(log types.Log) (*FunctionsCoordinatorInsufficientSubscriptionBalance, error)

	FilterInvalidRequestID(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorInvalidRequestIDIterator, error)

	WatchInvalidRequestID(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorInvalidRequestID, requestId [][32]byte) (event.Subscription, error)

	ParseInvalidRequestID(log types.Log) (*FunctionsCoordinatorInvalidRequestID, error)

	FilterOracleRequest(opts *bind.FilterOpts, requestId [][32]byte, requestingContract []common.Address) (*FunctionsCoordinatorOracleRequestIterator, error)

	WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOracleRequest, requestId [][32]byte, requestingContract []common.Address) (event.Subscription, error)

	ParseOracleRequest(log types.Log) (*FunctionsCoordinatorOracleRequest, error)

	FilterOracleResponse(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorOracleResponseIterator, error)

	WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOracleResponse, requestId [][32]byte) (event.Subscription, error)

	ParseOracleResponse(log types.Log) (*FunctionsCoordinatorOracleResponse, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionsCoordinatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FunctionsCoordinatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionsCoordinatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FunctionsCoordinatorOwnershipTransferred, error)

	FilterRequestTimedOut(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsCoordinatorRequestTimedOutIterator, error)

	WatchRequestTimedOut(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorRequestTimedOut, requestId [][32]byte) (event.Subscription, error)

	ParseRequestTimedOut(log types.Log) (*FunctionsCoordinatorRequestTimedOut, error)

	FilterTransmitted(opts *bind.FilterOpts) (*FunctionsCoordinatorTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *FunctionsCoordinatorTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*FunctionsCoordinatorTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
