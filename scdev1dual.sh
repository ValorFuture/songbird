#!/bin/bash
# Catalina users: tired of seeing network nag dialogs:
# I deleted avalanchego from firewall, then did: https://stackoverflow.com/questions/9845502/how-do-i-get-the-mac-os-x-firewall-to-permanently-allow-my-ios-app
# If you pass --existing, then the script will bypass the build.
# If you pass --clean, any existing chain database will be blasted.
if [[ $(pwd) =~ " " ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
if [ -z ${GOPATH+x} ]; then echo "GOPATH is not set, visit https://github.com/golang/go/wiki/SettingGOPATH" && exit; fi
if [ -z ${XRP_APIs+x} ] || [ "$XRP_APIs" == "url1, url2, ..., urlN" ]; then echo "XRP_APIs is not set, please set it using the form: $ export XRP_APIs=\"url1, url2, ..., urlN\"" && exit; fi
if [[ $(go version) != *"go1.15.5"* ]]; then echo "Go version is not go1.15.5" && exit; fi
XRP_APIs_JOINED="$(echo -e "${XRP_APIs}" | tr -d '[:space:]')"
printf "\x1b[34mFlare Network 1-Node Smart Contract Dev Team Static IP Address Dual Node Local Deployment\x1b[0m\n\n"
AVALANCHEGO_VERSION=@v1.3.2
CORETH_VERSION=@v0.4.2-rc.4

EXEC_DIR=$(pwd)
LOG_DIR=$(pwd)/logs
CONFIG_DIR=$(pwd)/config
PKG_DIR=$GOPATH/pkg/mod/github.com/ava-labs
NODE_DIR=$PKG_DIR/avalanchego$AVALANCHEGO_VERSION
CORETH_DIR=$PKG_DIR/coreth$CORETH_VERSION

if echo $1 | grep -e "--existing" -q
then
	cd $NODE_DIR
  if echo $2 | grep -e "--clean" -q
  then
    echo "Cleaning chain db..."
    rm -rf $NODE_DIR/db/
  fi
else
	rm -rf logs
	mkdir logs
	rm -rf $NODE_DIR
	mkdir -p $PKG_DIR
	cp -r fba-avalanche/avalanchego $NODE_DIR
	rm -rf $CORETH_DIR
	cp -r fba-avalanche/coreth $CORETH_DIR
	cd $NODE_DIR
	rm -rf $NODE_DIR/db/
	mkdir -p $LOG_DIR/node00
	mkdir -p $LOG_DIR/node01
	printf "Building Flare Core...\n"
	./scripts/build.sh
fi

# NODE 0
printf "Launching Node 0 at 192.168.2.1:9660\n"
./build/avalanchego \
    --public-ip=192.168.2.1 \
    --snow-sample-size=1 \
    --snow-quorum-size=1 \
    --http-port=9660 \
    --staking-port=9661 \
    --db-dir=$(pwd)/db/node00/ \
    --staking-enabled=true \
    --network-id=scdev \
    --bootstrap-ips= \
    --bootstrap-ids= \
    --staking-tls-cert-file=$(pwd)/config/keys/node00/node.crt \
    --staking-tls-key-file=$(pwd)/config/keys/node00/node.key \
    --log-level=debug \
    --log-dir=$LOG_DIR/node00 \
    --validators-file=$(pwd)/config/validators/scdev/1619370000.json \
    --alert-apis="https://flare.network" \
    --xrp-apis=$XRP_APIs_JOINED \
    &
sleep 2
NODE_00_PID=`lsof -n -i4TCP:9660 | grep LISTEN | cut -d ' ' -f2`

# NODE 1
printf "Launching Node 1 at 192.168.2.2:9662\n"
./build/avalanchego \
    --public-ip=192.168.2.2 \
    --snow-sample-size=1 \
    --snow-quorum-size=1 \
    --http-host=192.168.2.2 \
    --http-port=9662 \
    --staking-port=9663 \
    --db-dir=$(pwd)/db/node01/ \
    --staking-enabled=true \
    --network-id=scdev \
    --bootstrap-ips=192.168.2.1:9661 \
    --bootstrap-ids=NodeID-D4jrXzkioNZkqbPNuPmk3hR9Ee8oXLvDJ \
    --staking-tls-cert-file=$(pwd)/config/keys/node01/node.crt \
    --staking-tls-key-file=$(pwd)/config/keys/node01/node.key \
    --log-level=debug \
    --log-dir=$LOG_DIR/node01 \
    --validators-file=$(pwd)/config/validators/scdev/1619370000.json \
    --alert-apis="https://flare.network" \
    --xrp-apis=$XRP_APIs_JOINED \
    &
sleep 2
NODE_01_PID=`lsof -n -i4TCP:9662 | grep LISTEN | cut -d ' ' -f2`

printf "\n"
read -p "Press enter to stop background node processes"
kill $NODE_00_PID
kill $NODE_01_PID