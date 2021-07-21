import { BADNAME } from 'dns';
import * as fs from 'fs';
import {createFlareAirdropGenesisData, validateFile} from "./utils/processFile";
import { writeError } from './utils/utils';
const Web3Utils = require('web3-utils');
const parse = require('csv-parse/lib/sync');
import BigNumber from "bignumber.js";
import { createGenesisFileData } from './utils/genesisFile';

BigNumber.config({ ROUNDING_MODE: BigNumber.ROUND_FLOOR, DECIMAL_PLACES: 20 })

// CONSTANTS
const initialAirdropPercentage:BigNumber = new BigNumber(0.15);
const conversionFactor:BigNumber = new BigNumber(1.0073);

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
let logPath = "files/logs/";
if(parameters.includes("--log-path")){
    logPath = parameters[parameters.indexOf("--log-path")+1];
} 
const now = new Date()
const logFileName = logPath+`${now.toISOString()}_airdrop_data_gen_log.txt`;
console.log(logFileName)
fs.writeFileSync(logFileName, `Log file created at ${now.toISOString()} GMT(+0)\n`);

// log path (--contingent-percentage)
let contgPer = "100"
if(parameters.includes("--contingent-percentage")){
    contgPer = parameters[parameters.indexOf("--contingent-percentage")+1];
} 
const contingentPercentage:BigNumber = new BigNumber(contgPer).dividedBy(100);

const inputRepString = `Script run with 
--snapshot-file                             : ${snapshotFile}
--genesis-file                              : ${goGenesisFile}
--override                                  : ${canOverwriteGenesis}
--log-path                                  : ${logPath}`
fs.appendFileSync(logFileName, inputRepString + "\n");
console.log(inputRepString);

const constantRepString = separatorLine + `Constants
Contingent Percentages                      : ${contingentPercentage.multipliedBy(100).toFixed()} %
Initial Airdrop percentage                  : ${initialAirdropPercentage.multipliedBy(100).toFixed()} %
Conversion Factor                           : ${conversionFactor.toFixed()}`
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
console.log(`Total XPR valid balance read                : ${validatedData.totalXPRBalance.toFixed()}`)
fs.appendFileSync(logFileName, `Total valid XPR balance read                : ${validatedData.totalXPRBalance.toFixed()} \n`);
let expectedFlrToDistribute:BigNumber = new BigNumber(0);
expectedFlrToDistribute = validatedData.totalXPRBalance;
expectedFlrToDistribute = expectedFlrToDistribute.multipliedBy(conversionFactor).multipliedBy(contingentPercentage).multipliedBy(initialAirdropPercentage);
console.log(`Expected Flare to distribute                : ${expectedFlrToDistribute.toFixed()}`)
fs.appendFileSync(logFileName, `Expected Flare to distribute                : ${expectedFlrToDistribute.toFixed()} \n`);

// Calculating conversion factor
console.log(separatorLine+"Input file processing")
fs.appendFileSync(logFileName, separatorLine+"Input file processing\n");
// Create Flare balance json
let convertedAirdropData = createFlareAirdropGenesisData(parsed_file, validatedData,
     contingentPercentage, conversionFactor, initialAirdropPercentage);
// Log balance created
const zeroPad = (num:any, places:any) => String(num).padStart(places, '0')
console.log(`Number of processed accounts                : ${convertedAirdropData.processedAccountsLen}`)
fs.appendFileSync(logFileName, `Number of processed accounts                : ${convertedAirdropData.processedAccountsLen}\n`);
console.log(`Number of Flare accounts added to genesis   : ${convertedAirdropData.processedAccounts.length}`)
fs.appendFileSync(logFileName, `Number of Flare accounts added to genesis   : ${convertedAirdropData.processedAccounts.length}\n`);
for(let i=0; i<convertedAirdropData.accountsDistribution.length; i++){
    if(convertedAirdropData.accountsDistribution[i]>0){
        console.log(`Number of Flare addresses added ${zeroPad(i,3)} times   : ${convertedAirdropData.accountsDistribution[i]}`)
        fs.appendFileSync(logFileName, `Number of Flare addresses added ${zeroPad(i,3)} times   : ${convertedAirdropData.accountsDistribution[i]}\n`);
    }
}
console.log(`Total FLR added to accounts                 : ${convertedAirdropData.processedWei.toFixed()}`)
fs.appendFileSync(logFileName, `Total FLR added to accounts                 : ${convertedAirdropData.processedWei.toFixed()}\n`);


// **********************
// Do final health checks
let healthy = true;
let accounts_match = convertedAirdropData.processedAccountsLen == validatedData.validAccountsLen;
healthy = healthy && accounts_match;
console.log(separatorLine+"Health checks")
fs.appendFileSync(logFileName, separatorLine+"Health checks\n");
console.log(`Read and processed account number match     : ${accounts_match}`)
fs.appendFileSync(logFileName, `Read and processed account number match     : ${accounts_match} \n`);

if(healthy){
    const fileData = createGenesisFileData(convertedAirdropData.processedAccounts.join("\n              "))
    fs.appendFileSync(goGenesisFile, fileData);
    console.log(`Created output genesis file at              : ${goGenesisFile}`)
    fs.appendFileSync(logFileName, `Created output genesis file at              : ${goGenesisFile} \n`); 
}

