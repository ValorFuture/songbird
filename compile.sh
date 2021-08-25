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

# Make sure permissions are set so we can fiddle with ava and grpc
chmod -R 775 $GOPATH/src/github.com/ava-labs
chmod -R 775 $GOPATH/pkg/mod/github.com/ava-labs
chmod -R 755 $GOPATH/pkg/mod/google.golang.org/grpc@v1.37.0

# Start fresh
rm -rf $GOPATH/src/github.com/ava-labs
rm -rf $GOPATH/pkg/mod/github.com/ava-labs
rm -rf $GOPATH/pkg/mod/google.golang.org/grpc@v1.37.0
rm -rf $WORKING_DIR/tmp

# Get Avalanchego source
go get -v -d github.com/ava-labs/avalanchego/...
cd $GOPATH/src/github.com/ava-labs/avalanchego
git config --global advice.detachedHead false
# Hard-coded commit to tag v1.5.2, at the time of this authoring
# https://github.com/ava-labs/avalanchego/releases/tag/v1.5.2
git checkout f2e51d790430a171e6d39f72911d98f134942a55

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

chmod -R 755 $GOPATH/pkg/mod/google.golang.org/grpc@v1.37.0

cp -R $WORKING_DIR/src/grpc@v1.37.0/. $GOPATH/pkg/mod/google.golang.org/grpc@v1.37.0
cd $GOPATH/pkg/mod/google.golang.org/grpc@v1.37.0

# build
make build

cd $GOPATH/src/github.com/ava-labs/avalanchego

export ROCKSDBALLOWED=true
./scripts/build.sh
rm -rf ./scripts/coreth_changes
rm -rf ./scripts/grpc_changes