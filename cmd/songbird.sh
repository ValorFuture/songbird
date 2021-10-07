#!/bin/bash

# (c) 2021, Flare Networks Limited. All rights reserved.
# Please see the file LICENSE for licensing terms.

if [[ $(pwd) =~ " " ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
printf "\x1b[34mSongbird Canary Network Deployment\x1b[0m\n\n"

LAUNCH_DIR=$(pwd)

DB_TYPE=rocksdb
if [ "$(uname)" != "Linux" ] || [ "$(uname -m)" != "x86_64" ]; then DB_TYPE=leveldb; fi

# Test and export underlying chain APIs you chose to use for the state connector
source ./conf/export_chain_apis.sh $LAUNCH_DIR/conf/songbird/chain_apis.json

export FBA_VALs=$LAUNCH_DIR/conf/songbird/fba_validators.json
AVALANCHE_DIR=$LAUNCH_DIR/avalanchego
cd $AVALANCHE_DIR

# NODE 1
printf "Launching Songbird Node at 127.0.0.1:9650\n"
export WEB3_API=enabled
nohup ./build/avalanchego \
--http-host= \
--public-ip=127.0.0.1 \
--http-port=9650 \
--staking-port=9651 \
--log-dir=$LAUNCH_DIR/logs/songbird \
--db-dir=$LAUNCH_DIR/db/songbird \
--bootstrap-ips="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeIP" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.ip")" \
--bootstrap-ids="$(curl -m 10 -sX POST --data '{ "jsonrpc":"2.0", "id":1, "method":"info.getNodeID" }' -H 'content-type:application/json;' https://songbird.flare.network/ext/info | jq -r ".result.nodeID")" \
--db-type=$DB_TYPE \
--log-level=debug > /dev/null 2>&1 &
NODE_PID=`echo $!`

printf "Songbird node successfully launched on PID: ${NODE_PID}"
