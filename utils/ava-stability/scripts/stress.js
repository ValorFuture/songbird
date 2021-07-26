// We require the Hardhat Runtime Environment explicitly here. This is optional 
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const counters = [];
const promises = [];
const stats = [];
let signers = [];
let executionTimeMs = 0;
let maxTimeMs = 0;
let minTimeMs = Number.MAX_VALUE;
let startMs = Date.now();

const makeIncrementPromise = (i, offset) => {
  return new Promise(async (resolve) => {
    stats[i].attempts += 1;
    stats[i].lastStartTs = Date.now();
    await counters[i].connect(signers[i + offset]).increment()
      .then(async (tx) => {
        // wait until the transaction is mined
        await tx.wait()
          .then(async (receipt) => {
            await counters[i].count()
              .then((count) => {
                let elapsed = Date.now() - stats[i].lastStartTs;
                executionTimeMs += elapsed;
                if (elapsed < minTimeMs) {
                  minTimeMs = elapsed;
                }
                if (elapsed > maxTimeMs) {
                  maxTimeMs = elapsed;
                }
                stats[i].successes = count;
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
  .option("o", {
    alias: "offset",
    describe: "The number signers to offset in case multiple client applications are run.",
    default: 0,
    type: "number",
    nargs: 1,
  })
  .describe("help", "Show help."); // Override --help usage message.

  const { threads, offset } = argv;

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

  // Prime promises and stats for each contract
  for (let i = 0; i < threads; i++) {
    stats[i] = { attempts: 0, successes: 0, lastStartTs: 0 }
    promises[i] = makeIncrementPromise(i, offset);
  }
  
  // Spew out stats
  setInterval(() => {
    let txCount = 0;
    let txAttempt = 0;
    stats.map((stat) => {txCount += Number(stat.successes)});
    stats.map((stat) => {txAttempt += Number(stat.attempts)});
    console.log(`Attempts = ${txAttempt}; Tx count = ${txCount}; TPS = ${txCount / ((Date.now() - startMs) / 1000)}; Avg ms = ${executionTimeMs / txCount}; Max ms = ${maxTimeMs}; Min ms = ${minTimeMs}`)
  }, 5000);

  // Spin forever and pound the validator with transactions
  while(true) {
    const index = await Promise.race(promises);
    promises[index] = makeIncrementPromise(index, offset);
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
