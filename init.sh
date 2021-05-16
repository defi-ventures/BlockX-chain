#!/bin/bash

KEY="tokntestkey"
CHAINID="tokn-1"
MONIKER="localtestnet"

# remove existing daemon and client
rm -rf ~/.tokn*

make install

./build/tokncli config keyring-backend test

# Set up config for CLI
./build/tokncli config chain-id $CHAINID
./build/tokncli config output json
./build/tokncli config indent true
./build/tokncli config trust-node true

# if $KEY exists it should be deleted
./build/tokncli keys add $KEY

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
./build/toknd init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to atokn
cat $HOME/.toknd/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="atokn"' > $HOME/.toknd/config/tmp_genesis.json && mv $HOME/.toknd/config/tmp_genesis.json $HOME/.toknd/config/genesis.json
cat $HOME/.toknd/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="atokn"' > $HOME/.toknd/config/tmp_genesis.json && mv $HOME/.toknd/config/tmp_genesis.json $HOME/.toknd/config/genesis.json
cat $HOME/.toknd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="atokn"' > $HOME/.toknd/config/tmp_genesis.json && mv $HOME/.toknd/config/tmp_genesis.json $HOME/.toknd/config/genesis.json
cat $HOME/.toknd/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="atokn"' > $HOME/.toknd/config/tmp_genesis.json && mv $HOME/.toknd/config/tmp_genesis.json $HOME/.toknd/config/genesis.json

# increase block time (?)
cat $HOME/.toknd/config/genesis.json | jq '.consensus_params["block"]["time_iota_ms"]="30000"' > $HOME/.toknd/config/tmp_genesis.json && mv $HOME/.toknd/config/tmp_genesis.json $HOME/.toknd/config/genesis.json

if [[ $1 == "pending" ]]; then
    echo "pending mode on; block times will be set to 30s."
    # sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.toknd/config/config.toml
    sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.toknd/config/config.toml
    sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.toknd/config/config.toml
    sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.toknd/config/config.toml
    sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.toknd/config/config.toml
    sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.toknd/config/config.toml
    sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.toknd/config/config.toml
    sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.toknd/config/config.toml
fi

# Allocate genesis accounts (cosmos formatted addresses)
./build/toknd add-genesis-account $(./build/tokncli keys show $KEY -a) 100000000000000000000atokn

# Sign genesis transaction
./build/toknd gentx --name $KEY --amount=1000000000000000000atokn --keyring-backend test

# Collect genesis tx
./build/toknd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
./build/toknd validate-genesis

# Command to run the rest server in a different terminal/window
echo -e '\nrun the following command in a different terminal/window to run the REST server and JSON-RPC:'
echo -e "./build/tokncli rest-server --laddr \"tcp://localhost:8545\" --unlock-key $KEY --chain-id $CHAINID --trace --rpc-api web3,eth,debug,personal,net\n"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
./build/toknd start --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info" --trace
