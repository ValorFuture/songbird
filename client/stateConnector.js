'use strict';
process.env.NODE_ENV = 'production';
const Web3 = require('web3');
const web3 = new Web3();
const Tx = require('ethereumjs-tx').Transaction;
const Common = require('ethereumjs-common').default;
const fs = require('fs');
const fetch = require('node-fetch');
const express = require('express');
const app = express();

const stateConnectorContract = "0x1000000000000000000000000000000000000001";
const chains = {
	'btc': {
		chainId: 0
	},
	'ltc': {
		chainId: 1
	},
	'doge': {
		chainId: 2
	},
	'xrp': {
		chainId: 3
	},
	'xlm': {
		chainId: 4
	}
};

var active,
	config,
	customCommon,
	stateConnector;

async function postData(url = '', data = {}) {
	const response = await fetch(url, {
	  method: 'POST',
	  headers: {
		'Content-Type': 'application/json'
	  },
	  body: JSON.stringify(data)
	});
	return response.json();
  }

// ===============================================================
// XRP Specific Items
// ===============================================================

async function xrplProcessLedger(genesisLedger, claimPeriodIndex, claimPeriodLength, isCommit) {
	console.log('\nRetrieving XRPL state hash from ledger:', genesisLedger + (claimPeriodIndex+1)*claimPeriodLength - 1);
	const currLedger = genesisLedger + (claimPeriodIndex+1)*claimPeriodLength - 1;
	const command = 'ledger';
	const params = [{
		'ledger_index': currLedger,
		'binary': false,
		'full': false,
		'accounts': false,
		'transactions': false,
		'expand': false,
		'owner_funds': false
	}];
	postData('https://xrpl.flare.network:443', { method: command, params: params })
	.then(data => {
		return proveClaimPeriodFinality(chains['xrp'].chainId, genesisLedger + (claimPeriodIndex+1)*claimPeriodLength, claimPeriodIndex, web3.utils.sha3(data.result.ledger_hash), isCommit);
	})
	.catch(error => {
		processFailure(error);
	})
}

// ===============================================================
// Chain Invariant Functions
// ===============================================================

async function run(chainId, minLedger) {
	console.log('\n\x1b[34mState Connector System connected at', Date(Date.now()).toString(), '\x1b[0m' );
	stateConnector.methods.getLatestIndex(parseInt(chainId)).call().catch(initialiseChains)
	.then(result => {
		if (result != undefined) {
			if (chainId == 3) {
				// chains.xrp.api.getLedgerVersion().catch(processFailure)
				const command = 'ledger';
				const params = [{
					'ledger_index': "validated",
					'binary': false,
					'full': false,
					'accounts': false,
					'transactions': false,
					'expand': false,
					'owner_funds': false
				}];
				postData('https://xrpl.flare.network:443', { method: command, params: params })
				.then(data => {
					console.log("Finalised claim period:\t\x1b[33m", parseInt(result.finalisedClaimPeriodIndex)-1, 
						"\n\x1b[0mFinalised Ledger Index:\t\x1b[33m", parseInt(result.finalisedLedgerIndex),
						"\n\x1b[0mCurrent Ledger Index:\t\x1b[33m", data.result.ledger_index,
						"\n\x1b[0mFinalised Timestamp:\t\x1b[33m", parseInt(result.finalisedTimestamp),
						"\n\x1b[0mCurrent Timestamp:\t\x1b[33m", parseInt(Date.now()/1000),
						"\n\x1b[0mDiff Avg (sec):\t\t\x1b[33m", parseInt(result.timeDiffAvg));
					const currTime = parseInt(Date.now()/1000);
					var deferTime;
					if (parseInt(result.finalisedTimestamp) > 0) {
						// Time to commit the proof
						if (parseInt(result.timeDiffAvg) < 60) {
							deferTime = parseInt(2*parseInt(result.timeDiffAvg)/3 - (currTime-parseInt(result.finalisedTimestamp)));
						} else {
							deferTime = parseInt(parseInt(result.timeDiffAvg) - (currTime-parseInt(result.finalisedTimestamp)) - 15);
						}
						if (deferTime > 0) {
							console.log("Not enough time elapsed since prior finality, deferring for", deferTime, "seconds.");
							return setTimeout(() => {run(chainId, minLedger)}, 1000*(deferTime+1));
						} else if (data.result.ledger_index >= parseInt(result.genesisLedger) + (parseInt(result.finalisedClaimPeriodIndex)+1)*parseInt(result.claimPeriodLength)) {
							return xrplProcessLedger(parseInt(result.genesisLedger), parseInt(result.finalisedClaimPeriodIndex), parseInt(result.claimPeriodLength), true);
						} else {
							return xrplClaimProcessingCompleted('Reached latest state, waiting for new ledgers.');
						}
					} else {
						// Time to reveal the proof
						if (currTime > parseInt(result.timeDiffAvg)) {
							return xrplProcessLedger(parseInt(result.genesisLedger), parseInt(result.finalisedClaimPeriodIndex), parseInt(result.claimPeriodLength), false);
						} else {
							deferTime = parseInt(result.timeDiffAvg)-currTime;
							console.log("Not enough time elapsed since proof commit, deferring for", deferTime, "seconds.");
							return setTimeout(() => {run(chainId, minLedger)}, 1000*(deferTime+1));
						}
					}
				})
			} else {
				return processFailure('Invalid chainId.');
			}
		}	
	})
}

async function proveClaimPeriodFinality(chainId, ledger, claimPeriodIndex, claimPeriodHash, isCommit) {
	stateConnector.methods.getClaimPeriodIndexFinality(
					parseInt(chainId),
					claimPeriodIndex).call({
		from: config.accounts[0].address,
		gas: config.flare.gas,
		gasPrice: config.flare.gasPrice
	}).catch(processFailure)
	.then(result => {
		console.log('\x1b[0mClaim period:\t\t\x1b[33m', claimPeriodIndex, '\x1b[0m\nProof reveal:\t\t\x1b[33m', !isCommit, '\x1b[0m\nclaimPeriodHash:\t\x1b[33m', claimPeriodHash, '\x1b[0m');
		if (result == true) {
			if (chainId == 3) {
				console.log('This claim period already registered.');
				setTimeout(() => {return process.exit()}, 5000);
			} else {
				return processFailure('Invalid chainId.');
			}
		} else {
			web3.eth.getTransactionCount(config.accounts[0].address)
			.then(nonce => {
				if (isCommit) {
					return [stateConnector.methods.proveClaimPeriodFinality(
						chainId,
						ledger,
						claimPeriodIndex,
						web3.utils.soliditySha3(config.accounts[0].address, claimPeriodHash)).encodeABI(), nonce];
				} else {
					return [stateConnector.methods.proveClaimPeriodFinality(
						chainId,
						ledger,
						claimPeriodIndex,
						claimPeriodHash).encodeABI(), nonce];
				}
			})
			.then(txData => {
				var rawTx = {
					nonce: txData[1],
					gasPrice: web3.utils.toHex(parseInt(config.flare.gasPrice)),
					gas: web3.utils.toHex(config.flare.gas),
					to: stateConnector.options.address,
					from: config.accounts[0].address,
					data: txData[0]
				};
				var tx = new Tx(rawTx, {common: customCommon});
				var key = Buffer.from(config.accounts[0].privateKey, 'hex');
				tx.sign(key);
				var serializedTx = tx.serialize();
				const txHash = web3.utils.sha3(serializedTx);

				console.log('Delivering transaction:\t\x1b[33m', txHash, '\x1b[0m');
				web3.eth.getTransaction(txHash)
				.then(txResult => {
					if (txResult == null) {
						web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'))
						.on('receipt', receipt => {
							if (receipt.status == false) {
								return processFailure('receipt.status == false');
							} else {
								console.log('Transaction delivered:\t \x1b[33m' + receipt.transactionHash + '\x1b[0m');
								return setTimeout(() => {run(chainId, ledger)}, 5000);
							}
						})
						.on('error', error => {
							return processFailure(error);
						});
					} else {
						console.log('Already waiting for this transaction to be delivered.');
						setTimeout(() => {return process.exit()}, 5000);
					}
				})
			})
		}
	})
}

async function initialiseChains() {
	console.log('Initialising chains');
	web3.eth.getTransactionCount(config.accounts[0].address)
	.then(nonce => {
		return [stateConnector.methods.initialiseChains().encodeABI(), nonce];
	})
	.then(contractData => {
		var rawTx = {
			nonce: contractData[1],
			gasPrice: web3.utils.toHex(config.flare.gasPrice),
			gas: web3.utils.toHex(config.flare.gas),
			chainId: config.flare.chainId,
			from: config.accounts[0].address,
			to: stateConnector.options.address,
			data: contractData[0]
		}
		var tx = new Tx(rawTx, {common: customCommon});
		var key = Buffer.from(config.accounts[0].privateKey, 'hex');
		tx.sign(key);
		var serializedTx = tx.serialize();
		
		web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'))
		.on('receipt', receipt => {
			if (receipt.status == false) {
				return processFailure('receipt.status == false');
			} else {
				console.log("State-connector chains initialised.");
				setTimeout(() => {return process.exit()}, 5000);
			}
		})
		.on('error', error => {
			processFailure(error);
		});
	}).catch(processFailure);
}

async function configure(chainId) {
	let rawConfig = fs.readFileSync('config.json');
	config = JSON.parse(rawConfig);
	// console.log(config);
	web3.setProvider(new web3.providers.HttpProvider(config.flare.url));
	web3.eth.handleRevert = true;
	customCommon = Common.forCustomChain('ropsten',
		{
			name: 'coston',
			networkId: config.flare.chainId,
			chainId: config.flare.chainId,
		},
		'petersburg',);
	web3.eth.getBalance(config.accounts[0].address)
	.then(balance => {
		console.log(balance);
		if (parseInt(web3.utils.fromWei(balance, "ether")) < 1000000) {
			console.log("Not enough FLR reserved in your account, need 1M FLR.");
			sleep(5000);
			process.exit();
		} else {
			// Read the compiled contract code
			let source = fs.readFileSync("../bin/contracts/StateConnector.json");
			let contract = JSON.parse(source);
			// Create Contract proxy class
			stateConnector = new web3.eth.Contract(contract.abi);
			// Smart contract EVM bytecode as hex
			stateConnector.options.data = '0x' + contract.deployedBytecode;
			stateConnector.options.from = config.accounts[0].address;
			stateConnector.options.address = stateConnectorContract;
			return run(chainId, 0);
		}
	})
}

async function processFailure(error) {
	console.error('error:', error);
	setTimeout(() => {return process.exit()}, 2500);
}


async function sleep(ms) {
	return new Promise((resolve) => {
		setTimeout(resolve, ms);
	});
}

setTimeout(() => {return process.exit()}, 600000);
app.get('/', (req, res) => {
	if ("prove" in req.query) {
		if (req.query.prove in chains) {
			if (active) {
				res.status(200).send('State Connector already active on this port.').end();
			} else {
				active = true;
				res.status(200).send('State Connector initiated.').end();
				return configure(chains[req.query.prove].chainId);
			}
		} else {
			res.status(404).send('Unknown chain.');
		}
	} else {
		res.status(200).send('Healthy.');
	}
});
// Start the server
const PORT = process.env.PORT || parseInt(process.argv[2]);
app.listen(PORT, () => {
});
module.exports = app;
