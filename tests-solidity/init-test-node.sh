#!/bin/bash

CHAINID="ethermint-1337"
MONIKER="localtestnet"

VAL_KEY="localkey"
VAL_MNEMONIC="gesture inject test cycle original hollow east ridge hen combine junk child bacon zero hope comfort vacuum milk pitch cage oppose unhappy lunar seat"

USER1_KEY="user1"
USER1_MNEMONIC="copper push brief egg scan entry inform record adjust fossil boss egg comic alien upon aspect dry avoid interest fury window hint race symptom"

USER2_KEY="user2"
USER2_MNEMONIC="maximum display century economy unlock van census kite error heart snow filter midnight usage egg venture cash kick motor survey drastic edge muffin visual"

# remove existing daemon and client
rm -rf ~/.ethermint*

tokncli config keyring-backend test

# Set up config for CLI
tokncli config chain-id $CHAINID
tokncli config output json
tokncli config indent true
tokncli config trust-node true

# Import keys from mnemonics
echo $VAL_MNEMONIC | tokncli keys add $VAL_KEY --recover
echo $USER1_MNEMONIC | tokncli keys add $USER1_KEY --recover
echo $USER2_MNEMONIC | tokncli keys add $USER2_KEY --recover

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
toknd init $MONIKER --chain-id $CHAINID

# Allocate genesis accounts (cosmos formatted addresses)
toknd add-genesis-account $(tokncli keys show $VAL_KEY -a) 1000000000000000000000atokn,10000000000000000stake
toknd add-genesis-account $(tokncli keys show $USER1_KEY -a) 1000000000000000000000atokn,10000000000000000stake
toknd add-genesis-account $(tokncli keys show $USER2_KEY -a) 1000000000000000000000atokn,10000000000000000stake

# Sign genesis transaction
toknd gentx --name $VAL_KEY --keyring-backend test

# Collect genesis tx
toknd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
toknd validate-genesis

# Command to run the rest server in a different terminal/window
echo -e '\nrun the following command in a different terminal/window to run the REST server and JSON-RPC:'
echo -e "tokncli rest-server --laddr \"tcp://localhost:8545\" --wsport 8546 --unlock-key $VAL_KEY,$USER1_KEY,$USER2_KEY --chain-id $CHAINID --trace\n"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
toknd start --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info" --trace
