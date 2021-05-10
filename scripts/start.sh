#!/bin/sh
toknd --home /ethermint/node$ID/toknd/ start > toknd.log &
sleep 5
tokncli rest-server --laddr "tcp://localhost:8545" --chain-id "ethermint-7305661614933169792" --trace --rpc-api="web3,eth,net,personal" > tokncli.log &
tail -f /dev/null