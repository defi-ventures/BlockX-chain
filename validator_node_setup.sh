#!/bin/bash

KEY="tokntestkey-1"
CHAINID="tokn-11"
MONIKER="localtestnet-1"
MNEMONIC=""

# remove existing daemon and client
rm -rf ~/.tokn*

make build

./build/tokncli config keyring-backend test

# Set up config for CLI
./build/tokncli config chain-id $CHAINID
./build/tokncli config output json
./build/tokncli config indent true
./build/tokncli config trust-node true

# Set moniker and chain-id for Tokn (Moniker can be anything, chain-id must be an integer)
./build/toknd init $MONIKER --chain-id $CHAINID

# if $KEY exists it should be deleted
echo $MNEMONIC | ./build/tokncli keys add $KEY --recover