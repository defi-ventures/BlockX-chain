***System Requirements***
OS - ubuntu 18.04
Memory - 4 GB RAM
CPU - 2vCPU

***Public RPC details***

URL - https://testnet.blockxnet.com

Chain ID - 11

***Chain ID (testnet v1)***

tokn-11

***Steps to setup a validator for testnet v1***

1. Clone the repository

```bash
cd ~/
git clone https://github.com/defi-ventures/ethermint.git
```

2. Install the following for setup

```bash
apt get update
apt install make build-essential jq
wget https://golang.org/dl/go1.16.6.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

3. Change the value of KEY, CHAINID, MONIKER, MNEMONIC in validator_node_setup.sh before running the validator node setup.
You need to generate a mnemonic for the setup
```bash
cd ~/ethermint
./validator_node_setup.sh
```

4. Replace the genesis file in ~/.toknd/config/
```
{
   "genesis_time":"2021-09-23T20:36:21.028545148Z",
   "chain_id":"tokn-11",
   "consensus_params":{
      "block":{
         "max_bytes":"22020096",
         "max_gas":"-1",
         "time_iota_ms":"30000"
      },
      "evidence":{
         "max_age_num_blocks":"100000",
         "max_age_duration":"172800000000000"
      },
      "validator":{
         "pub_key_types":[
            "ed25519"
         ]
      }
   },
   "app_hash":"",
   "app_state":{
      "params":null,
      "upgrade":{
         
      },
      "supply":{
         "supply":[
            
         ]
      },
      "gov":{
         "starting_proposal_id":"1",
         "deposits":null,
         "votes":null,
         "proposals":null,
         "deposit_params":{
            "min_deposit":[
               {
                  "denom":"atokn",
                  "amount":"10000000"
               }
            ],
            "max_deposit_period":"172800000000000"
         },
         "voting_params":{
            "voting_period":"172800000000000"
         },
         "tally_params":{
            "quorum":"0.334000000000000000",
            "threshold":"0.500000000000000000",
            "veto":"0.334000000000000000"
         }
      },
      "evm":{
         "accounts":[
            
         ],
         "txs_logs":[
            
         ],
         "chain_config":{
            "homestead_block":"0",
            "dao_fork_block":"0",
            "dao_fork_support":true,
            "eip150_block":"0",
            "eip150_hash":"0x0000000000000000000000000000000000000000000000000000000000000000",
            "eip155_block":"0",
            "eip158_block":"0",
            "byzantium_block":"0",
            "constantinople_block":"0",
            "petersburg_block":"0",
            "istanbul_block":"0",
            "muir_glacier_block":"0",
            "yoloV2_block":"-1",
            "ewasm_block":"-1"
         },
         "params":{
            "evm_denom":"atokn",
            "enable_create":true,
            "enable_call":true,
            "extra_eips":null
         }
      },
      "bank":{
         "send_enabled":true
      },
      "slashing":{
         "params":{
            "signed_blocks_window":"100",
            "min_signed_per_window":"0.500000000000000000",
            "downtime_jail_duration":"600000000000",
            "slash_fraction_double_sign":"0.050000000000000000",
            "slash_fraction_downtime":"0.010000000000000000"
         },
         "signing_infos":{
            
         },
         "missed_blocks":{
            
         }
      },
      "staking":{
         "params":{
            "unbonding_time":"1814400000000000",
            "max_validators":100,
            "max_entries":7,
            "historical_entries":0,
            "bond_denom":"atokn"
         },
         "last_total_power":"0",
         "last_validator_powers":null,
         "validators":null,
         "delegations":null,
         "unbonding_delegations":null,
         "redelegations":null,
         "exported":false
      },
      "genutil":{
         "gentxs":[
            {
               "type":"cosmos-sdk/StdTx",
               "value":{
                  "msg":[
                     {
                        "type":"cosmos-sdk/MsgCreateValidator",
                        "value":{
                           "description":{
                              "moniker":"localtestnet-1",
                              "identity":"",
                              "website":"",
                              "security_contact":"",
                              "details":""
                           },
                           "commission":{
                              "rate":"0.100000000000000000",
                              "max_rate":"0.200000000000000000",
                              "max_change_rate":"0.010000000000000000"
                           },
                           "min_self_delegation":"1",
                           "delegator_address":"eth12fkhumzzmdjlph5dcj5hemzqlc663pf5k8md7y",
                           "validator_address":"ethvaloper12fkhumzzmdjlph5dcj5hemzqlc663pf5g49es6",
                           "pubkey":"ethvalconspub1zcjduepqwasywlrykyjtuvg4r3yve6zh0stg5tt43cvcy44e4lkz4mmpwgcs8jsm0x",
                           "value":{
                              "denom":"atokn",
                              "amount":"10000000000000000000000"
                           }
                        }
                     }
                  ],
                  "fee":{
                     "amount":[
                        
                     ],
                     "gas":"200000"
                  },
                  "signatures":[
                     {
                        "pub_key":{
                           "type":"ethermint/PubKeyEthSecp256k1",
                           "value":"A5ZbQPTCM4L6oSiHsIbUQf0dUguqR4r07DOtXxbtTvUn"
                        },
                        "signature":"8XBpNAtECSeg25p8k9Aze21LnhX2cz8bwavTz4oloLxFGCTDm/+MQEI4wNgP0SkQFDXupl3pbkzw2vvMy/oVbAA="
                     }
                  ],
                  "memo":"af7abd00b81255e96736190d249c9266bf6590ad@10.12.20.61:26656"
               }
            }
         ]
      },
      "mint":{
         "minter":{
            "inflation":"0.130000000000000000",
            "annual_provisions":"0.000000000000000000"
         },
         "params":{
            "mint_denom":"atokn",
            "inflation_rate_change":"0.130000000000000000",
            "inflation_max":"0.200000000000000000",
            "inflation_min":"0.070000000000000000",
            "goal_bonded":"0.670000000000000000",
            "blocks_per_year":"6311520"
         }
      },
      "evidence":{
         "params":{
            "max_evidence_age":"120000000000"
         },
         "evidence":[
            
         ]
      },
      "distribution":{
         "params":{
            "community_tax":"0.020000000000000000",
            "base_proposer_reward":"0.010000000000000000",
            "bonus_proposer_reward":"0.040000000000000000",
            "withdraw_addr_enabled":true
         },
         "fee_pool":{
            "community_pool":[
               
            ]
         },
         "delegator_withdraw_infos":[
            
         ],
         "previous_proposer":"",
         "outstanding_rewards":[
            
         ],
         "validator_accumulated_commissions":[
            
         ],
         "validator_historical_rewards":[
            
         ],
         "validator_current_rewards":[
            
         ],
         "delegator_starting_infos":[
            
         ],
         "validator_slash_events":[
            
         ]
      },
      "crisis":{
         "constant_fee":{
            "denom":"atokn",
            "amount":"1000"
         }
      },
      "auth":{
         "params":{
            "max_memo_characters":"256",
            "tx_sig_limit":"7",
            "tx_size_cost_per_byte":"10",
            "sig_verify_cost_ed25519":"590",
            "sig_verify_cost_secp256k1":"1000"
         },
         "accounts":[
            {
               "type":"ethermint/EthAccount",
               "value":{
                  "address":"eth12fkhumzzmdjlph5dcj5hemzqlc663pf5k8md7y",
                  "eth_address":"0x526d7e6c42DB65F0DE8dC4A97cEc40FE35A88534",
                  "coins":[
                     {
                        "denom":"atokn",
                        "amount":"1000000000000000000000000000"
                     }
                  ],
                  "public_key":"",
                  "account_number":0,
                  "sequence":0,
                  "code_hash":"c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
               }
            }
         ]
      }
   }
}
```

5. Add the following in seeds, persistent_peers in ~/.toknd/config/config.toml
```
af7abd00b81255e96736190d249c9266bf6590ad@52.71.20.235:26656,8d6c9fec090000627dba8e20ea3863dc05140663@54.166.134.59:26656
```

6. Reset the local chain config
```bash
cd ~/ethermint
./build/toknd unsafe-reset-all
```

7. Start local node and check if its syncing
```bash
./build/toknd start --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info" --trace
```

8. Start RPC (in a different terminal)
```bash
./build/tokncli rest-server --laddr "tcp://0.0.0.0:8545" --chain-id <> --trace --rpc-api eth,net,web3,personal --unsafe-cors
```

9. Acquire test tokens from the team for the address generated from the mnemonic

10. Run create validator command to become a validator in the network after the blockchain syncs completely(change values in commands accordingly).
Amount should be of the format - <x>atokn
```bash
./build/tokncli tx staking create-validator --amount=<> --pubkey=$(./build/toknd tendermint show-validator) --moniker=<> --chain-id=<> --commission-rate="0.10" --commission-max-rate="0.20" --commission-max-change-rate="0.01" --min-self-delegation="1" --gas="auto" --from=<>
```
