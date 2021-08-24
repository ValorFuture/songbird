#!/bin/bash
if [[ $(pwd) =~ *" "* ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
if [ -z ${GOPATH+x} ]; then echo "GOPATH is not set, visit https://github.com/golang/go/wiki/SettingGOPATH" && exit; fi
if [[ $(go version) != *"go1.15"* ]]; then echo "Go version is not go1.15" && exit; fi
if [ "$(uname -m)" != "x86_64" ]; then echo "Machine architecture is not x86_64" && exit; fi

WORKING_DIR=$(pwd)
GRPC_SRC_PATH=$GOPATH/src/google.golang.org/grpc

if [ -d $GOPATH/src/github.com/ava-labs ]; then
  echo "Removing old version..."
  chmod -R 775 $GOPATH/src/github.com/ava-labs && rm -rf $GOPATH/src/github.com/ava-labs
  chmod -R 775 $GOPATH/pkg/mod/github.com/ava-labs && rm -rf $GOPATH/pkg/mod/github.com/ava-labs
fi

echo "Downloading AvalancheGo..."
go get -v -d github.com/ava-labs/avalanchego/... &> /dev/null
cd $GOPATH/src/github.com/ava-labs/avalanchego
git config --global advice.detachedHead false
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

# Apply changes to coreth
echo "Copying Flare-specific changes for coreth building..."
mkdir ./scripts/coreth_changes
cp -R $WORKING_DIR/src/coreth/. ./scripts/coreth_changes/coreth

echo "Applying Flare-specific changes to grpc..."

# Grab grpc from repo
mkdir $WORKING_DIR/tmp
cd $WORKING_DIR/tmp
git clone https://github.com/grpc/grpc-go
cd grpc-go
# Switch to version required by Avalanche
git checkout v1.37.0

# Apply changes to grpc
cp $WORKING_DIR/src/grpc@v1.37.0/server.go .

# Go to Avalanche source directory to begin build
cd $GOPATH/src/github.com/ava-labs/avalanchego

# Copy grpc changes to location where ava coreth build script knows where to fetch
mkdir ./scripts/grpc_changes
cp -R $WORKING_DIR/tmp/grpc-go/. ./scripts/grpc_changes/grpc

# Modify grpc dependency to point to the vendored version for AvalancheGo
go mod edit -replace=google.golang.org/grpc@v1.37.0=$WORKING_DIR/tmp/grpc-go
go mod tidy

cd $GOPATH/src/github.com/ava-labs/avalanchego

export ROCKSDBALLOWED=true
./scripts/build.sh
rm -rf ./scripts/coreth_changes
chmod -R 775 $GOPATH/src/github.com/ava-labs
chmod -R 775 $GOPATH/pkg/mod/github.com/ava-labs
