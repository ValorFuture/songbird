const Web3Utils = require('web3-utils');
const RippleAPI = require('ripple-lib').RippleAPI;
import * as fs from 'fs';
import { writeError } from './utils';
import BigNumber from "bignumber.js";
import { removeUndefined } from 'ripple-lib/dist/npm/common';

const RippleApi = new RippleAPI({
    server: 'wss://s1.ripple.com' // Public rippled server hosted by Ripple, Inc.
  });

interface LineItem {
    XPRAddress: string,
    FlareAddress: string,
    XPRBalance: string
}
interface validateRes {
    validAccounts: boolean[],
    validAccountsLen: number
}

interface conversionFactorRes {
    conversionFactor: BigNumber,
    totalXPRBalance: BigNumber
}

interface airdropGenesisRes {
    processedAccounts: string[],
    processedAccountsLen: number,
    processedWei: BigNumber
}

export function validateFile(parsedFile: LineItem[], logFile: string):validateRes {
    let validAccountsLen:number = 0;
    let validAccounts:boolean[] = []
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let lineItem = parsedFile[lineIndex];
        let isValid = true;
        if(!RippleApi.isValidAddress(lineItem.XPRAddress)){
            console.log(`Line ${lineIndex + 2}: XPR address is invalid`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: XPR address is invalid \n`);
            isValid = false;
        }
        if(!Web3Utils.isAddress(lineItem.FlareAddress)){
            console.log(`Line ${lineIndex + 2}: Flare address is invalid`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: Flare address is invalid \n`);
            isValid = false;
        }
        let numberBalance = +lineItem.XPRBalance;
        if(isNaN(numberBalance)){
            console.log(`Line ${lineIndex + 2}: Balance is not a valid number`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: Balance is not a valid number \n`);
            isValid = false;
        }
        validAccounts[lineIndex] = isValid
        if(isValid){
            validAccountsLen += 1;
        }
    } 
    return {validAccounts, validAccountsLen};
}

export function calculateConversionFactor
(parsedFile: LineItem[], validAccounts: validateRes, expected_total:BigNumber): conversionFactorRes{
    let totalXPRBalance = new BigNumber(0);
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        if(validAccounts.validAccounts[lineIndex]){
            let lineItem = parsedFile[lineIndex];
            totalXPRBalance = totalXPRBalance.plus(lineItem.XPRBalance);
        }
    }
    let expectedTot = new BigNumber(expected_total);
    const conversionFactor = expectedTot.div(totalXPRBalance);
    return {
        conversionFactor,
        totalXPRBalance
    };
}

export function createFlareAirdropGenesisData
(parsedFile: LineItem[], validAccounts: validateRes, contingentPercentage: number,
conversionFactor: BigNumber, initialAirdropPercentage: number ):airdropGenesisRes{
    let processedAccountsLen:number = 0;
    let processedAccounts:string[] = [];
    let processedWei = new BigNumber(0)
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        if(validAccounts.validAccounts[lineIndex]){
            let lineItem = parsedFile[lineIndex];
            processedAccountsLen += 1
            let accBalance = new BigNumber(lineItem.XPRBalance);
            accBalance = accBalance.multipliedBy(contingentPercentage)
            accBalance = accBalance.multipliedBy(conversionFactor)
            accBalance = accBalance.multipliedBy(initialAirdropPercentage);
            // rounding down to 0 decimal places
            accBalance = accBalance.dp(0, BigNumber.ROUND_FLOOR);
            processedWei = processedWei.plus(accBalance);
            processedAccounts[lineIndex] = `"${lineItem.FlareAddress}": {"balance": "0x${accBalance.toString(16)}" },`;
        }
        else{
            processedAccounts[lineIndex] = ``;
        }
    }
    return {
        processedAccounts,
        processedAccountsLen,
        processedWei
    }
}