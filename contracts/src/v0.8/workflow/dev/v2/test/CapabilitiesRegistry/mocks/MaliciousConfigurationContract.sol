// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {CapabilitiesRegistry} from "../../../CapabilitiesRegistry.sol";
import {ICapabilityConfiguration} from "../../../interfaces/ICapabilityConfiguration.sol";

import {Constants} from "../Constants.t.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

contract MaliciousConfigurationContract is ICapabilityConfiguration, IERC165, Constants {
  string internal s_capabilityWithConfigurationContractId;

  constructor(
    string memory capabilityWithConfigContractId
  ) {
    s_capabilityWithConfigurationContractId = capabilityWithConfigContractId;
  }

  function getCapabilityConfiguration(
    uint32
  ) external pure returns (bytes memory configuration) {
    return bytes("");
  }

  function beforeCapabilityConfigSet(bytes32[] calldata, bytes calldata, uint64, uint32) external {
    CapabilitiesRegistry.NodeParams[] memory nodes = new CapabilitiesRegistry.NodeParams[](2);
    string[] memory capabilityIds = new string[](1);

    capabilityIds[0] = s_capabilityWithConfigurationContractId;

    // Set node one's signer to another address
    nodes[0] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID,
      signer: NODE_OPERATOR_ONE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY,
      csaKey: TEST_CSA_KEY,
      capabilityIds: capabilityIds
    });

    nodes[1] = CapabilitiesRegistry.NodeParams({
      nodeOperatorId: TEST_NODE_OPERATOR_ONE_ID,
      p2pId: P2P_ID_THREE,
      signer: NODE_OPERATOR_THREE_SIGNER_ADDRESS,
      encryptionPublicKey: TEST_ENCRYPTION_PUBLIC_KEY_THREE,
      csaKey: TEST_CSA_KEY_THREE,
      capabilityIds: capabilityIds
    });

    CapabilitiesRegistry(msg.sender).updateNodes(nodes);
  }

  function supportsInterface(
    bytes4 interfaceId
  ) public pure returns (bool) {
    return interfaceId == type(ICapabilityConfiguration).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
