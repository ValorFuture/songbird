require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-web3");

let fs = require('fs');

accounts = [
  ...JSON.parse(fs.readFileSync('test-1020-accounts.json'))
];

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async () => {
  const accounts = await ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */

module.exports = {
  defaultNetwork: "scdev",

  networks: {
    scdev: {
      url: "http://127.0.0.1:9660/ext/bc/C/rpc",
      gas: 8000000,
      timeout: 40000,
      accounts: accounts.map((x) => x.privateKey)
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
        version: "0.8.4",
        settings: {
          optimizer: {
            enabled: false,
            runs: 200
          }
        }
      }
    ]
  }
}
