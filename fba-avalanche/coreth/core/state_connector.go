package core

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetMinReserve(blockNumber *big.Int) *big.Int {
	switch {
	default:
		minReserve, _ := new(big.Int).SetString("1000000000000000000000000", 10)
		return minReserve
	}
}

func GetStateConnectorGasDivisor(blockNumber *big.Int) uint64 {
	switch {
	default:
		return 100
	}
}

func GetKeeperGasMultiplier(blockNumber *big.Int) uint64 {
	switch {
	default:
		return 100
	}
}

func GetMaxAllowedChains(blockNumber *big.Int) uint32 {
	switch {
	default:
		return 4
	}
}

func GetGovernanceContractAddr(blockNumber *big.Int) string {
	switch {
	default:
		return "0xfffEc6C83c8BF5c3F4AE0cCF8c45CE20E4560BD7"
	}
}

func GetStateConnectorContractAddr(blockNumber *big.Int) string {
	switch {
	default:
		return "0x1000000000000000000000000000000000000001"
	}
}

func GetSystemTriggerContractAddr(blockNumber *big.Int) string {
	switch {
	default:
		return "0x1000000000000000000000000000000000000002"
	}
}

func GetInflationContractAddr(blockNumber *big.Int) string {
	switch {
	default:
		return GetSystemTriggerContractAddr(blockNumber)
	}
}

func GetProveClaimPeriodFinalitySelector(blockNumber *big.Int) []byte {
	switch {
	default:
		return []byte{0xa5, 0x7d, 0x0e, 0x25}
	}
}

func GetProvePaymentFinalitySelector(blockNumber *big.Int) []byte {
	switch {
	default:
		return []byte{0x38, 0x84, 0x92, 0xdd}
	}
}

func GetDisprovePaymentFinalitySelector(blockNumber *big.Int) []byte {
	switch {
	default:
		return []byte{0x7f, 0x58, 0x24, 0x32}
	}
}

func GetSystemTriggerSelector(blockNumber *big.Int) []byte {
	switch {
	default:
		return []byte{0x7f, 0xec, 0x8d, 0x38}
	}
}

var (
	tr = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
	}
	client = &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
)

type StateHashes struct {
	Hashes []string `json:"hashes"`
}

// =======================================================
// XRP
// =======================================================

type GetXRPBlockRequestParams struct {
	LedgerIndex  uint64 `json:"ledger_index"`
	Full         bool   `json:"full"`
	Accounts     bool   `json:"accounts"`
	Transactions bool   `json:"transactions"`
	Expand       bool   `json:"expand"`
	OwnerFunds   bool   `json:"owner_funds"`
}
type GetXRPBlockRequestPayload struct {
	Method string                     `json:"method"`
	Params []GetXRPBlockRequestParams `json:"params"`
}
type CheckXRPErrorResponse struct {
	Error string `json:"error"`
}
type GetXRPBlockResponse struct {
	LedgerHash  string `json:"ledger_hash"`
	LedgerIndex int    `json:"ledger_index"`
	Validated   bool   `json:"validated"`
}

func GetXRPBlock(ledger uint64, chainURL string) (string, bool) {
	data := GetXRPBlockRequestPayload{
		Method: "ledger",
		Params: []GetXRPBlockRequestParams{
			GetXRPBlockRequestParams{
				LedgerIndex:  ledger,
				Full:         false,
				Accounts:     false,
				Transactions: false,
				Expand:       false,
				OwnerFunds:   false,
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", true
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", chainURL, body)
	if err != nil {
		return "", true
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", true
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", true
		}
		var checkErrorResp map[string]CheckXRPErrorResponse
		err = json.Unmarshal(respBody, &checkErrorResp)
		if err != nil {
			return "", true
		}
		if checkErrorResp["result"].Error != "" {
			return "", true
		}
		var jsonResp map[string]GetXRPBlockResponse
		err = json.Unmarshal(respBody, &jsonResp)
		if err != nil {
			return "", true
		}
		if !jsonResp["result"].Validated {
			return "", true
		}
		return jsonResp["result"].LedgerHash, false
	}
	return "", true
}

func ProveClaimPeriodFinalityXRP(checkRet []byte, chainURL string) (bool, bool) {
	ledger := binary.BigEndian.Uint64(checkRet[56:64])
	ledgerHashString, err := GetXRPBlock(ledger, chainURL)
	if err {
		return false, true
	}
	if ledgerHashString != "" && bytes.Equal(crypto.Keccak256([]byte(ledgerHashString)), checkRet[96:128]) {
		return true, false
	}
	return false, false
}

type GetXRPTxRequestParams struct {
	Transaction string `json:"transaction"`
	Binary      bool   `json:"binary"`
}
type GetXRPTxRequestPayload struct {
	Method string                  `json:"method"`
	Params []GetXRPTxRequestParams `json:"params"`
}
type GetXRPTxResponse struct {
	Account         string      `json:"Account"`
	Amount          interface{} `json:"Amount"`
	Destination     string      `json:"Destination"`
	DestinationTag  int         `json:"DestinationTag"`
	TransactionType string      `json:"TransactionType"`
	Hash            string      `json:"hash"`
	InLedger        int         `json:"inLedger"`
	Flags           int         `json:"Flags"`
	Validated       bool        `json:"validated"`
}

type GetXRPTxIssuedCurrency struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"`
	Value    string `json:"value"`
}

func GetXRPTx(txHash string, latestAvailableLedger uint64, chainURL string) ([]byte, uint64, bool) {
	data := GetXRPTxRequestPayload{
		Method: "tx",
		Params: []GetXRPTxRequestParams{
			GetXRPTxRequestParams{
				Transaction: txHash,
				Binary:      false,
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return []byte{}, 0, true
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", chainURL, body)
	if err != nil {
		return []byte{}, 0, true
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, 0, true
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, 0, true
		}
		var checkErrorResp map[string]CheckXRPErrorResponse
		err = json.Unmarshal(respBody, &checkErrorResp)
		if err != nil {
			return []byte{}, 0, true
		}
		respErrString := checkErrorResp["result"].Error
		if respErrString != "" {
			if respErrString == "amendmentBlocked" ||
				respErrString == "failedToForward" ||
				respErrString == "invalid_API_version" ||
				respErrString == "noClosed" ||
				respErrString == "noCurrent" ||
				respErrString == "noNetwork" ||
				respErrString == "tooBusy" {
				return []byte{}, 0, true
			} else {
				return []byte{}, 0, false
			}
		}
		var jsonResp map[string]GetXRPTxResponse
		err = json.Unmarshal(respBody, &jsonResp)
		if err != nil {
			return []byte{}, 0, false
		}
		if jsonResp["result"].TransactionType == "Payment" {
			inLedger := uint64(jsonResp["result"].InLedger)
			if jsonResp["result"].Flags != 131072 && inLedger > 0 && inLedger < latestAvailableLedger && jsonResp["result"].Validated {
				var amount uint64
				amountInterface := jsonResp["result"].Amount
				amountType := reflect.TypeOf(amountInterface)
				var currency string
				if amountType.Name() == "string" {
					amount, err = strconv.ParseUint(jsonResp["result"].Amount.(string), 10, 64)
					if err != nil {
						return []byte{}, 0, false
					}
					currency = "XRP"
				} else {
					amountStruct, err := json.Marshal(amountInterface)
					if err != nil {
						return []byte{}, 0, false
					}
					var issuedCurrencyResp GetXRPTxIssuedCurrency
					err = json.Unmarshal(amountStruct, &issuedCurrencyResp)
					if err != nil {
						return []byte{}, 0, false
					}
					amountBigFloat, _, err := big.ParseFloat(issuedCurrencyResp.Value, 10, 256, big.ToZero)
					if err != nil {
						return []byte{}, 0, false
					}
					amountBigFloat.Mul(amountBigFloat, big.NewFloat(float64(1000000)))
					amount, _ = amountBigFloat.Uint64()
					currency = issuedCurrencyResp.Currency + issuedCurrencyResp.Issuer
				}
				txIdHash := crypto.Keccak256([]byte(jsonResp["result"].Hash))
				sourceHash := crypto.Keccak256([]byte(jsonResp["result"].Account))
				destinationHash := crypto.Keccak256([]byte(jsonResp["result"].Destination))
				destinationTagHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(jsonResp["result"].DestinationTag))), 32))
				amountHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(amount))), 32))
				currencyHash := crypto.Keccak256([]byte(currency))
				return crypto.Keccak256(txIdHash, sourceHash, destinationHash, destinationTagHash, amountHash, currencyHash), inLedger, false
			} else {
				return []byte{}, 0, false
			}
		} else {
			return []byte{}, 0, false
		}
	}
	return []byte{}, 0, true
}

func ProvePaymentFinalityXRP(checkRet []byte, chainURL string) (bool, bool) {
	paymentHash, inLedger, err := GetXRPTx(string(checkRet[192:]), binary.BigEndian.Uint64(checkRet[88:96]), chainURL)
	if !err {
		if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger == binary.BigEndian.Uint64(checkRet[56:64]) {
			return true, false
		}
		return false, false
	}
	return false, true
}

func DisprovePaymentFinalityXRP(checkRet []byte, chainURL string) (bool, bool) {
	paymentHash, inLedger, err := GetXRPTx(string(checkRet[192:]), binary.BigEndian.Uint64(checkRet[88:96]), chainURL)
	if !err {
		if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger > binary.BigEndian.Uint64(checkRet[56:64]) {
			return true, false
		} else if len(paymentHash) == 0 {
			return true, false
		}
		return false, false
	}
	return false, true
}

func ProveXRP(sender common.Address, blockNumber *big.Int, functionSelector []byte, checkRet []byte, evmAddresses string, chainURL string) (bool, bool) {
	if bytes.Equal(functionSelector, GetProveClaimPeriodFinalitySelector(blockNumber)) {
		for _, evmAddress := range strings.Split(evmAddresses, ",") {
			if len(evmAddress) == 45 {
				if evmAddress[:3] == "xrp" && common.HexToAddress(evmAddress[3:]) == sender {
					return ProveClaimPeriodFinalityXRP(checkRet, chainURL)
				}
			}
		}
		return false, false
	} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockNumber)) {
		return ProvePaymentFinalityXRP(checkRet, chainURL)
	} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockNumber)) {
		return DisprovePaymentFinalityXRP(checkRet, chainURL)
	}
	return false, false
}

// =======================================================
// LTC
// =======================================================

func ProveClaimPeriodFinalityLTC(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func ProvePaymentFinalityLTC(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func DisprovePaymentFinalityLTC(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func ProveLTC(sender common.Address, blockNumber *big.Int, functionSelector []byte, checkRet []byte, evmAddresses string, chainURL string) (bool, bool) {
	if bytes.Equal(functionSelector, GetProveClaimPeriodFinalitySelector(blockNumber)) {
		for _, evmAddress := range strings.Split(evmAddresses, ",") {
			if len(evmAddress) == 45 {
				if evmAddress[:3] == "ltc" && common.HexToAddress(evmAddress[3:]) == sender {
					return ProveClaimPeriodFinalityLTC(checkRet, chainURL)
				}
			}
		}
		return false, false
	} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockNumber)) {
		return ProvePaymentFinalityLTC(checkRet, chainURL)
	} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockNumber)) {
		return DisprovePaymentFinalityLTC(checkRet, chainURL)
	}
	return false, false
}

// =======================================================
// XLM
// =======================================================

func ProveClaimPeriodFinalityXLM(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func ProvePaymentFinalityXLM(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func DisprovePaymentFinalityXLM(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func ProveXLM(sender common.Address, blockNumber *big.Int, functionSelector []byte, checkRet []byte, evmAddresses string, chainURL string) (bool, bool) {
	if bytes.Equal(functionSelector, GetProveClaimPeriodFinalitySelector(blockNumber)) {
		for _, evmAddress := range strings.Split(evmAddresses, ",") {
			if len(evmAddress) == 45 {
				if evmAddress[:3] == "xlm" && common.HexToAddress(evmAddress[3:]) == sender {
					return ProveClaimPeriodFinalityXLM(checkRet, chainURL)
				}
			}
		}
		return false, false
	} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockNumber)) {
		return ProvePaymentFinalityXLM(checkRet, chainURL)
	} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockNumber)) {
		return DisprovePaymentFinalityXLM(checkRet, chainURL)
	}
	return false, false
}

// =======================================================
// DOGE
// =======================================================

func ProveClaimPeriodFinalityDOGE(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func ProvePaymentFinalityDOGE(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func DisprovePaymentFinalityDOGE(checkRet []byte, chainURL string) (bool, bool) {
	return true, false
}

func ProveDOGE(sender common.Address, blockNumber *big.Int, functionSelector []byte, checkRet []byte, evmAddresses string, chainURL string) (bool, bool) {
	if bytes.Equal(functionSelector, GetProveClaimPeriodFinalitySelector(blockNumber)) {
		for _, evmAddress := range strings.Split(evmAddresses, ",") {
			if len(evmAddress) == 45 {
				if evmAddress[:3] == "dog" && common.HexToAddress(evmAddress[3:]) == sender {
					return ProveClaimPeriodFinalityDOGE(checkRet, chainURL)
				}
			}
		}
		return false, false
	} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockNumber)) {
		return ProvePaymentFinalityDOGE(checkRet, chainURL)
	} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockNumber)) {
		return DisprovePaymentFinalityDOGE(checkRet, chainURL)
	}
	return false, false
}

// =======================================================
// Common
// =======================================================

func ProveChain(sender common.Address, blockNumber *big.Int, functionSelector []byte, checkRet []byte, chainId uint32, evmAddresses string, chainURL string) (bool, bool) {
	switch chainId {
	case 0:
		return ProveXRP(sender, blockNumber, functionSelector, checkRet, evmAddresses, chainURL)
	case 1:
		return ProveLTC(sender, blockNumber, functionSelector, checkRet, evmAddresses, chainURL)
	case 2:
		return ProveXLM(sender, blockNumber, functionSelector, checkRet, evmAddresses, chainURL)
	case 3:
		return ProveDOGE(sender, blockNumber, functionSelector, checkRet, evmAddresses, chainURL)
	default:
		return false, true
	}
}

func ReadChain(sender common.Address, blockNumber *big.Int, functionSelector []byte, checkRet []byte, alertURLs string, evmAddresses string, chainURLs []string) bool {
	chainId := binary.BigEndian.Uint32(checkRet[28:32])
	if uint32(len(chainURLs)) <= chainId {
		// This is already checked at avalanchego/main/params.go on launch, but a fail-safe
		// is included here regardless for increased coverage
		return false
	}
	for {
		for _, chainURL := range strings.Split(chainURLs[chainId], ",") {
			if chainURL != "" {
				verified, err := ProveChain(sender, blockNumber, functionSelector, checkRet, chainId, evmAddresses, chainURL)
				if !verified && err {
					continue
				} else {
					return verified
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// Verify proof against underlying chain
func StateConnectorCall(sender common.Address, blockNumber *big.Int, functionSelector []byte, checkRet []byte, stateConnectorConfig []string) bool {
	return ReadChain(sender, blockNumber, functionSelector, checkRet, stateConnectorConfig[2], stateConnectorConfig[3], stateConnectorConfig[4:])
}
