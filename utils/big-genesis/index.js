async function main() {
//  await hre.run('compile');
//  signers = await hre.ethers.getSigners();

  const { argv } = require("yargs")
  .scriptName("big-genesis")
  .usage("Usage: $0 -c num")
  .example(
    "$0 -c 10",
    "Create a random address json file."
  )
  .option("c", {
    alias: "count",
    describe: "The number of addresses.",
    demandOption: "The number of addresses to create are required.",
    type: "number",
    nargs: 1,
  })
  .describe("help", "Show help."); // Override --help usage message.

  const { threads } = argv;
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
