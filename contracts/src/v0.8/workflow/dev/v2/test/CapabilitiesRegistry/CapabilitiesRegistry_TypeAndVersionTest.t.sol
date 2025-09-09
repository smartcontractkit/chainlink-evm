// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {BaseTest} from "./BaseTest.t.sol";

contract CapabilitiesRegistry_TypeAndVersionTest is BaseTest {
  function test_TypeAndVersion() public view {
    assertEq(s_CapabilitiesRegistry.typeAndVersion(), "CapabilitiesRegistry 2.0.0-dev");
  }
}
