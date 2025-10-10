// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IReceiver} from "../../keystone/interfaces/IReceiver.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/interfaces/IERC165.sol";

// This contract is used to trigger log trigger for testing evm chain capabilities
contract LoggerTester is IReceiver {
    // --- Roles ---
    address public immutable deployer;
    mapping(address => bool) public owners;
    mapping(address => bool) public allowedEmitters;

    // --- State ---
    uint256 public logCounter;

    // --- Events ---
    // id = shared incremental id; manual = true if via emitLog(), false if via onReport()
    event LogEmitted(uint256 id, bool manual);

    // --- Errors ---
    error NotOwner(address caller);
    error UnauthorizedEmitter(address caller);

    modifier onlyOwner() {
        if (!owners[msg.sender]) revert NotOwner(msg.sender);
        _;
    }

    constructor() {
        deployer = msg.sender;
        owners[msg.sender] = true;
    }

    // --- ERC165 ---
    function supportsInterface(
        bytes4 interfaceId
    ) public pure override returns (bool) {
        return interfaceId == type(IReceiver).interfaceId ||
            interfaceId == type(IERC165).interfaceId;
    }

    // --- Owner Management ---
    function setOwner(address user, bool isOwner) external onlyOwner {
        owners[user] = isOwner;
    }

    function setOwners(address[] calldata users, bool isOwner) external onlyOwner {
        for (uint256 i = 0; i < users.length; ++i) {
            owners[users[i]] = isOwner;
        }
    }

    // --- Emitters ---
    function setAllowedEmitter(address user, bool isAllowed) external onlyOwner {
        allowedEmitters[user] = isAllowed;
    }

    function setAllowedEmitters(address[] calldata users, bool isAllowed) external onlyOwner {
        for (uint256 i = 0; i < users.length; ++i) {
            allowedEmitters[users[i]] = isAllowed;
        }
    }

    // --- User Functions ---
    function emitLog() external {
        if (!allowedEmitters[msg.sender]) revert UnauthorizedEmitter(msg.sender);
        logCounter++;
        emit LogEmitted(logCounter, true); // manual
    }

    // IReceiver
    function onReport(bytes calldata /*metadata*/, bytes calldata /*report*/) external override {
        if (!allowedEmitters[msg.sender]) revert UnauthorizedEmitter(msg.sender);
        logCounter++;
        emit LogEmitted(logCounter, false); // via onReport
    }
}
