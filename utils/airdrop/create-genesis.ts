function MakeRandomHex(iLen: number): string {
    let sRes: string = "";
    let sSource: string = "0123456789abcdef";

    for (let a: number = 0; (a < iLen); a++) {
        sRes += sSource[ Math.ceil( Math.random() * (sSource.length-1) ) ];
    }
    
    return sRes;
}

const yargs = require('yargs/yargs')

let argv = yargs(process.argv.slice(2)).options({
  a: { type: 'number', alias: 'addresses' , demandOption: true },
}).argv

let iAddresses : number = argv.a;

console.log("Creating " + iAddresses + " addresses..." );

import * as fs from 'fs';

//  address len 40
//  balance len 28
let sAddresses: string = "";
for (let a: number = 0; (a < iAddresses); a++) 
{   
    sAddresses +="          \"" + MakeRandomHex(40) + "\": {\n             \"balance\": \"0x" + MakeRandomHex(30) + "\"\n	      },\n";
}

fs.writeFile( "genesis.random.go" , sAddresses ,function(err) {
        if (err) {
            return console.error(err);
        }
        console.log("File created");
    } );