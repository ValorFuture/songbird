import { BADNAME } from 'dns';
import * as fs from 'fs';
import {calculateConversionFactor, createFlareAirdropGenesisData, validateFile} from "./utils/processFile";
import { writeError } from './utils/utils';
const Web3Utils = require('web3-utils');
const parse = require('csv-parse/lib/sync');
import BigNumber from "bignumber.js";
import { createGenesisFileData } from './utils/genesisFile';

const separatorLine = "------------------------------------------------------------\n"

const now = new Date()
const logFileName = `files/logs/${now.toISOString()}_airdrop_data_gen_log.txt`;
console.log(logFileName)

fs.writeFileSync(logFileName, `Log file created at ${now.toISOString()} GMT(+0)\n`);

// parse CLI parameter
const parameters = process.argv.slice(2)
// Snapshot file 
if(!parameters.includes("--snapshot-file")){
    console.log("You must provide snapshot file with --snapshot-file flag");
    process.exit(0);
}
const snapshotFile = parameters[parameters.indexOf("--snapshot-file")+1]
// go Genesis output file 
if(!parameters.includes("--genesis-file")){
    console.log("You must provide go output file for genesis with --genesis-file flag");
    process.exit(0);
}
const goGenesisFile = parameters[parameters.indexOf("--genesis-file")+1]
// go genesis override flag
const canOverwriteGenesis = parameters.includes("--override")
if (fs.existsSync(goGenesisFile)) {
    if(!canOverwriteGenesis){
        console.log("go Genesis file already exist, if you want to overwrite it provide --override");
        process.exit(0);
    }
    else {
        fs.writeFile(goGenesisFile, '', function (err) {
            if (err) {
                console.log("Can't create file at provided destination")
                throw err
            };
          });
    }
    // File exists in path
  } else {
    fs.writeFile(goGenesisFile, '', function (err) {
        if (err) {
            console.log("Can't create file at provided destination")
            throw err
        };
      });
  }

const inputRepString = `Script run with 
--snapshot-file              : ${snapshotFile}
--genesis-file               : ${goGenesisFile}`
fs.appendFileSync(logFileName, inputRepString + "\n");
console.log(inputRepString);

// Constants
const contingentPercentage:number = 1;    // The percentage of promised airdrop distributed 
const initialAirdropPercentage:number = 0.15  // 
const expectedDistributedWei:BigNumber = new BigNumber(45*10**9*10**18);   
// We want this to be 45 bil Spark token so 45 * 10^9 * 10^18 Wei 

const constantRepString = separatorLine + `Constants
Contingent Percentages       : ${contingentPercentage * 100} %
Initial Airdrop percentage   : ${initialAirdropPercentage * 100} %
Total distributed Wei        : ${expectedDistributedWei.toFixed()}
Wei distributed at Airdrop   : ${expectedDistributedWei.multipliedBy(initialAirdropPercentage).toFixed()}`

fs.appendFileSync(logFileName, constantRepString + "\n");
console.log(constantRepString);

// Parse the CSV file
let data = fs.readFileSync(snapshotFile, "utf8");
const parsed_file = parse(data, {
  columns: true,
  skip_empty_lines: true,
  delimiter: ','
})

console.log(separatorLine+"Input file problems")
fs.appendFileSync(logFileName, separatorLine+"Input file problems \n");
// Validate Input CSV File
let validatedData = validateFile(parsed_file,logFileName);
// Log Validation results
console.log(`Number of valid accounts     : ${validatedData.validAccountsLen}`)
fs.appendFileSync(logFileName, `Number of valid accounts     : ${validatedData.validAccountsLen}\n`);

// Calculating conversion factor
let conversionFactor = calculateConversionFactor(parsed_file, validatedData, expectedDistributedWei);
// Log conversion factor results
console.log(separatorLine+`Conversion factor            : ${conversionFactor.conversionFactor.toString()}`)
fs.appendFileSync(logFileName, separatorLine+`Conversion factor            : ${conversionFactor.conversionFactor.toString()} \n`);

// Create Flare balance json
let convertedAirdropData = createFlareAirdropGenesisData(parsed_file, validatedData,
     contingentPercentage, conversionFactor.conversionFactor, initialAirdropPercentage);
// Log balance created
console.log(`Number of processed accounts : ${convertedAirdropData.processedAccountsLen}`)
fs.appendFileSync(logFileName, `Number of processed accounts : ${convertedAirdropData.processedAccountsLen}\n`);
console.log(`Total FLR added to accounts  : ${convertedAirdropData.processedWei.toFixed()}`)
fs.appendFileSync(logFileName, `Total FLR added to accounts  : ${convertedAirdropData.processedWei.toFixed()}\n`);
console.log(convertedAirdropData);


// Create go genesis file
const fileData = createGenesisFileData(convertedAirdropData.processedAccounts.join("\n              "))
fs.appendFileSync(goGenesisFile, fileData);