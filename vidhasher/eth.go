package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
	"math/big"
	"os"
)

var (
	EthAccount accounts.Account
	EthWallet  *hdwallet.Wallet
)

func etherum_dial() {
	var err error
	EthClient, err = ethclient.Dial(EthRpcAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if EthAccountStr != "" {
		maddr, err := common.NewMixedcaseAddressFromString(EthAccountStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		EthAccountAddr = maddr.Address()
	}
}

func make_wallet() {
	var err error

	EthChainId = big.NewInt(-1)
	_, ok := EthChainId.SetString(EthChainIdStr, 10)
	if ok != true {
		Log.FatalF("Failed to parse chain-id as integer")
	}

	EthWallet, err = hdwallet.NewFromMnemonic(EthMnemonic)
	if err != nil {
		Log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	EthAccount, err = EthWallet.Derive(path, false)
	if err != nil {
		Log.Fatal(err)
	}

	EthAccountSK, err = EthWallet.PrivateKey(EthAccount)
	if err != nil {
		Log.Fatal(err)
	}
}

func depoly_new_contract(cmd *cobra.Command, args []string) {
	etherum_dial()
	make_wallet()

	fmt.Printf("Using address: %s\n", EthAccount.Address.Hex())

	// Default [chash] parameter
	if len(args) < 2 {
		args = append(args, "multihash")
	}
	alg_phash := args[0]
	alg_chash := args[1]

	auth, err := bind.NewKeyedTransactorWithChainID(EthAccountSK, EthChainId)
	if err != nil {
		Log.FatalF("Failed to make transactor: %v", err)
	}
	contract_addr, tx, _, err := DeployHashDictionary(auth, EthClient, alg_phash, alg_chash)
	if err != nil {
		Log.FatalF("Failed to deploy a new HashDictionary contract: %v", err)
	}
	fmt.Printf("Contract pending deployment at: 0x%x\n", contract_addr)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
}

func make_new_wallet(cmd *cobra.Command, args []string) {
	var err error
	EthMnemonic, err = hdwallet.NewMnemonic(256)
	if err != nil {
		Log.Fatal(err)
	}
	make_wallet()

	fmt.Println(EthMnemonic)
}

func find_phash(cmd *cobra.Command, args []string) {
	etherum_dial()
	var err error

	query_str := args[0]
	query_hex := common.Hex2Bytes(query_str)

	dict, err := NewHashDictionary(common.HexToAddress(ContractAddr), EthClient)
	if err != nil {
		Log.FatalF("Failed to instantiate a HashDictionary contract: %v", err)
	}
	phash_alg, err := dict.PhashAlg(nil)
	if err != nil {
		Log.FatalF("Failed to retrieve phash_alg: %v", err)
	}
	chash_alg, err := dict.ChashAlg(nil)
	if err != nil {
		Log.FatalF("Failed to retrieve chash_alg: %v", err)
	}
	fmt.Printf("====================== HashDictioanry ======================\n")
	fmt.Printf("Contract address: %s\n", ContractAddr)
	fmt.Printf("pHash algorithm:  %s\n", phash_alg)
	fmt.Printf("cHash algorithm:  %s\n", chash_alg)
	fmt.Printf("========================== Search ==========================\n")
	fmt.Printf("phash = %s\n", query_str)
	fmt.Printf("========================== Results =========================\n")
	entries, err := dict.GetEntriesByPHash(nil, query_hex)
	for _, entry := range entries {
		fmt.Printf("chash = %s\n", common.Bytes2Hex(entry.Hpair.Chash))
		fmt.Printf("votes = %s YES / %s NO\n", entry.VotesRight.String(), entry.VotesWrong.String())
	}
}

func vote_entry(phash, chash string, vote uint8) {
	dict, err := NewHashDictionary(common.HexToAddress(ContractAddr), EthClient)
	if err != nil {
		Log.FatalF("Failed to instantiate a HashDictionary contract: %v", err)
	}
	hpair := HashDictionaryHashPair{
		Phash: common.Hex2Bytes(phash),
		Chash: common.Hex2Bytes(chash),
	}
	auth, err := bind.NewKeyedTransactorWithChainID(EthAccountSK, EthChainId)
	if err != nil {
		Log.FatalF("Failed to make transactor: %v", err)
	}
	tx, err := dict.VoteEntry(auth, hpair, vote)
	if err != nil {
		Log.FatalF("Failed to retrieve chash_alg: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
}

func find_chash(cmd *cobra.Command, args []string) {
	etherum_dial()
	var err error

	query_str := args[0]
	query_hex := common.Hex2Bytes(query_str)

	dict, err := NewHashDictionary(common.HexToAddress(ContractAddr), EthClient)
	if err != nil {
		Log.FatalF("Failed to instantiate a HashDictionary contract: %v", err)
	}
	phash_alg, err := dict.PhashAlg(nil)
	if err != nil {
		Log.FatalF("Failed to retrieve phash_alg: %v", err)
	}
	chash_alg, err := dict.ChashAlg(nil)
	if err != nil {
		Log.FatalF("Failed to retrieve chash_alg: %v", err)
	}
	fmt.Printf("====================== HashDictioanry ======================\n")
	fmt.Printf("Contract address: %s\n", ContractAddr)
	fmt.Printf("pHash algorithm:  %s\n", phash_alg)
	fmt.Printf("cHash algorithm:  %s\n", chash_alg)
	fmt.Printf("========================== Search ==========================\n")
	fmt.Printf("chash = %s\n", query_str)
	fmt.Printf("========================== Results =========================\n")
	entries, err := dict.GetEntriesByCHash(nil, query_hex)
	for _, entry := range entries {
		fmt.Printf("phash = %s\n", common.Bytes2Hex(entry.Hpair.Phash))
		fmt.Printf("votes = %s YES / %s NO\n", entry.VotesRight.String(), entry.VotesWrong.String())
	}
}
