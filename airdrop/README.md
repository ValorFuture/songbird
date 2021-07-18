# Generating Airdrop files

This module is used to generate Airdrop genesis file and all files used within Distribution contract. It is build as a local yarn package that handles all the dependencies for running the script in locally closed environment.

It is also used as validator for Airdrop file from XPR snapshot on 12.12.2020

## Running the script

In order to run the script one must navigate into `/airdrop` folder.
From there one must initially install all dependencies:
```
yarn
``` 

One can then run the script 
```
yarn ts-node scripts/createAirdropGenesisDataFile.ts
``` 
Note that one needs to provide additional parameters that are described 
more thoroughly in input parameters section.

Required parameters:
```
--snapshot-file 
--genesis-file 
``` 

Optional parameters:
```
--override
``` 


## Input parameters
### Snapshot file

The file that holds snapshot data 

```
r11D6PPwznQcvNGCPbt7M27vguskJ826c,0x28BCD249FFD09D3FAF8D014683C5DB2A7CE36199,12953990545629
r11L3HhmYjTRVpueMwKZwPDeb6hBCSdBn,0x22577CC04C6EA5F0E1CDE6BD2663761549995BA0,207503719416
r12zYzJzTcf2j1BPsb5kUtZnLA1Wn7445,0x2A6687E2FDD6A66AC868AC62AD12C01245E72CBB,5593567199584
```

Fields are defined as:
1. XRP address
2. FLR address
3. XRP balance at time of snapshot, to six decimals, base 10.


## What does script do

Script does multiple things 

1. it generates genesis_scdev file 


In order to do that we do the following computations and checkups:

For each line of airdrop file we do the following computation:
```
FLR to distribute = XRP balance * conversion factor * initial airdrop percentage * contingent percentage

```
In order to generate a line for each account as such:
```
	      "ff50eF6F4b0568493175defa3655b10d68Bf41FB": {
	        "balance": "0x314dc6448d9338c15B0a00000000"
	      },
```

Doing so we:
1. Check validity of each XPR address
2. Check validity of each Flare address
3. Check that each balance is of an expected format
4. Maintain the amount of lines read (valid and invalid lines)
5. Check that there are no duplicate XPR addresses in input file
6. Join the duplicate Flare addresses ant their balances into one balance (assuming they came from two separate XPR addresses)
7. Maintain the total XPR read from input file 
8. Maintain the total airdropped wei, and check that it is close enough to expected airdrop wei number

## Future TODO's

We want to expand the script (or add new one) that generates either data file used in smart-contract repo
 to add accounts to Distribution contract, or create a transaction list with in order nonce logic, that can be 
 imported to some script run shortly after genesis that will fill the Distribution contract with necessary 
 data