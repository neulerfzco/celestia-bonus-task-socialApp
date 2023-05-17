## Setup guide to deploy the rollup on blockspacerace network

### celestia incentivized testnet: blockspacerace rollup: 
In order to deploy it on the blockspacerace as a rollup, please make the following modifications:  

#### 1. Settings on this repo 
Execute the following commands in the top directory of this repository
`
go mod edit -replace github.com/cosmos/cosmos-sdk=github.com/rollkit/cosmos-sdk@v0.46.7-rollkit-v0.7.3-no-fraud-proofs
go mod edit -replace github.com/tendermint/tendermint=github.com/celestiaorg/tendermint@v0.34.22-0.20221202214355-3605c597500d
go mod tidy
go mod download
`
#### 2. Deploy a Light node

`
celestia light start --core.ip https://rpc-blockspacerace.pops.one --gateway --gateway.addr 127.0.0.1 --gateway.port 26659 --p2p.network blockspacerace
`

#### 3. Start the rollup 
Execute the following commands in the top directory of this repository

```
VALIDATOR_NAME=social-validator1
CHAIN_ID=socialapp
KEY_NAME=socialapp-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

NAMESPACE_ID=$(openssl rand -hex 8)
echo $NAMESPACE_ID
DA_BLOCK_HEIGHT=$(curl https://rpc-blockspacerace.pops.one/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT
ignite chain build
socialappd tendermint unsafe-reset-all
socialappd init $VALIDATOR_NAME --chain-id $CHAIN_ID

socialappd keys add $KEY_NAME --keyring-backend test
socialappd add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
socialappd gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
socialappd collect-gentxs
socialappd start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":200000,"gas_limit":8000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT --rollkit.lazy_aggregator
```
