#!/bin/bash

# Ensure script is run with sudo
if [ "$EUID" -ne 0 ]; then
  echo "The root the root the root is on fire, please run via sudo like the README.md says!"
  exit
fi

echo "Updating packages..."
apt-get update

echo "Adding Ethereum PPA..."
apt-get install -y software-properties-common
add-apt-repository -y ppa:ethereum/ethereum

echo "Installing Solidity Compiler"
apt-get update
apt-get install -y solc

solc --version
echo "Installation complete."