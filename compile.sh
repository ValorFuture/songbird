#!/bin/bash
if [[ $(pwd) =~ *" "* ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
if [ -z ${GOPATH+x} ]; then echo "GOPATH is not set, visit https://github.com/golang/go/wiki/SettingGOPATH" && exit; fi
if [[ $(go version) != *"go1.15"* ]]; then echo "Go version is not go1.15" && exit; fi
if [ "$(uname -m)" != "x86_64" ]; then echo "Machine architecture is not x86_64" && exit; fi

WORKING_DIR=$(pwd)

#GENESIS_FILE=genesis_coston.go
#GENESIS_FILE=genesis_scdev_160k.go

if [ -z "$GENESIS_FILE" ]; then
   GENESIS_FILE=genesis_scdev_airdrop.go
fi

echo "Using genesis file '$GENESIS_FILE'" 

if [ $# -ne 0 ]
  then
    GENESIS_FILE=$1
fi

sudo rm -rf $GOPATH/src/github.com/ava-labs
sudo rm -rf $GOPATH/pkg/mod/github.com/ava-labs
go get -v -d github.com/ava-labs/avalanchego/...
cd $GOPATH/src/github.com/ava-labs/avalanchego
# Hard-coded commit to tag v1.4.12, at the time of this authoring
# https://github.com/ava-labs/avalanchego/releases/tag/v1.4.12
git checkout cae93d95c1bcdc02e1370d38ed1c9d87f1c8c814

echo "Applying Flare-specific changes to AvalancheGo..."

# Apply changes to avalanchego
cp $WORKING_DIR/src/genesis/$GENESIS_FILE ./genesis/genesis_coston.go
cp $WORKING_DIR/src/avalanchego/flags.go ./config/flags.go
cp $WORKING_DIR/src/avalanchego/beacons.go ./genesis/beacons.go
cp $WORKING_DIR/src/avalanchego/genesis_fuji.go ./genesis/genesis_fuji.go
cp $WORKING_DIR/src/avalanchego/unparsed_config.go ./genesis/unparsed_config.go
cp $WORKING_DIR/src/avalanchego/set.go ./snow/validators/set.go
cp $WORKING_DIR/src/avalanchego/build_coreth.sh ./scripts/build_coreth.sh
mkdir ./scripts/coreth_changes
cp $WORKING_DIR/src/coreth/state_transition.go ./scripts/coreth_changes/state_transition.go
cp $WORKING_DIR/src/stateco/state_connector.go ./scripts/coreth_changes/state_connector.go
cp $WORKING_DIR/src/keeper/keeper.go ./scripts/coreth_changes/keeper.go
cp $WORKING_DIR/src/keeper/keeper_test.go ./scripts/coreth_changes/keeper_test.go

export ROCKSDBALLOWED=1
./scripts/build.sh
rm -rf ./scripts/coreth_changes