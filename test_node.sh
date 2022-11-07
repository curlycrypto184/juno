# Ensure juno is installed first.

set -e

KEY="juno1"
CHAINID="juno-t1"
MONIKER="localjuno"
KEYALGO="secp256k1"
KEYRING="test" # export juno_KEYRING="TEST"
LOGLjunoL="info"
TRACE="" # "--trace"
APP_HOME="$HOME/.juno"


junod config keyring-backend $KEYRING
junod config chain-id $CHAINID
junod config output "json"

if [ -z "$APP_HOME" ]; then
  echo "APP_HOME is not set to a valid dir...   "
  exit 1
fi
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

from_scratch () {

  make install

  # remove existing daemon
  rm -rf ~/.juno/*

  # if $KEY exists it should be deleted
  # decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry
  # juno1hj5fveer5cjtn4wd6wstzugjfdxzl0xps73ftl
  echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | junod keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover | true
  junod init $MONIKER --chain-id $CHAINID 

  # Function updates the config based on a jq argument as a string
  update_test_genesis () {
    # update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
    cat $APP_HOME/config/genesis.json | jq "$1" > $APP_HOME/config/tmp_genesis.json && mv $APP_HOME/config/tmp_genesis.json $APP_HOME/config/genesis.json
  }

  # Set gas limit in genesis
  update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
  update_test_genesis '.app_state["gov"]["voting_params"]["voting_period"]="15s"'  
  update_test_genesis '.app_state["staking"]["params"]["bond_denom"]="ujuno"'  


  update_test_genesis '.app_state["mint"]["params"]["mint_denom"]="ujuno"'  
  update_test_genesis '.app_state["gov"]["deposit_params"]["min_deposit"]=[{"denom": "ujuno","amount": "1000000"}]' # 1 juno right now
  update_test_genesis '.app_state["crisis"]["constant_fee"]={"denom": "ujuno","amount": "1000"}'


  # Allocate genesis accounts
  # 10 juno (1 of which is used for validator)
  junod add-genesis-account $KEY 10000000ujuno --keyring-backend $KEYRING

  # create gentx with 1 juno
  junod gentx $KEY 1000000ujuno --keyring-backend $KEYRING --chain-id $CHAINID

  # Collect genesis tx
  junod collect-gentxs

  # Run this to ensure junorything worked and that the genesis file is setup correctly
  junod validate-genesis
}

from_scratch

# Opens the RPC endpoint to outside connections
sed -i '/laddr = "tcp:\/\/127.0.0.1:26657"/c\laddr = "tcp:\/\/0.0.0.0:26657"' $APP_HOME/config/config.toml
sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' $APP_HOME/config/config.toml
# cors_allowed_origins = []

# # Start the node (remove the --pruning=nothing flag if historical queries are not needed)
junod start --pruning=nothing  --minimum-gas-prices=0ujuno #--mode validator     