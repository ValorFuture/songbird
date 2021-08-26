
import "@nomiclabs/hardhat-truffle5";
import "@nomiclabs/hardhat-web3";
import { HardhatUserConfig } from "hardhat/config";
import '@typechain/hardhat'
import '@nomiclabs/hardhat-ethers'
import '@nomiclabs/hardhat-waffle'

let fs = require('fs');

let accounts  = [
  ...JSON.parse(fs.readFileSync('./tests/test-1020-accounts.json'))
];

const config: HardhatUserConfig = {
  defaultNetwork: "hardhat",

  networks: {
    scdev: {
      url: "http://127.0.0.1:9660/ext/bc/C/rpc",
      gas: 80000000,
      timeout: 40000,
      accounts: accounts.map((x: any) => x.privateKey)
    },
    hardhat: {
      accounts,
      initialDate: "2021-01-01",  // no time - get UTC @ 00:00:00
      blockGasLimit: 125000000 // 10x ETH gas
    },
    local: {
      url: 'http://127.0.0.1:8545',
      chainId: 31337
    }
  },
  solidity: {
    compilers: [
      {
        version: "0.7.6",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          }
        }
      }
    ],  
  },

  paths: {
    sources: "./stateco/",
    tests: "./tests",
    cache: "./cache",
    artifacts: "./artifacts",
    deploy: 'deploy',
    deployments: 'deployments',
    imports: 'imports'
  },

  mocha: {
    timeout: 1000000000
  },

  typechain: {
    outDir: "./types",
    target: "ethers-v5",
  },
};

export default config;