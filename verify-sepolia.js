const fs = require("fs");
const { execSync } = require("child_process");

// Try to load dotenv if available
try {
  require("dotenv").config();
} catch (e) {
  // Read .env file manually if dotenv is not installed
  if (fs.existsSync(".env")) {
    const envConfig = fs.readFileSync(".env", "utf8");
    envConfig.split("\n").forEach(line => {
      const [key, value] = line.split("=");
      if (key && value) {
        process.env[key.trim()] = value.trim();
      }
    });
  }
}

async function main() {
  // Check for deployment info
  if (!fs.existsSync("deployment.json")) {
    console.error("deployment.json not found. Please deploy the contract first.");
    process.exit(1);
  }

  const deployment = JSON.parse(fs.readFileSync("deployment.json", "utf8"));
  const contractAddress = deployment.address;

  console.log("Verifying MockKeystoneForwarder on Sepolia");
  console.log("Contract Address:", contractAddress);

  // Check for Etherscan API key
  if (!process.env.ETHERSCAN_API_KEY) {
    console.error("\nMissing ETHERSCAN_API_KEY in .env file");
    console.error("Add to your .env file:");
    console.error("ETHERSCAN_API_KEY=your-etherscan-api-key");
    console.error("\nGet your API key from: https://etherscan.io/apis");
    process.exit(1);
  }

  try {
    // Add delay to avoid rate limit
    console.log("\nWaiting 3 seconds to avoid rate limit...");
    await new Promise(resolve => setTimeout(resolve, 3000));
    
    const cmd = `cd contracts && forge verify-contract \
      --chain sepolia \
      --etherscan-api-key ${process.env.ETHERSCAN_API_KEY} \
      --watch \
      --retry 3 \
      ${contractAddress} \
      src/v0.8/keystone/MockKeystoneForwarder.sol:MockKeystoneForwarder`;

    console.log("\nRunning verification...");
    execSync(cmd, { stdio: 'inherit' });
    
    console.log("\n✅ Contract verified successfully!");
    console.log(`View on Sepolia Etherscan: https://sepolia.etherscan.io/address/${contractAddress}#code`);
    
  } catch (error) {
    console.error("\nVerification failed. You can verify manually at:");
    console.error(`https://sepolia.etherscan.io/verifyContract`);
    console.error("\nContract details:");
    console.error("- Address:", contractAddress);
    console.error("- Compiler: v0.8.19");
    console.error("- Optimization: Enabled (1000000 runs)");
    console.error("- Contract name: MockKeystoneForwarder");
  }
}

main().catch(error => {
  console.error("\nError:", error);
  process.exit(1);
});