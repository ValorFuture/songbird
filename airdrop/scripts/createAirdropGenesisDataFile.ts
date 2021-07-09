const Web3Utils = require('web3-utils');


const address = '0xc1912fee45d61c87cc5ea59dae31190fffff232d'
const bl = Web3Utils.isAddress(address)
// let vara  = web3.utils.isAddress(address)

console.log(bl)

const a:number = 10;
const b:string = "Flare SC team is the best"
console.log(b);

const c = process.argv.slice(2)
console.log(c);