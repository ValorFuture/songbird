const Web3Utils = require('web3-utils');
const RippleAPI = require('ripple-lib').RippleAPI;
import * as fs from 'fs';
import { writeError } from './utils';
import BigNumber from "bignumber.js";
import { removeUndefined } from 'ripple-lib/dist/npm/common';

const TEN = new BigNumber(10);
BigNumber.config({ ROUNDING_MODE: BigNumber.ROUND_FLOOR, DECIMAL_PLACES: 20 })

const RippleApi = new RippleAPI({
    server: 'wss://s1.ripple.com' // Public rippled server hosted by Ripple, Inc.
  });

export interface LineItem {
    XRPAddress: string,
    FlareAddress: string,
    XRPBalance: string
}
interface validateRes {
    validAccounts: boolean[],
    validAccountsLen: number,
    invalidAccountsLen: number,
    totalXPRBalance: BigNumber,
    invalidXPRBalance: BigNumber,
}
interface airdropGenesisRes {
    processedAccounts: string[],
    processedAccountsLen: number,
    processedWei: BigNumber,
    accountsDistribution: number[]
}

export function validateFile(parsedFile: LineItem[], logFile: string):validateRes {
    let validAccountsLen:number = 0;
    let validAccounts:boolean[] = [];
    let invalidAccountsLen:number = 0;
    let seenXPRAddresses = new Set();
    let totalXPRBalance = new BigNumber(0);
    let invalidXPRBalance = new BigNumber(0);
    let seenXPRAddressesDetail: {[name: string]: number } = {};
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let lineItem = parsedFile[lineIndex];
        let isValid = true;
        let isValidNum = true;
        let readableIndex = lineIndex + 2;
        if(!RippleApi.isValidAddress(lineItem.XRPAddress)){
            console.log(`Line ${readableIndex}: XPR address is invalid`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: XPR address is invalid \n`);
            isValid = false;
        }
        if(seenXPRAddresses.has(lineItem.XRPAddress)){
            // We have already seen this XPR address
            console.log(`Line ${readableIndex}: XPR address is duplicate of line ${seenXPRAddressesDetail[lineItem.XRPAddress]}`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: XPR address is duplicate of line ${seenXPRAddressesDetail[lineItem.XRPAddress]}\n`);
            isValid = false;
        }
        if(!seenXPRAddresses.has(lineItem.XRPAddress)){
            seenXPRAddresses.add(lineItem.XRPAddress);
            seenXPRAddressesDetail[lineItem.XRPAddress] = lineIndex;
        } 
        if(!Web3Utils.isAddress(lineItem.FlareAddress)){
            console.log(`Line ${readableIndex}: Flare address is invalid`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: Flare address is invalid \n`);
            isValid = false;
        }
        let numberBalance = parseInt(lineItem.XRPBalance,10);
        if(isNaN(numberBalance)){
            console.log(`Line ${readableIndex}: Balance is not a valid number`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: Balance is not a valid number \n`);
            isValid = false;
            isValidNum = false;
        }
        validAccounts[lineIndex] = isValid;
        if(isValid){
            validAccountsLen += 1;
            totalXPRBalance = totalXPRBalance.plus(lineItem.XRPBalance);
        } else {
            invalidAccountsLen += 1;
            if (isValidNum) {
                invalidXPRBalance = invalidXPRBalance.plus(lineItem.XRPBalance);
            }
        }
    } 
    return {validAccounts, validAccountsLen, invalidAccountsLen, totalXPRBalance, invalidXPRBalance};
}

export function createFlareAirdropGenesisData
(parsedFile: LineItem[], validAccounts: validateRes, contingentPercentage: BigNumber,
conversionFactor: BigNumber, initialAirdropPercentage: BigNumber ):airdropGenesisRes{
    let processedAccountsLen:number = 0;
    let processedAccounts:string[] = [];
    let processedWei = new BigNumber(0);
    let seenFlareAddresses = new Set<string>();
    let flrAddDetail: {[name: string]: {balance: BigNumber, num: number} } = {};
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        if(validAccounts.validAccounts[lineIndex]){
            let lineItem = parsedFile[lineIndex];
            processedAccountsLen += 1
            let accBalance = new BigNumber(lineItem.XRPBalance);     
            accBalance = accBalance.multipliedBy(contingentPercentage);
            accBalance = accBalance.multipliedBy(conversionFactor);
            accBalance = accBalance.multipliedBy(initialAirdropPercentage);
            // To get from XPR to 6 decimal places to Wei (Flare to 18 decimal places)
            accBalance = accBalance.multipliedBy(TEN.pow(12));
            // rounding down to 0 decimal places
            accBalance = accBalance.dp(0, BigNumber.ROUND_FLOOR);
            processedWei = processedWei.plus(accBalance);
            if(seenFlareAddresses.has(lineItem.FlareAddress)){
                flrAddDetail[lineItem.FlareAddress].balance = flrAddDetail[lineItem.FlareAddress].balance.plus(accBalance);
                flrAddDetail[lineItem.FlareAddress].num += 1;
            }
            else {
                seenFlareAddresses.add(lineItem.FlareAddress);
                flrAddDetail[lineItem.FlareAddress] = {balance: accBalance, num: 1};
            }
        }
    }
    let accountsDistribution:number[] = [];
    for(let flrAdd of seenFlareAddresses){
        processedAccounts.push(`"${flrAdd.substring(2)}": {"balance": "0x${flrAddDetail[flrAdd].balance.toString(16)}" },`);
        if(accountsDistribution[flrAddDetail[flrAdd].num]){
            accountsDistribution[flrAddDetail[flrAdd].num] += 1;  
        } else {
            accountsDistribution[flrAddDetail[flrAdd].num] = 1;  
        }     
    }
    return {
        processedAccounts,
        processedAccountsLen,
        processedWei,
        accountsDistribution
    }
}