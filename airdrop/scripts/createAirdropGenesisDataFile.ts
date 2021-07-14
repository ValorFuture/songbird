import { BADNAME } from 'dns';
import * as fs from 'fs';
import {calculateConversionFactor, validateFile} from "./utils/processFile";
import { writeError } from './utils/utils';
const Web3Utils = require('web3-utils');
const parse = require('csv-parse/lib/sync');
import BigNumber from "bignumber.js";

const separatorLine = "------------------------------------------------------------\n"

const now = new Date()
const logFileName = `files/logs/${now.toISOString()}_airdrop_data_gen_log.txt`;
console.log(logFileName)

fs.writeFileSync(logFileName, `Log file created at ${now.toISOString()} \n`);

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
const initialAirdropPercentage:number = 15  // 
const expectedDistributedWei:BigNumber = new BigNumber(45*10**9*10**18);   
// We want this to be 45 bil Spark token so 45 * 10^9 * 10^18 Wei 

const constantRepString = `Constants
Contingent Percentages      : ${contingentPercentage} %
Expected distributed Wei    : ${expectedDistributedWei.toFixed()}
Initial Airdrop percentage  : ${initialAirdropPercentage} %`

fs.appendFileSync(logFileName, constantRepString + "\n");
console.log(constantRepString);

// Parse the CSV file
let data = fs.readFileSync(snapshotFile, "utf8");
const parsed_file = parse(data, {
  columns: true,
  skip_empty_lines: true,
  delimiter: ','
})

// Validate Input CSV File
console.log(separatorLine+"Input file problems")
fs.appendFileSync(logFileName, separatorLine+"Input file problems \n");
validateFile(parsed_file,logFileName);

// Calculating conversion factor
let conversionFactor = calculateConversionFactor(parsed_file,expectedDistributedWei);
console.log(separatorLine+`Conversion factor : ${conversionFactor.toString()}`)
fs.appendFileSync(logFileName, separatorLine+`Conversion factor : ${conversionFactor.toString()} \n`);

