#!/bin/bash
if [[ $(pwd) =~ *" "* ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
if [ -z ${GOPATH+x} ]; then echo "GOPATH is not set, visit https://github.com/golang/go/wiki/SettingGOPATH" && exit; fi
if [[ $(go version) != *"go1.15"* ]]; then echo "Go version is not go1.15" && exit; fi
if [ "$(uname -m)" != "x86_64" ]; then echo "Machine architecture is not x86_64" && exit; fi

WORKING_DIR=$(pwd)

if [ -z "$GENESIS_FILE" ]; then
   GENESIS_FILE=genesis_coston.go
fi

if [ $# -ne 0 ]
  then
    GENESIS_FILE=$1
fi

echo "Using genesis file '$GENESIS_FILE'" 

# Start fresh
sudo rm -rf $GOPATH/src/github.com/ava-labs
sudo rm -rf $GOPATH/pkg/mod/github.com/ava-labs
sudo rm -rf $WORKING_DIR/tmp

# Get Avalanchego source
go get -v -d github.com/ava-labs/avalanchego/...
cd $GOPATH/src/github.com/ava-labs/avalanchego

# Switch to supported version
# Hard-coded commit to tag v1.4.12, at the time of this authoring
# https://github.com/ava-labs/avalanchego/releases/tag/v1.4.12
git checkout cae93d95c1bcdc02e1370d38ed1c9d87f1c8c814

echo "Applying Flare-specific changes to AvalancheGo..."

# Apply changes to avalanchego

# copy active genesis file
cp $WORKING_DIR/src/genesis/$GENESIS_FILE ./genesis/genesis_coston.go

# copy flare avalanchego changes
cp -R $WORKING_DIR/src/avalanchego/. .

# copy flare coreth changes
mkdir ./scripts/coreth_changes
cp -R $WORKING_DIR/src/coreth/. ./scripts/coreth_changes

# copy flare gRPC changes
echo "Update gRPC..."

export GO111MODULE=on
go get google.golang.org/grpc@v1.37.0
export GO111MODULE=

cp -R $WORKING_DIR/src/grpc@v1.37.0/. $GOPATH/pkg/mod/google.golang.org/grpc@v1.37.0
cd $GOPATH/pkg/mod/google.golang.org/grpc@v1.37.0

# build
make build

cd $GOPATH/src/github.com/ava-labs/avalanchego

export ROCKSDBALLOWED=true
./scripts/build.sh
rm -rf ./scripts/coreth_changes
rm -rf ./scripts/grpc_changes