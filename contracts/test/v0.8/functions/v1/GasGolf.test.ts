import { ethers } from 'hardhat'
import { BigNumber } from 'ethers'
import {
  accessControlMockPrivateKey,
  FunctionsContracts,
  FunctionsRoles,
  getEventArg,
  getSetupFactory,
  ids,
} from './utils'
import { stringToBytes } from '../../../test-helpers/helpers'

const setup = getSetupFactory()
let contracts: FunctionsContracts
let roles: FunctionsRoles

const baselineGasUsed = 754223 // TODO: Update baseline
let currentGasUsed = 0

beforeEach(async () => {
  ;({ contracts, roles } = setup())
})

after(() => {
  const score = currentGasUsed - baselineGasUsed
  console.log(`\n               ⛳ Par   : ${baselineGasUsed} gas`)
  console.log(`\n               🏌️  You   : ${currentGasUsed} gas`)
  console.log(`\n               🚩 Score : ${score} gas`)
})

describe('Gas Golf', () => {
  it('taking a swing', async () => {
    // User signs Terms of Service
    const messageHash = await contracts.accessControl.getMessageHash(
      roles.consumerAddress,
      roles.consumerAddress,
    )
    const wallet = new ethers.Wallet(accessControlMockPrivateKey)
    const flatSignature = await wallet.signMessage(
      ethers.utils.arrayify(messageHash),
    )
    let { r, s, v } = ethers.utils.splitSignature(flatSignature)
    const acceptTermsOfServiceTx = await contracts.accessControl
      .connect(roles.consumer)
      .acceptTermsOfService(
        roles.consumerAddress,
        roles.consumerAddress,
        r,
        s,
        v,
      )
    const { gasUsed: acceptTermsOfServiceGasUsed } =
      await acceptTermsOfServiceTx.wait()

    // User creates a new Subscription
    const createSubscriptionTx = await contracts.router
      .connect(roles.consumer)
      .createSubscription()
    const createSubscriptionTxReceipt = await createSubscriptionTx.wait()
    const createSubscriptionTxGasUsed = createSubscriptionTxReceipt.gasUsed
    const subscriptionId =
      createSubscriptionTxReceipt.events[0].args['subscriptionId'].toNumber()

    // User adds a consuming contract to their Subscription
    const addConsumerTx = await contracts.router
      .connect(roles.consumer)
      .addConsumer(subscriptionId, contracts.client.address)
    const { gasUsed: addConsumerTxGasUsed } = await addConsumerTx.wait()

    // User funds their subscription
    const transferAndCallTx = await contracts.linkToken
      .connect(roles.subOwner)
      .transferAndCall(
        contracts.router.address,
        BigNumber.from('54666805176129187'),
        ethers.utils.defaultAbiCoder.encode(['uint64'], [subscriptionId]),
      )
    const { gasUsed: transferAndCallTxGasUsed } = await transferAndCallTx.wait()

    // User sends request
    const requestTx = await contracts.client.sendSimpleRequestWithJavaScript(
      'function myFancyFunction(){return "woah, thats fancy"}',
      subscriptionId,
      ids.donId,
      20_000,
    )
    const { gasUsed: requestTxGasUsed, events } = await requestTx.wait()
    const requestId = getEventArg(events, 'RequestSent', 0)

    // DON's transmitter submits a response
    const response = stringToBytes('woah, thats fancy')
    const error = stringToBytes('')
    const report = ethers.utils.defaultAbiCoder.encode(
      ['bytes32[]', 'bytes[]', 'bytes[]'],
      [[ethers.utils.hexZeroPad(requestId, 32)], [response], [error]],
    )
    const fulfillmentTx = await contracts.coordinator.callReport(report)
    const { gasUsed: fulfillmentTxGasUsed } = await fulfillmentTx.wait()

    currentGasUsed = acceptTermsOfServiceGasUsed
      .add(createSubscriptionTxGasUsed)
      .add(addConsumerTxGasUsed)
      .add(transferAndCallTxGasUsed)
      .add(requestTxGasUsed)
      .add(fulfillmentTxGasUsed)
      .toNumber()
  })
})
