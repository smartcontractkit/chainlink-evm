// SPDX-License-Identifier: BUSL 1.1
pragma solidity 0.8.26;

import {WorkflowRegistry} from "../../WorkflowRegistry.sol";
import {LinkingUtils} from "../../testhelpers/LinkingUtils.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable-next-line max-states-count
contract WorkflowRegistrySetup is Test {
  WorkflowRegistry internal s_registry;
  address internal s_owner;
  address internal s_stranger;
  address internal s_user;

  uint256 internal s_allowedSignerPrivateKey;
  address internal s_allowedSigner;
  uint256 internal s_validityTimestamp;
  bytes32 internal s_proof;

  string internal s_donFamily;
  string internal s_binaryUrl;
  string internal s_configUrl;
  string internal s_tag;
  string internal s_workflowName;
  bytes32 internal s_workflowId;
  bytes internal s_attributes;
  string internal s_invalidLongString;
  string internal s_invalidURL;

  function setUp() public virtual {
    s_owner = makeAddr("owner");
    s_stranger = makeAddr("stranger");
    s_allowedSignerPrivateKey = 0x200b7adf7bcce82338c9b5d8114629b511e4be583683449d90c60718739b683c;
    s_validityTimestamp = uint256(block.timestamp + 1 hours);
    s_proof = keccak256("test-proof");
    s_allowedSigner = vm.addr(s_allowedSignerPrivateKey);
    assertEq(s_allowedSigner, address(0x86f2cE81640Fd86e68CF3EB25c2801D6E1C62bd0));

    s_user = makeAddr("user");
    s_donFamily = "DON-A";
    s_binaryUrl = "ipfs://bin";
    s_configUrl = "ipfs://cfg";
    s_tag = "alpha";

    s_workflowName = "my-workflow";
    s_workflowId = keccak256("workflow1");
    s_attributes = hex"11223344556677889900aabbccddeeff";
    s_invalidLongString =
      "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcd";
    s_invalidURL =
      "https://www.example.com/this/is/a/very/long/url/that/keeps/going/on/and/on/to/ensure/that/it/exceeds/two/hundred/and/one/characters/in/length/for/testing/purposes/and/it/should/be/sufficiently/long/to/meet/your/requirements/for/this/test";

    vm.startPrank(s_owner);
    s_registry = new WorkflowRegistry();
    address[] memory signers = new address[](1);
    signers[0] = s_allowedSigner;
    s_registry.updateAllowedSigners(signers, true);
    vm.stopPrank();
  }

  function _setDONLimit() internal {
    vm.prank(s_owner);
    s_registry.setDONLimit(s_donFamily, 100, true);
  }

  // Helper to link an owner
  function _linkOwner(
    address owner
  ) internal {
    (bytes32 ownerProof, bytes memory sig) = _getLinkProofSignature(owner);
    vm.prank(owner);
    s_registry.linkOwner(s_validityTimestamp, ownerProof, sig);
  }

  function _getLinkProofSignature(
    address owner
  ) internal view returns (bytes32, bytes memory) {
    bytes32 ownerProof = keccak256(abi.encode(s_proof, owner));
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_LINK, address(s_registry), owner, s_validityTimestamp, ownerProof
      )
    );
    return (ownerProof, abi.encodePacked(r, s, v));
  }

  function _getUnlinkProofSignature(
    address owner
  ) internal view returns (bytes32, bytes memory) {
    bytes32 ownerProof = keccak256(abi.encode(s_proof, owner));
    (uint8 v, bytes32 r, bytes32 s) = vm.sign(
      s_allowedSignerPrivateKey,
      LinkingUtils.getMessageHash(
        LinkingUtils.REQUEST_TYPE_UNLINK, address(s_registry), owner, s_validityTimestamp, ownerProof
      )
    );
    return (ownerProof, abi.encodePacked(r, s, v));
  }

  // helper to upsert one test workflow
  function _upsertTestWorklow(WorkflowRegistry.WorkflowStatus status, bool keepAlive, address owner) internal {
    vm.prank(owner);
    s_registry.upsertWorkflow(
      s_workflowName, s_tag, s_workflowId, status, s_donFamily, s_binaryUrl, s_configUrl, s_attributes, keepAlive
    );
  }

  // helper to upsert 5 test workflows
  function _upsertTestWorklows(WorkflowRegistry.WorkflowStatus status, bool keepAlive, address owner) internal {
    // Workflow 1: Price Oracle
    bytes32 workflowId1 = keccak256("workflow1");
    string memory workflowName1 = "Price Oracle";
    string memory tag1 = "oracle-main";
    string memory binaryUrl1 = "https://example.com/binaries/price-oracle.wasm";
    string memory configUrl1 = "https://example.com/configs/price-oracle.json";
    bytes memory attributes1 = abi.encode("Price Oracle v1.0");

    vm.startPrank(owner);
    s_registry.upsertWorkflow(
      workflowName1, tag1, workflowId1, status, s_donFamily, binaryUrl1, configUrl1, attributes1, keepAlive
    );

    // Workflow 2: Weather Data Feeder
    bytes32 workflowId2 = keccak256("workflow2");
    string memory workflowName2 = "Weather Data Feeder";
    string memory tag2 = "weather-feed";
    string memory binaryUrl2 = "https://example.com/binaries/weather-data.wasm";
    string memory configUrl2 = "https://example.com/configs/weather-config.json";
    bytes memory attributes2 = abi.encode("Weather Data v2.1");

    s_registry.upsertWorkflow(
      workflowName2, tag2, workflowId2, status, s_donFamily, binaryUrl2, configUrl2, attributes2, keepAlive
    );

    // Workflow 3: NFT Metadata Service
    bytes32 workflowId3 = keccak256("workflow3");
    string memory workflowName3 = "NFT Metadata Service";
    string memory tag3 = "nft-meta";
    string memory binaryUrl3 = "https://example.com/binaries/nft-metadata.wasm";
    string memory configUrl3 = "https://example.com/configs/nft-settings.json";
    bytes memory attributes3 = abi.encode("NFT Metadata Service v1.2");

    s_registry.upsertWorkflow(
      workflowName3, tag3, workflowId3, status, s_donFamily, binaryUrl3, configUrl3, attributes3, keepAlive
    );

    // Workflow 4: Cross-Chain Bridge Monitor
    bytes32 workflowId4 = keccak256("workflow4");
    string memory workflowName4 = "Cross-Chain Bridge Monitor";
    string memory tag4 = "bridge-monitor";
    string memory binaryUrl4 = "https://example.com/binaries/bridge-monitor.wasm";
    string memory configUrl4 = "https://example.com/configs/bridge-config.json";
    bytes memory attributes4 = abi.encode("Bridge Monitor v3.0");

    s_registry.upsertWorkflow(
      workflowName4, tag4, workflowId4, status, s_donFamily, binaryUrl4, configUrl4, attributes4, keepAlive
    );

    // Workflow 5: Sports Data Feed
    bytes32 workflowId5 = keccak256("workflow5");
    string memory workflowName5 = "Sports Data Feed";
    string memory tag5 = "sports-feed";
    string memory binaryUrl5 = "https://example.com/binaries/sports-data.wasm";
    string memory configUrl5 = "https://example.com/configs/sports-config.json";
    bytes memory attributes5 = abi.encode("Sports Data Feed v1.5");

    s_registry.upsertWorkflow(
      workflowName5, tag5, workflowId5, status, s_donFamily, binaryUrl5, configUrl5, attributes5, keepAlive
    );

    vm.stopPrank();
  }
}
