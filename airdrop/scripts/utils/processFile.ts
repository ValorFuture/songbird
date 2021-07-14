const Web3Utils = require('web3-utils');
const RippleAPI = require('ripple-lib').RippleAPI;
import * as fs from 'fs';
import { writeError } from './utils';
import BigNumber from "bignumber.js";

const RippleApi = new RippleAPI({
    server: 'wss://s1.ripple.com' // Public rippled server hosted by Ripple, Inc.
  });

interface LineItem {
    XPRAddress: string,
    FlareAddress: string,
    SparkBalanceWei: string
}

export function validateFile(parsedFile: LineItem[], logFile: string) {
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let lineItem = parsedFile[lineIndex];
        if(!RippleApi.isValidAddress(lineItem.XPRAddress)){
            console.log(`Line ${lineIndex + 2}: XPR address is invalid`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: XPR address is invalid \n`);
        }
        if(!Web3Utils.isAddress(lineItem.FlareAddress)){
            console.log(`Line ${lineIndex + 2}: Flare address is invalid`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: Flare address is invalid \n`);
        }
        let numberBalance = +lineItem.SparkBalanceWei;
        if(isNaN(numberBalance)){
            console.log(`Line ${lineIndex + 2}: Balance is not a valid number`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: Balance is not a valid number \n`);
        }
    }
}

export function calculateConversionFactor(parsedFile: LineItem[], expected_total:any): any{
    let total = new BigNumber(0);
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let lineItem = parsedFile[lineIndex];
        if(!isNaN(+lineItem.SparkBalanceWei)){
            total = total.plus(lineItem.SparkBalanceWei);
        }
    }
    let expectedTot = new BigNumber(expected_total);
    return expectedTot.div(total);
}
