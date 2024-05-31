// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IReceiver} from "./interfaces/IReceiver.sol";
import {ConfirmedOwner} from "../shared/access/ConfirmedOwner.sol";

contract KeystoneFeedsConsumer is IReceiver, ConfirmedOwner {
  event MessageReceived(bytes10 indexed workflowName, address indexed workflowOwner, uint256 nReports);
  event FeedReceived(bytes32 indexed feedId, uint256 price, uint64 timestamp);

  error UnauthorizedSender(address sender);
  error UnauthorizedWorkflowOwner(address workflowOwner);
  error UnauthorizedWorkflowName(bytes10 workflowName);

  constructor() ConfirmedOwner(msg.sender) {}

  struct FeedReport {
    bytes32 FeedId;
    uint256 Price;
    uint64 Timestamp;
  }

  mapping(bytes32 feedId => uint256 price) internal s_prices;
  address[] internal allowedSenders;
  address[] internal allowedWorkflowOwners;
  bytes10[] internal allowedWorkflowNames;

  function setConfig(address[] calldata _allowedSenders, address[] calldata _allowedWorkflowOwners, bytes10[] calldata _allowedWorkflowNames) external onlyOwner {
    allowedSenders = _allowedSenders;
    allowedWorkflowOwners = _allowedWorkflowOwners;
    allowedWorkflowNames = _allowedWorkflowNames;
  }

  function onReport(bytes calldata metadata, bytes calldata rawReport) external {
    bool allowed = false;
    for (uint32 i = 0; i < allowedSenders.length; i++) {
      if (msg.sender == allowedSenders[i]) {
        allowed = true;
        break;
      }
    }
    if (!allowed) {
      revert UnauthorizedSender(msg.sender);
    }

    (bytes10 workflowName, address workflowOwner) = _getInfo(metadata);
    allowed = false;
    for (uint32 i = 0; i < allowedWorkflowNames.length; i++) {
      if (workflowName == allowedWorkflowNames[i]) {
        allowed = true;
        break;
      }
    }
    if (!allowed) {
      revert UnauthorizedWorkflowName(workflowName);
    }

    allowed = false;
    for (uint32 i = 0; i < allowedWorkflowOwners.length; i++) {
      if (workflowOwner == allowedWorkflowOwners[i]) {
        allowed = true;
        break;
      }
    }
    if (!allowed) {
      revert UnauthorizedWorkflowOwner(workflowOwner);
    }

    FeedReport[] memory feeds = abi.decode(rawReport, (FeedReport[]));
    for (uint32 i = 0; i < feeds.length; i++) {
      s_prices[feeds[i].FeedId] = feeds[i].Price;
      emit FeedReceived(feeds[i].FeedId, feeds[i].Price, feeds[i].Timestamp);
    }

    emit MessageReceived(workflowName, workflowOwner, feeds.length);
  }

  function _getInfo(
    bytes memory metadata
  ) internal pure returns (bytes10 workflowName, address workflowOwner) {
    // (first 32 bytes contain length of the byte array)
    // workflow_cid             // offset 32, size 32
    // workflow_name            // offset 64, size 10
    // workflow_owner           // offset 74, size 20
    // report_name              // offset 94, size  2
    assembly {
      // shift right by 22 bytes to get the actual value
      workflowName := shr(mul(22, 8), mload(add(metadata, 64)))
      // shift right by 12 bytes to get the actual value
      workflowOwner := shr(mul(12, 8), mload(add(metadata, 74)))
    }
  }

  function getPrice(bytes32 feedId) external view returns (uint256) {
    return s_prices[feedId];
  }
}
