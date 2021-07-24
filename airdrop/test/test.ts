var assert = require('assert');

var processFile = require('../scripts/utils/processFile');
const LineItem = processFile.LineItem;
// import {createFlareAirdropGenesisData, validateFile, LineItem} from "../scripts/utils/processFile";

describe('Processing testing', function() {

  describe('Validating lines', function() {
    it('Should work just fine', function() {
      let data = [
        {XPRAddress:"r11D6PPwznQcvNGCPbt7M27vguskJ826c",FlareAddress:"0x28BCD249FFD09D3FAF8D014683C5DB2A7CE36199",
        XPRBalance:"1295399054562900000000"},
        {XPRAddress:"r11L3HhmYjTRVpueMwKZwPDeb6hBCSdBn",FlareAddress:"0x22577CC04C6EA5F0E1CDE6BD2663761549995BA0",
        XPRBalance:"20750371941600000000"},
        {XPRAddress:"r12zYzJzTcf2j1BPsb5kUtZnLA1Wn7445",FlareAddress:"0x2A6687E2FDD6A66AC868AC62AD12C01245E72CBB",
        XPRBalance:"559356719958400000000"},
        {XPRAddress:"r1398Fmwd1oYz8uUUbeQUE5axgXHjcfTZ",FlareAddress:"0x38EA655165CC077A36E1F1ED745C003DFE83875D",
        XPRBalance:"607905473445200000000"},
        {XPRAddress:"r13m9n9y7TVwFLfJnsMh1tGPRsXjMiaKh",FlareAddress:"0x8BA3B8041146FB6769D76A900826BE705B1D669E",
        XPRBalance:"38277375824800000000"},
        {XPRAddress:"r14f8Luu4dYKzNEwFYV2KfA74YZcWVS5F",FlareAddress:"0x158E1998458203B4824241B9BC178EA55C532A30",
        XPRBalance:"1410219697810000000000"},
        {XPRAddress:"r14iqdWmMQD1M7ski2a1oL2yoL8saBrgS",FlareAddress:"0xD4D3E94C6A2059C3166D4BD5A4421AF101394C7C",
        XPRBalance:"1034081916122500000000"},
        {XPRAddress:"r15BXLNhkFuUP2jztomyDvzsxVLzYw7Yh",FlareAddress:"0x61BA6F4C8165E031DA5443ACFCA9E804FBE993C4",
        XPRBalance:"47043919812400000000"},
        {XPRAddress:"r15aAVY2acncVcTkShfQQ6ycAQS2b4yfa",FlareAddress:"0xD4E690B5DD199B64DEA5B8FC08FC79A7F2CF7E76",
        XPRBalance:"20145987912400000000"},
        {XPRAddress:"r16DDq7D5kbh7mY6oUWk73RMd2pHA9CKv",FlareAddress:"0x6CDE1C841812C3820C8A61B9BE548F105DC15DDF",
        XPRBalance:"14133728449708000000000"}
      ];
      let loging = ""
      let processedFileData = processFile.validateFile(data,loging);

      assert.equal(processedFileData.validAccounts.length, 10);
      assert.equal(processedFileData.validAccounts[0], true);
      assert.equal(processedFileData.validAccounts[1], true);
      assert.equal(processedFileData.validAccounts[2], true);
      assert.equal(processedFileData.validAccountsLen, 10);
      assert.equal(processedFileData.invalidAccountsLen, 0);
    });
  });
});