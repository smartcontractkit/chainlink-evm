const fs = require("fs");
const path = require("path");
const { execSync } = require("child_process");

// Try to load dotenv if available
try {
  require("dotenv").config();
} catch (e) {
  console.log("dotenv not installed, reading .env manually...");
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

// Load ethers dynamically
let ethers;
try {
  ethers = require("ethers");
} catch (e) {
  console.error("ethers.js not found. Installing...");
  execSync("npm install ethers@6", { stdio: 'inherit' });
  ethers = require("ethers");
}

async function main() {
  // Check environment variables
  if (!process.env.RPC_URL || !process.env.PRIVATE_KEY) {
    console.error("Missing RPC_URL or PRIVATE_KEY in .env file");
    console.error("\nCreate a .env file with:");
    console.error("RPC_URL=https://your-rpc-url");
    console.error("PRIVATE_KEY=your-private-key");
    process.exit(1);
  }

  console.log("Building MockKeystoneForwarder contract...");
  
  try {
    // Build only the MockKeystoneForwarder contract
    execSync("cd contracts && forge build --contracts src/v0.8/keystone/MockKeystoneForwarder.sol", { stdio: 'inherit' });
  } catch (error) {
    console.error("Forge build failed. Trying direct compilation with solc...");
    
    // Try direct solc compilation
    const contractPath = path.join(__dirname, "contracts/src/v0.8/keystone/MockKeystoneForwarder.sol");
    const buildDir = path.join(__dirname, "build");
    
    if (!fs.existsSync(buildDir)) {
      fs.mkdirSync(buildDir);
    }
    
    try {
      execSync(`cd contracts && solc --base-path . --include-path ./src --abi --bin --overwrite -o ../build src/v0.8/keystone/MockKeystoneForwarder.sol`);
      
      const abi = JSON.parse(fs.readFileSync(path.join(buildDir, "MockKeystoneForwarder.abi"), "utf8"));
      const bytecode = "0x" + fs.readFileSync(path.join(buildDir, "MockKeystoneForwarder.bin"), "utf8").trim();
      
      return await deployContract(abi, bytecode);
    } catch (solcError) {
      console.error("Direct compilation also failed:", solcError.message);
      process.exit(1);
    }
  }

  // Read the compiled artifact from forge
  const artifactPath = path.join(__dirname, "contracts/foundry-artifacts/MockKeystoneForwarder.sol/MockKeystoneForwarder.json");
  
  if (!fs.existsSync(artifactPath)) {
    console.error("Compiled artifact not found at:", artifactPath);
    // Try alternative path
    const altPath = path.join(__dirname, "foundry-artifacts/MockKeystoneForwarder.sol/MockKeystoneForwarder.json");
    if (fs.existsSync(altPath)) {
      const artifact = JSON.parse(fs.readFileSync(altPath, "utf8"));
      const abi = artifact.abi;
      const bytecode = artifact.bytecode.object;
      return await deployContract(abi, bytecode);
    }
    process.exit(1);
  }

  const artifact = JSON.parse(fs.readFileSync(artifactPath, "utf8"));
  const abi = artifact.abi;
  const bytecode = artifact.bytecode.object;
  
  await deployContract(abi, bytecode);
}

async function deployContract(abi, bytecode) {
  // Setup provider and wallet
  const provider = new ethers.JsonRpcProvider(process.env.RPC_URL);
  const wallet = new ethers.Wallet(process.env.PRIVATE_KEY, provider);

  console.log("\nDeploying from address:", wallet.address);
  
  // Check balance
  const balance = await provider.getBalance(wallet.address);
  console.log("Balance:", ethers.formatEther(balance), "ETH");

  if (balance === 0n) {
    console.error("Insufficient balance for deployment!");
    process.exit(1);
  }

  // Deploy contract
  const factory = new ethers.ContractFactory(abi, bytecode, wallet);
  console.log("\nDeploying MockKeystoneForwarder...");
  
  const contract = await factory.deploy();
  const tx = contract.deploymentTransaction();
  
  console.log("Transaction hash:", tx.hash);
  console.log("Waiting for confirmation...");
  
  await contract.waitForDeployment();
  const address = await contract.getAddress();
  
  console.log("\n✅ MockKeystoneForwarder deployed to:", address);

  // Verify deployment
  const typeAndVersion = await contract.typeAndVersion();
  console.log("Type and Version:", typeAndVersion);

  // Save deployment info
  const deploymentInfo = {
    address: address,
    deployer: wallet.address,
    transactionHash: tx.hash,
    timestamp: new Date().toISOString(),
    network: {
      chainId: Number((await provider.getNetwork()).chainId),
      rpcUrl: process.env.RPC_URL
    }
  };

  fs.writeFileSync("deployment.json", JSON.stringify(deploymentInfo, null, 2));
  console.log("\nDeployment info saved to deployment.json");
}

main().catch(error => {
  console.error("\nDeployment failed:", error);
  process.exit(1);
});