#!/bin/bash
if [[ $(pwd) =~ *" "* ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
if [ -z ${GOPATH+x} ]; then echo "GOPATH is not set, visit https://github.com/golang/go/wiki/SettingGOPATH" && exit; fi
if [[ $(go version) != *"go1.15"* ]]; then echo "Go version is not go1.15" && exit; fi
if [ "$(dpkg --print-architecture)" != "amd64" ]; then echo "Machine architecture is not amd64" && exit; fi

WORKING_DIR=$(pwd)

#GENESIS_FILE=genesis_coston.go
#GENESIS_FILE=genesis_scdev_160k.go

if [ -z "$GENESIS_FILE" ]; then
   GENESIS_FILE=genesis_coston.go
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
# Hard-coded commit to tag v1.4.11-rc.0, at the time of this authoring
git checkout ac32de45ffd6769007f250f123a5d5dae8230456

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