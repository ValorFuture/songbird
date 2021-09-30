# (c) 2021, Flare Networks Limited. All rights reserved.
# Please see the file LICENSE for licensing terms.

#!/bin/bash
if [[ $(pwd) =~ " " ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
if [ -z ${GOPATH+x} ]; then echo "GOPATH is not set, visit https://github.com/golang/go/wiki/SettingGOPATH" && exit; fi
if [ "$(uname -m)" != "x86_64" ]; then echo "Machine architecture is not x86_64" && exit; fi

WORKING_DIR=$(pwd)
AVALANCHE_DIR=$WORKING_DIR/avalanchego

if [ ! -d $AVALANCHE_DIR ]; then
  echo "Cloning Avalanche repository..."
  git clone https://github.com/ava-labs/avalanchego $AVALANCHE_DIR --quiet
fi

echo "Checking out correct Avalanche version..."
cd $AVALANCHE_DIR
# Hard-coded commit to tag v1.5.2, at the time of this authoring
# https://github.com/ava-labs/avalanchego/releases/tag/v1.5.2
git checkout master --force --quiet
git pull --quiet
git checkout f2e51d790430a171e6d39f72911d98f134942a55 --quiet

GENESIS_FILE=genesis_local.go
if [ $# -ne 0 ]
  then
    GENESIS_FILE=genesis_$1.go
fi
echo "Using ${GENESIS_FILE} as genesis file..."

echo "Applying Flare-specific changes to Avalanche..."
cp $WORKING_DIR/src/genesis/$GENESIS_FILE $AVALANCHE_DIR/genesis/genesis_testnet.go
cp $WORKING_DIR/src/avalanchego/flags.go $AVALANCHE_DIR/config/flags.go
cp $WORKING_DIR/src/avalanchego/genesis.go $AVALANCHE_DIR/genesis/genesis.go
cp $WORKING_DIR/src/avalanchego/beacons.go $AVALANCHE_DIR/genesis/beacons.go
cp $WORKING_DIR/src/avalanchego/genesis_fuji.go $AVALANCHE_DIR/genesis/genesis_fuji.go
cp $WORKING_DIR/src/avalanchego/unparsed_config.go $AVALANCHE_DIR/genesis/unparsed_config.go
cp $WORKING_DIR/src/avalanchego/node.go $AVALANCHE_DIR/node/node.go
cp $WORKING_DIR/src/avalanchego/vm.go $AVALANCHE_DIR/vms/platformvm/vm.go
cp $WORKING_DIR/src/avalanchego/set.go $AVALANCHE_DIR/snow/validators/set.go
cp $WORKING_DIR/src/avalanchego/build_coreth.sh $AVALANCHE_DIR/scripts/build_coreth.sh
cp $WORKING_DIR/src/avalanchego/versions.sh $AVALANCHE_DIR/scripts/versions.sh
cp $WORKING_DIR/src/avalanchego/constants.sh $AVALANCHE_DIR/scripts/constants.sh

echo "Calling Avalanche build script..."
export ROCKSDBALLOWED=1
$AVALANCHE_DIR/scripts/build.sh
