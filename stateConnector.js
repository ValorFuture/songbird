'use strict';
process.env.NODE_ENV = 'production';
const Web3 = require('web3');
const web3 = new Web3();
const Tx = require('ethereumjs-tx').Transaction;
const Common = require('ethereumjs-common').default;
const fs = require('fs');
const express = require('express');
const app = express();
const { MerkleTree } = require('merkletreejs');
const keccak256 = require('keccak256');

const minFee = 1;
var config;
var customCommon;
var chainAPI;
var stateConnector;
var claimsInProgress = false;

// ===============================================================
// XRPL Specific Functions
// ===============================================================

const RippleAPI = require('ripple-lib').RippleAPI;
const RippleKeys = require('ripple-keypairs');

async function xrplProcessLedgers(payloads, genesisLedger, claimPeriodIndex, claimPeriodLength, ledger, registrationFee) {
	console.log('\nRetrieving XRPL state from ledgers:', ledger, 'to', genesisLedger + (claimPeriodIndex+1)*claimPeriodLength - 1);
	async function xrplProcessLedger(currLedger) {
		const command = 'ledger';
		const params = {
			'ledger_index': currLedger,
			'binary': false,
			'full': false,
			'accounts': false,
			'transactions': true,
			'expand': true,
			'owner_funds': false
		};
		return chainAPI.request(command, params)
		.then(response => {
			async function responseIterate(response) {
				async function transactionIterate(item, i, numTransactions) {
					if (item.TransactionType == 'Payment' && typeof item.Amount == 'string' && item.metaData.TransactionResult == 'tesSUCCESS') {
						const prevLength = payloads.length;
						const leafPromise = new Promise((resolve, reject) => {
							var destinationTag;
							if (!("DestinationTag" in item)) {
								destinationTag = 0;
							} else {
								destinationTag = item.DestinationTag;
							}
							console.log('chainId: \t\t', '0', '\n',
								'ledger: \t\t', response.ledger.seqNum, '\n',
								'txId: \t\t\t', item.hash, '\n',
								'source: \t\t', item.Account, '\n',
								'destination: \t\t', item.Destination, '\n',
								'destinationTag: \t', String(destinationTag), '\n',
								'amount: \t\t', parseInt(item.metaData.delivered_amount), '\n');
							const chainIdHash = web3.utils.soliditySha3('0');
							const ledgerHash = web3.utils.soliditySha3(response.ledger.seqNum);
							const txHash = web3.utils.soliditySha3(item.hash);
							const accountsHash = web3.utils.soliditySha3(item.Account, item.Destination, destinationTag);
							const amountHash = web3.utils.soliditySha3(item.metaData.delivered_amount);
							const leafHash = web3.utils.soliditySha3(chainIdHash, ledgerHash, txHash, accountsHash, amountHash);
							resolve(leafHash);
						}).catch(processFailure)
						return await leafPromise.then(newPayload => {
							payloads[payloads.length] = newPayload;
							if (payloads.length == prevLength + 1) {
								if (i+1 < numTransactions) {
									return transactionIterate(response.ledger.transactions[i+1], i+1, numTransactions);
								} else {
									return checkResponseCompletion(response);
								}
							} else {
								return processFailure("Unable to append payload:", tx.hash);
							}
						}).catch(error => {
							return processFailure("Unable to intepret payload:", error, tx.hash);
						})
					} else {
						if (i+1 < numTransactions) {
							return transactionIterate(response.ledger.transactions[i+1], i+1, numTransactions);
						} else {
							return checkResponseCompletion(response);
						}
					}
				}
				async function checkResponseCompletion(response) {
					if (chainAPI.hasNextPage(response) == true) {
						chainAPI.requestNextPage(command, params, response)
						.then(next_response => {
							responseIterate(next_response);
						})
					} else if (parseInt(currLedger)+1 < genesisLedger + (claimPeriodIndex+1)*claimPeriodLength) {
						return xrplProcessLedger(parseInt(currLedger)+1);
					} else {
						var root;
						if (payloads.length > 0) {
							const tree = new MerkleTree(payloads, keccak256, {sort: true});
							root = tree.getHexRoot();
						} else {
							root = "0x0000000000000000000000000000000000000000000000000000000000000000";
						}
						console.log('Num Payloads:\t\t', payloads.length);
						return registerClaimPeriod(0, genesisLedger + (claimPeriodIndex+1)*claimPeriodLength, claimPeriodIndex, root, registrationFee);
					}
				}
				if (response.ledger.transactions.length > 0) {
					return transactionIterate(response.ledger.transactions[0], 0, response.ledger.transactions.length);
				} else {
					return checkResponseCompletion(response);
				}
			}
			responseIterate(response);
		})
		.catch(error => {
			processFailure(error);
		})
	}
	return xrplProcessLedger(ledger);
}

async function xrplConfig() {
	let rawConfig = fs.readFileSync('config/config.json');
	config = JSON.parse(rawConfig);
	chainAPI = new RippleAPI({
	  server: config.chains[0].url,
	  timeout: 60000
	});
	web3.setProvider(new web3.providers.HttpProvider(config.flare.url));
	web3.eth.handleRevert = true;
	customCommon = Common.forCustomChain('ropsten',
						{
							name: 'coston',
							networkId: config.flare.chainId,
							chainId: config.flare.chainId,
						},
        				'petersburg',);
	chainAPI.on('connected', () => {
		return run(0);
	})
}

function xrplClaimProcessingCompleted(message) {
	chainAPI.disconnect().catch(processFailure)
	.then(() => {
		console.log(message);
		setTimeout(() => {return process.exit()}, 2500);
	})
}

async function xrplConnectRetry(error) {
	console.log('XRPL connecting...')
	sleep(1000).then(() => {
		chainAPI.connect().catch(xrplConnectRetry);
	})
}

// ===============================================================
// Chain Invariant Functions
// ===============================================================

async function run(chainId) {
	console.log('\n\x1b[34mState Connector System connected at', Date(Date.now()).toString(), '\x1b[0m' );
	stateConnector.methods.getlatestIndex(parseInt(chainId)).call({
		from: config.stateConnector.address,
		gas: config.flare.gas,
		gasPrice: config.flare.gasPrice
	}).catch(processFailure)
	.then(result => {
		return [parseInt(result.genesisLedger), parseInt(result.finalisedClaimPeriodIndex), parseInt(result.claimPeriodLength), 
		parseInt(result.finalisedLedgerIndex), parseInt(result._registrationFee)];
	})
	.then(result => {
		if (chainId == 0) {
			chainAPI.getLedgerVersion().catch(processFailure)
			.then(sampledLedger => {
				console.log("Finalised claim period:\t\x1b[33m", result[1]-1, 
					"\n\x1b[0mFinalised Ledger Index:\t\x1b[33m", result[3], '\n\x1b[0mCurrent Ledger Index:\t\x1b[33m', sampledLedger);
				if (sampledLedger > result[0] + (result[1]+1)*result[2]) {
					return xrplProcessLedgers([], result[0], result[1], result[2], result[3], result[4]);
				} else {
					return xrplClaimProcessingCompleted('Reached latest state, waiting for new ledgers.');
				}
			})
		} else {
			return processFailure('Invalid chainId.');
		}
	})
}

async function registerClaimPeriod(chainId, ledger, claimPeriodIndex, claimPeriodHash, registrationFee) {
	stateConnector.methods.checkFinality(
					parseInt(chainId),
					claimPeriodIndex).call({
		from: config.stateConnector.address,
		gas: config.flare.gas,
		gasPrice: config.flare.gasPrice
	}).catch(processFailure)
	.then(result => {
		console.log('Claim period:\t\t\x1b[33m', claimPeriodIndex, '\x1b[0m\nclaimPeriodHash:\t\x1b[33m', claimPeriodHash, '\x1b[0m');
		if (result == true) {
			if (chainId == 0) {
				return xrplClaimProcessingCompleted('Latest claim period already registered, waiting for new ledgers.');
			} else {
				return processFailure('Invalid chainId.');
			}
		} else {
			web3.eth.getTransactionCount(config.stateConnector.address)
			.then(nonce => {
				return [stateConnector.methods.registerClaimPeriod(
							chainId,
							ledger,
							claimPeriodIndex,
							claimPeriodHash).encodeABI(), nonce];
			})
			.then(txData => {
				var rawTx = {
					nonce: txData[1],
					gasPrice: web3.utils.toHex(parseInt(config.flare.gasPrice)),
					gas: web3.utils.toHex(config.flare.gas),
					to: stateConnector.options.address,
					from: config.stateConnector.address,
					value: parseInt(registrationFee),
					data: txData[0]
				};
				var tx = new Tx(rawTx, {common: customCommon});
				var key = Buffer.from(config.stateConnector.privateKey, 'hex');
				tx.sign(key);
				var serializedTx = tx.serialize();
				const txHash = web3.utils.sha3(serializedTx);

				console.log('Delivering transaction:\t\x1b[33m', txHash, '\x1b[0m');
				return web3.eth.getTransaction(txHash)
				.then(txResult => {
					if (txResult == null) {
						web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex'))
						.on('receipt', receipt => {
							if (receipt.status == false) {
								return processFailure('receipt.status == false');
							} else {
								console.log('Transaction finalised:\t \x1b[33m' + receipt.transactionHash + '\x1b[0m');
								return setTimeout(() => {return run(0)}, 5000);
							}
						})
						.on('error', error => {
							return processFailure(error);
						});
					} else {
						return processFailure('txResult != null');
					}
				})
			})
		}
	})
}

async function contract() {
	// Read the compiled contract code
	let source = fs.readFileSync("solidity/stateConnector.json");
	let contracts = JSON.parse(source)["contracts"];
	// ABI description as JSON structure
	let abi = JSON.parse(contracts['stateConnector.sol:stateConnector'].abi);
	// Create Contract proxy class
	stateConnector = new web3.eth.Contract(abi);
	// Smart contract EVM bytecode as hex
	stateConnector.options.data = '0x' + contracts['stateConnector.sol:stateConnector'].bin;
	stateConnector.options.from = config.stateConnector.address;
	stateConnector.options.address = config.stateConnector.contract;
}

async function processFailure(error) {
	console.error('error:', error);
	setTimeout(() => {return process.exit()}, 2500);
}

async function updateClaimsInProgress(status) {
	claimsInProgress = status;
	return claimsInProgress;
}

async function sleep(ms) {
	return new Promise((resolve) => {
		setTimeout(resolve, ms);
	});
}

app.get('/stateConnector', (req, res) => {
	if (claimsInProgress == true) {
		res.status(200).send('Claims already being processed.').end();
	} else {
		updateClaimsInProgress(true)
		.then(result => {
			if (result == true) {
				res.status(200).send('State Connector initiated.').end();
				const chainId = parseInt(process.argv[2]);
				if (chainId == 0) {
					xrplConfig().catch(processFailure)
					.then(() => {
						return contract().catch(processFailure);
					})
					.then(() => {
						return chainAPI.connect().catch(xrplConnectRetry);
					})
				} else {
					processFailure('Invalid chainId');
				}
			} else {
				return processFailure('Error updating claimsInProgress.');
			}
		})
	}
});
// Start the server
const PORT = process.env.PORT || 8080+parseInt(process.argv[2]);
app.listen(PORT, () => {
});
module.exports = app;