const Web3Utils = require('web3-utils');
const RippleAPI = require('ripple-lib').RippleAPI;

const RippleApi = new RippleAPI({
    server: 'wss://s1.ripple.com' // Public rippled server hosted by Ripple, Inc.
  });

interface LineIntem {
    XPRAddress: string,
    FlareAdress: string,
    SparkBalanceWei: string
}

export function validateFile(parsedFile: LineIntem[]) {
    for(let lineIndex = 0; lineIndex < parsedFile.length; lineIndex++){
        let lineItem = parsedFile[lineIndex];
        if(!RippleApi.isValidAddress(lineItem.XPRAddress)){
            console.log(`Line ${lineIndex + 2}: XPR address is invalid`)
        }
        if(!Web3Utils.isAddress(lineItem.FlareAdress)){
            console.log(`Line ${lineIndex + 2}: Flare address is invalid`)
        }
        let numberBalance = +lineItem.SparkBalanceWei;
        if(isNaN(numberBalance)){
            console.log(`Line ${lineIndex + 2}: Balance is not a valid number`)
        }
    }
}

