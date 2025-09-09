/**
 * @description this script generates a master interface for interacting with the automation registry
 * @notice run this script with pnpm tsx ./scripts/generate-automation-master-interface-v2_3.ts
 */
import Registry from '../artifacts/src/v0.8/automation/v2_3/AutomationRegistry2_3.sol/AutomationRegistry2_3.json'
import RegistryLogicA from '../artifacts/src/v0.8/automation/v2_3/AutomationRegistryLogicA2_3.sol/AutomationRegistryLogicA2_3.json'
import RegistryLogicB from '../artifacts/src/v0.8/automation/v2_3/AutomationRegistryLogicB2_3.sol/AutomationRegistryLogicB2_3.json'
import RegistryLogicC from '../artifacts/src/v0.8/automation/v2_3/AutomationRegistryLogicC2_3.sol/AutomationRegistryLogicC2_3.json'

import fs from 'fs'
import { exec } from 'child_process'
import { createHash } from 'crypto'

const dest = 'src/v0.8/automation/interfaces/v2_3'
const srcDest = `${dest}/IAutomationRegistryMaster2_3.sol`
const tmpDest = `${dest}/tmp.txt`

const combinedABI = []
const abiSet = new Set()
const abis = [
  Registry.abi,
  RegistryLogicA.abi,
  RegistryLogicB.abi,
  RegistryLogicC.abi,
]

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
cat ${tmpDest} | pnpm abi-to-sol --solidity-version ^0.8.4 --license MIT > ${srcDest} IAutomationRegistryMaster2_3;
echo "// abi-checksum: ${checksum}" | cat - ${srcDest} > ${tmpDest} && mv ${tmpDest} ${srcDest};
export FOUNDRY_PROFILE=automation; forge fmt ${srcDest};
`

exec(cmd)

console.log('generated new master interface for automation registry')
