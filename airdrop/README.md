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

