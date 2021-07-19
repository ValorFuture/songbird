const Web3Utils = require('web3-utils');
const RippleAPI = require('ripple-lib').RippleAPI;
import * as fs from 'fs';
import { writeError } from './utils';
import BigNumber from "bignumber.js";
import { removeUndefined } from 'ripple-lib/dist/npm/common';

BigNumber.config({ ROUNDING_MODE: BigNumber.ROUND_FLOOR, DECIMAL_PLACES: 20 })

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
    validAccountsLen: number,
    invalidAccountsLen: number
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
    let invalidAccountsLen:number = 0
    let seenXPRAddresses = new Set();
    let seenXPRAddressesDetail: {[name: string]: number } = {};;
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let lineItem = parsedFile[lineIndex];
        let isValid = true;
        if(!RippleApi.isValidAddress(lineItem.XPRAddress)){
            console.log(`Line ${lineIndex + 2}: XPR address is invalid`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: XPR address is invalid \n`);
            isValid = false;
        }
        if(seenXPRAddresses.has(lineItem.XPRAddress)){
            // We have already seen this XPR address
            console.log(`Line ${lineIndex + 2}: XPR address is duplicate of line ${seenXPRAddressesDetail[lineItem.XPRAddress]}`);
            fs.appendFileSync(logFile, `Line ${lineIndex + 2}: XPR address is duplicate of line ${seenXPRAddressesDetail[lineItem.XPRAddress]}\n`);
            isValid = false;
        }
        if(!seenXPRAddresses.has(lineItem.XPRAddress)){
            seenXPRAddresses.add(lineItem.XPRAddress);
            seenXPRAddressesDetail[lineItem.XPRAddress] = lineIndex;
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
        } else {
            invalidAccountsLen += 1;
        }
    } 
    return {validAccounts, validAccountsLen, invalidAccountsLen};
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
(parsedFile: LineItem[], validAccounts: validateRes, contingentPercentage: BigNumber,
conversionFactor: BigNumber, initialAirdropPercentage: BigNumber ):airdropGenesisRes{
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
            processedAccounts[lineIndex] = `"${lineItem.FlareAddress.substring(2)}": {"balance": "0x${accBalance.toString(16)}" },`;
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