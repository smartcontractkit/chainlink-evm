/**
 * @description this script generates a master interface for interacting with the automation registry
 * @notice run this script with pnpm tsx ./scripts/generate-automation-master-interface.ts
 */
import Registry from '../artifacts/src/v0.8/automation/v2_2/AutomationRegistry2_2.sol/AutomationRegistry2_2.json'
import RegistryLogicA from '../artifacts/src/v0.8/automation/v2_2/AutomationRegistryLogicA2_2.sol/AutomationRegistryLogicA2_2.json'
import RegistryLogicB from '../artifacts/src/v0.8/automation/v2_2/AutomationRegistryLogicB2_2.sol/AutomationRegistryLogicB2_2.json'

import fs from 'fs'
import { exec } from 'child_process'
import { createHash } from 'crypto'

const dest = 'src/v0.8/automation/interfaces/v2_2'
const srcDest = `${dest}/IAutomationRegistryMaster.sol`
const tmpDest = `${dest}/tmp.txt`

const combinedABI = []
const abiSet = new Set()
const abis = [Registry.abi, RegistryLogicA.abi, RegistryLogicB.abi]

for (const abi of abis) {
  for (const entry of abi) {
    const id = createHash('sha256').update(JSON.stringify(entry)).digest('hex')
    if (!abiSet.has(id)) {
      abiSet.add(id)
      if (
        entry.type === 'function' &&
        (entry.name === 'checkUpkeep' ||
          entry.name === 'checkCallback' ||
          entry.name === 'simulatePerformUpkeep')
      ) {
        entry.stateMutability = 'view' // override stateMutability for check / callback / simulate functions
      }
      combinedABI.push(entry)
    }
  }
}

const checksum = createHash('sha256').update(JSON.stringify(abis)).digest('hex')

fs.writeFileSync(`${tmpDest}`, JSON.stringify(combinedABI))

const cmd = `
cat ${tmpDest} | pnpm abi-to-sol --solidity-version ^0.8.4 --license MIT > ${srcDest} IAutomationRegistryMaster;
echo "// abi-checksum: ${checksum}" | cat - ${srcDest} > ${tmpDest} && mv ${tmpDest} ${srcDest};
export FOUNDRY_PROFILE=automation; forge fmt ${srcDest};
`

exec(cmd)

console.log('generated new master interface for automation registry')
