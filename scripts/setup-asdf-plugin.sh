#!/bin/bash

# Define the plugins and their respective repositories
declare -A plugins
plugins=(
  ["golang"]="https://github.com/asdf-community/asdf-golang.git"
  ["protoc"]="https://github.com/paxosglobal/asdf-protoc.git"
  ["starknet-foundry"]="https://github.com/foundry-rs/asdf-starknet-foundry.git"
)

# Read the .tool-versions file and install the plugins
while read -r line; do
  plugin=$(echo "$line" | awk '{print $1}')
  version=$(echo "$line" | awk '{print $2}')

  if [[ -n "${plugins[$plugin]}" ]]; then
    echo "Installing $plugin $version..."
    asdf plugin add "$plugin" "${plugins[$plugin]}"
  else
    echo "No repository URL found for plugin $plugin. Skipping..."
  fi
done < .tool-versions

echo "All plugins installed successfully."
