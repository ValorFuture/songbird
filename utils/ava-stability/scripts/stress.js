// We require the Hardhat Runtime Environment explicitly here. This is optional 
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const counters = [];
const promises = [];
let signers = [];

const makeIncrementPromise = (i) => {
  return new Promise(async (resolve) => {
    await counters[i].connect(signers[i]).increment()
      .then(async (tx) => {
        // wait until the transaction is mined
        await tx.wait()
          .then(async (receipt) => {
            await counters[i].count()
              .then((count) => {
                console.log(`counters[${i}] = ${counters[i].address}; count = ${count.toString()}`);
                resolve(i);
              })
              .catch((err) => {
                console.log(`Error getting count for counters[${i}] = ${err}`);
                resolve(i);
              });
          })
          .catch((err) => {
            console.log(`Error waiting for increment to mine for counters[${i}] = ${err}`);
            resolve(i);
          });
      })
      .catch((err) => {
        console.log(`Error waiting for increment tx for counters[${i}] = ${err}`);
        resolve(i);
      });
  });      
};

async function main() {
  await hre.run('compile');
  signers = await hre.ethers.getSigners();

  const { argv } = require("yargs")
  .scriptName("stress")
  .usage("Usage: $0 -t num")
  .example(
    "$0 -t 10",
    "Runs a stress test to simultaneous transactions to Ava validator."
  )
  .option("t", {
    alias: "threads",
    describe: "The number of threads to start.",
    demandOption: "The number of threads are required.",
    type: "number",
    nargs: 1,
  })
  .describe("help", "Show help."); // Override --help usage message.

  const { threads } = argv;

  // Deploy the Counter contract
  const Counter = await hre.ethers.getContractFactory("Counter");

  // Deploy threads number of contracts
  deploys = [];
  for (let i = 0; i < threads; i++) {
    counters[i] = await Counter.deploy();
    deploys.push(counters[i].deployed());
  }
  await Promise.all(deploys);

//  nonce = await hre.web3.eth.getTransactionCount(hre.web3.eth.accounts.privateKeyToAccount(hre.network.config.accounts[0]).address);

  // Prime promises for each contract
  for (let i = 0; i < threads; i++) {
    promises[i] = makeIncrementPromise(i);
  }

  // Spin forever and pound the validator with transactions
  while(true) {
    const index = await Promise.race(promises);
    promises[index] = makeIncrementPromise(index);
  }
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
