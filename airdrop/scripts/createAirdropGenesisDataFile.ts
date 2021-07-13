import { BADNAME } from 'dns';
import * as fs from 'fs';
import {validateFile} from "./utils/processFile";
import { writeError } from './utils/utils';
const Web3Utils = require('web3-utils');
const parse = require('csv-parse/lib/sync');
const BN = Web3Utils.toBN;

const now = new Date()
const logFileName = `files/logs/${now.toISOString()}_airdrop_data_gen_log.txt`;
console.log(logFileName)

fs.writeFile(logFileName, `Log file created at ${now.toISOString()} \n`, writeError);

// parse CLI parameter
const parameters = process.argv.slice(2)
// Snapshot file 
if(!parameters.includes("--snapshot-file")){
    console.log("You must provide snapshot file with --snapshot-file flag");
    process.exit(0);
}
const snapshotFile = parameters[parameters.indexOf("--snapshot-file")+1]
console.log(`Script run with: --snapshot-file : ${snapshotFile}`)

// Constants
const contingentPercentage:number = 100;    // The percentage of promised airdrop distributed
const conversionFactor:number = 1.0073;     // The factor for converting XPR balance to Flare Wei 
const initialAirdropPercentage:number = 15  // 
const expectedDistributedWei:number = BN(45)    // We want this to be 45 bil Spark token

fs.appendFile(logFileName, `Constants: \n`, writeError);
fs.appendFile(logFileName, `Contingent Percentages    : ${contingentPercentage} % \n`, writeError);
fs.appendFile(logFileName, `Expected distributed Wei  : ${expectedDistributedWei} \n`, writeError);


// Parse the CSV file
let data = fs.readFileSync(snapshotFile, "utf8");
const parsed_file = parse(data, {
  columns: true,
  skip_empty_lines: true,
  delimiter: ','
})

validateFile(parsed_file);


