#!/bin/bash

echo "📦 Installing dependencies..."
npm install ethers dotenv

echo -e "\n🔧 Running deployment script..."
node deploy-mock-keystone-forwarder.js