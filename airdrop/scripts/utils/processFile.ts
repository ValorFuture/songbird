const Web3Utils = require('web3-utils');
const RippleAPI = require('ripple-lib').RippleAPI;
import * as fs from 'fs';
import { isBaseTenNumber, writeError } from './utils';
import BigNumber from "bignumber.js";
import { removeUndefined } from 'ripple-lib/dist/npm/common';

const TEN = new BigNumber(10);
const indexOffset = 2;
BigNumber.config({ ROUNDING_MODE: BigNumber.ROUND_FLOOR, DECIMAL_PLACES: 20 })

const RippleApi = new RippleAPI({
    server: 'wss://s1.ripple.com' // Public rippled server hosted by Ripple, Inc.
  });

export interface LineItem {
    XRPAddress: string,
    FlareAddress: string,
    XRPBalance: string,
    FlareBalance: string
}
export interface ProcessedAccount {
    FlareAddress: string,
    FlareBalance: string
}
interface validateRes {
    validAccounts: boolean[],
    validAccountsLen: number,
    invalidAccountsLen: number,
    totalXRPBalance: BigNumber,
    invalidXRPBalance: BigNumber,
    totalFLRBalance: BigNumber,
    invalidFLRBalance: BigNumber,
    lineErrors: number
}
interface airdropGenesisRes {
    processedAccounts: ProcessedAccount[],
    processedAccountsLen: number,
    processedWei: BigNumber,
    accountsDistribution: number[]
}

export function validateFile(parsedFile: LineItem[], logFile: string, logConsole: boolean = true):validateRes {
    let validAccountsLen:number = 0;
    let validAccounts:boolean[] = [];
    let invalidAccountsLen:number = 0;
    let lineErrors = 0;
    let seenXRPAddresses = new Set();
    let totalXRPBalance = new BigNumber(0);
    let invalidXRPBalance = new BigNumber(0);
    let totalFLRBalance = new BigNumber(0);
    let invalidFLRBalance = new BigNumber(0);
    let seenXRPAddressesDetail: {[name: string]: number } = {};
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let lineItem = parsedFile[lineIndex];
        let isValid = true;
        let isValidXRP = true;
        let isValidFLR = true;
        let readableIndex = lineIndex + indexOffset;
        if(!RippleApi.isValidAddress(lineItem.XRPAddress)){
            if(logConsole) console.log(`Line ${readableIndex}: XRP address is invalid ${lineItem.XRPAddress}`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: XRP address is invalid ${lineItem.XRPAddress}\n`);
            isValid = false;
            lineErrors += 1;
        }
        if(seenXRPAddresses.has(lineItem.XRPAddress)){
            // We have already seen this XRP address
            if(logConsole) console.log(`Line ${readableIndex}: XRP address is duplicate of line ${seenXRPAddressesDetail[lineItem.XRPAddress]}`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: XRP address is duplicate of line ${seenXRPAddressesDetail[lineItem.XRPAddress]}\n`);
            isValid = false;
            lineErrors += 1;
        }
        if(!seenXRPAddresses.has(lineItem.XRPAddress)){
            seenXRPAddresses.add(lineItem.XRPAddress);
            seenXRPAddressesDetail[lineItem.XRPAddress] = lineIndex;
        } 
        if(!Web3Utils.isAddress(lineItem.FlareAddress)){
            if(logConsole) console.log(`Line ${readableIndex}: Flare address is invalid ${lineItem.FlareAddress}`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: Flare address is invalid ${lineItem.FlareAddress}\n`);
            isValid = false;
            lineErrors += 1;
        }
        if(!isBaseTenNumber(lineItem.XRPBalance)){
            if(logConsole) console.log(`Line ${readableIndex}: XRP Balance is not a valid number`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: XRP Balance is not a valid number \n`);
            isValid = false;
            lineErrors += 1;
            isValidXRP = false;
        }
        if(!isBaseTenNumber(lineItem.FlareBalance)){
            if(logConsole) console.log(`Line ${readableIndex}: FLR Balance is not a valid number`);
            fs.appendFileSync(logFile, `Line ${readableIndex}: FLR Balance is not a valid number \n`);
            isValid = false;
            lineErrors += 1;
            isValidFLR = false;
        }
        validAccounts[lineIndex] = isValid;
        if(isValid){
            validAccountsLen += 1;
            totalXRPBalance = totalXRPBalance.plus(lineItem.XRPBalance);
            totalFLRBalance = totalFLRBalance.plus(lineItem.FlareBalance);
        } else {
            invalidAccountsLen += 1;
            if (isValidXRP) {
                invalidXRPBalance = invalidXRPBalance.plus(lineItem.XRPBalance);
            }
            if (isValidFLR) {
                invalidFLRBalance = invalidFLRBalance.plus(lineItem.FlareBalance);
            }
        }
    } 
    return {
        validAccounts,
        validAccountsLen,
        invalidAccountsLen,
        totalXRPBalance,
        invalidXRPBalance,
        totalFLRBalance,
        invalidFLRBalance,
        lineErrors
    };
}

export function createFlareAirdropGenesisData
(parsedFile: LineItem[],
validAccounts: validateRes,
contingentPercentage: BigNumber,
conversionFactor: BigNumber,
initialAirdropPercentage: BigNumber,
logFile: string, 
logConsole: boolean = true):airdropGenesisRes{
    let processedAccountsLen:number = 0;
    let processedAccounts:ProcessedAccount[] = [];
    let processedWei = new BigNumber(0);
    let seenFlareAddresses = new Set<string>();
    let flrAddDetail: {[name: string]: {balance: BigNumber, num: number} } = {};
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let readableIndex = lineIndex + indexOffset;
        if(validAccounts.validAccounts[lineIndex]){
            let lineItem = parsedFile[lineIndex];
            processedAccountsLen += 1
            let accBalance = new BigNumber(lineItem.XRPBalance);  
            accBalance = accBalance.multipliedBy(conversionFactor);  
            let expectedBalance = new BigNumber(lineItem.FlareBalance);
            // Check that balances are calculated properly
            if(!accBalance.multipliedBy(TEN.pow(12)).isEqualTo(expectedBalance)){
                if(logConsole) console.log(`Line ${readableIndex}: Flare balance error`);
                fs.appendFileSync(logFile, `Line ${readableIndex}: Flare balance error \n`);
            }
            // Calculate account balance 
            accBalance = accBalance.multipliedBy(contingentPercentage);
            accBalance = accBalance.multipliedBy(initialAirdropPercentage);
            // To get from XRP to 6 decimal places to Wei (Flare to 18 decimal places)
            accBalance = accBalance.multipliedBy(TEN.pow(12));
            // rounding down to 0 decimal places
            accBalance = accBalance.dp(0, BigNumber.ROUND_FLOOR);
            // Total Wei book keeping
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
        processedAccounts.push(
            {
                FlareAddress: flrAdd,
                FlareBalance: flrAddDetail[flrAdd].balance.toString(16)
            }
        )
        // processedAccounts.push(`"${flrAdd.substring(2)}": {"balance": "0x${flrAddDetail[flrAdd].balance.toString(16)}" },`);
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