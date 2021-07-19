import { BADNAME } from 'dns';
import * as fs from 'fs';
import {calculateConversionFactor, createFlareAirdropGenesisData, validateFile} from "./utils/processFile";
import { writeError } from './utils/utils';
const Web3Utils = require('web3-utils');
const parse = require('csv-parse/lib/sync');
import BigNumber from "bignumber.js";
import { createGenesisFileData } from './utils/genesisFile';

const separatorLine = "--------------------------------------------------------------------------------\n"

// parse CLI parameter
const parameters = process.argv.slice(2)
// Snapshot file (--snapshot-file)
if(!parameters.includes("--snapshot-file")){
    console.log("You must provide snapshot file with --snapshot-file flag");
    process.exit(0);
}
const snapshotFile = parameters[parameters.indexOf("--snapshot-file")+1]
// go Genesis output file (--genesis-file)
if(!parameters.includes("--genesis-file")){
    console.log("You must provide go output file for genesis with --genesis-file flag");
    process.exit(0);
}
const goGenesisFile = parameters[parameters.indexOf("--genesis-file")+1]
// go genesis override flag (--override)
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
// log path (--log-path)
let logPath = ""
if(parameters.includes("--log-path")){
    logPath = parameters[parameters.indexOf("--log-path")+1];
} 
if(!parameters.includes("--log-path")) {
    logPath = "files/logs/";
}
const now = new Date()
const logFileName = logPath+`${now.toISOString()}_airdrop_data_gen_log.txt`;
console.log(logFileName)
fs.writeFileSync(logFileName, `Log file created at ${now.toISOString()} GMT(+0)\n`);



const inputRepString = `Script run with 
--snapshot-file                             : ${snapshotFile}
--genesis-file                              : ${goGenesisFile}
--override                                  : ${canOverwriteGenesis}`
fs.appendFileSync(logFileName, inputRepString + "\n");
console.log(inputRepString);

// Constants
const contingentPercentage:number = 1;    // The percentage of promised airdrop distributed 
const initialAirdropPercentage:number = 0.15  // 
const expectedDistributedWei:BigNumber = new BigNumber(45*10**9*10**18);   
// We want this to be 45 bil Spark token so 45 * 10^9 * 10^18 Wei 

const constantRepString = separatorLine + `Constants
Contingent Percentages                      : ${contingentPercentage * 100} %
Initial Airdrop percentage                  : ${initialAirdropPercentage * 100} %
Total distributed Wei                       : ${expectedDistributedWei.toFixed()}
Wei distributed at Airdrop                  : ${expectedDistributedWei.multipliedBy(initialAirdropPercentage).toFixed()}`

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
console.log(`Number of valid accounts                    : ${validatedData.validAccountsLen}`)
fs.appendFileSync(logFileName, `Number of valid accounts                    : ${validatedData.validAccountsLen}\n`);
console.log(`Number of invalid accounts                  : ${validatedData.invalidAccountsLen}`)
fs.appendFileSync(logFileName, `Number of invalid accounts                  : ${validatedData.invalidAccountsLen}\n`);

// Calculating conversion factor
console.log(separatorLine+"Input file processing")
fs.appendFileSync(logFileName, separatorLine+"Input file processing\n");
let conversionFactor = calculateConversionFactor(parsed_file, validatedData, expectedDistributedWei);
// Log conversion factor results
console.log(`Conversion factor                           : ${conversionFactor.conversionFactor.toString()}`)
fs.appendFileSync(logFileName, `Conversion factor                           : ${conversionFactor.conversionFactor.toString()} \n`);
console.log(`Total XPR balance read                      : ${conversionFactor.totalXPRBalance.toFixed()}`)
fs.appendFileSync(logFileName, `Total XPR balance read                      : ${conversionFactor.totalXPRBalance.toFixed()} \n`);

// Create Flare balance json
let convertedAirdropData = createFlareAirdropGenesisData(parsed_file, validatedData,
     contingentPercentage, conversionFactor.conversionFactor, initialAirdropPercentage);
// Log balance created
console.log(`Number of processed accounts                : ${convertedAirdropData.processedAccountsLen}`)
fs.appendFileSync(logFileName, `Number of processed accounts                : ${convertedAirdropData.processedAccountsLen}\n`);
console.log(`Total FLR added to accounts                 : ${convertedAirdropData.processedWei.toFixed()}`)
fs.appendFileSync(logFileName, `Total FLR added to accounts                 : ${convertedAirdropData.processedWei.toFixed()}\n`);

// **********************
// Do final health checks
let accounts_match = convertedAirdropData.processedAccountsLen == validatedData.validAccountsLen;
console.log(separatorLine+"Health checks")
fs.appendFileSync(logFileName, separatorLine+"Health checks\n");
console.log(`Read and processed account number match     : ${accounts_match}`)
fs.appendFileSync(logFileName, `Read and processed account number match     : ${accounts_match} \n`);


// Create go genesis file
const fileData = createGenesisFileData(convertedAirdropData.processedAccounts.join("\n              "))
fs.appendFileSync(goGenesisFile, fileData);
