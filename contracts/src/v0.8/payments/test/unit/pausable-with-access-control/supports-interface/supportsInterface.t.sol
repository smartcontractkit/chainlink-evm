// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.26;

import {BaseUnitTest} from "../../BaseUnitTest.t.sol";

import {IAccessControl} from "@openzeppelin/contracts@5.0.2/access/IAccessControl.sol";
import {IAccessControlDefaultAdminRules} from
  "@openzeppelin/contracts@5.0.2/access/extensions/IAccessControlDefaultAdminRules.sol";
import {IAccessControlEnumerable} from "@openzeppelin/contracts@5.0.2/access/extensions/IAccessControlEnumerable.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract SupportsInterfaceUnitTests is BaseUnitTest {
  function test_supportsInterface() public performForAllContracts(CommonContracts.PAUSABLE_WITH_ACCESS_CONTROL) {
    assertTrue(IERC165(s_contractUnderTest).supportsInterface(type(IAccessControlEnumerable).interfaceId));
    assertTrue(IERC165(s_contractUnderTest).supportsInterface(type(IERC165).interfaceId));
    assertTrue(IERC165(s_contractUnderTest).supportsInterface(type(IAccessControlDefaultAdminRules).interfaceId));
    assertTrue(IERC165(s_contractUnderTest).supportsInterface(type(IAccessControl).interfaceId));
  }
}
