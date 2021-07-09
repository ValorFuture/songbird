import * as fs from 'fs';
import {validateFile} from "./utils/processFile";
const Web3Utils = require('web3-utils');
const parse = require('csv-parse/lib/sync');

// parse CLI parameter
const parameters = process.argv.slice(2)
// Snapshot file 
if(!parameters.includes("--snapshot-file")){
    console.log("You must provide snapshot file with --snapshot-file flag");
    process.exit(0);
}
const snapshotFile = parameters[parameters.indexOf("--snapshot-file")+1]
console.log(`Script run with: --snapshot-file : ${snapshotFile}`)


// Parse the CSV file
let data = fs.readFileSync(snapshotFile, "utf8");
const parsed_file = parse(data, {
  columns: true,
  skip_empty_lines: true,
  delimiter: ','
})

validateFile(parsed_file);


