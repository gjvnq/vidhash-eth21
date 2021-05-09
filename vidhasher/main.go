package main

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gjvnq/go-logger"
	"github.com/spf13/cobra"
	"io"
	"math/big"
	"os"
)

var Log *logger.Logger

var rootCmd = &cobra.Command{
	Use:   "vidhasher",
	Short: "Vidhasher is a prototype for a video hashing tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var voteCmd = &cobra.Command{
	Use:   "vote [phash] [chash] [vote]",
	Short: "Votes on an equivalence entry (vote = N/R/W)",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		etherum_dial()
		make_wallet()

		vote := uint8(0)
		if args[2] == "R" || args[2] == "r" {
			vote = 1
		} else if args[2] == "W" || args[2] == "w" {
			vote = 2
		} else if args[2] == "N" || args[2] == "n" {
			vote = 0
		} else {
			Log.FatalF("Invalid vote choice: %s", args[2])
		}
		vote_entry(args[0], args[1], vote)
	},
}

var voteFileCmd = &cobra.Command{
	Use:   "vote-file [path]",
	Short: "Votes on an equivalence entry based on a local file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		etherum_dial()
		make_wallet()

		hasher := sha1.New()
		f, err := os.Open(args[0])
		if err != nil {
			Log.Fatal(err)
		}
		if _, err := io.Copy(hasher, f); err != nil {
			Log.Fatal(err)
		}
		f.Close()
		chash := common.Bytes2Hex(hasher.Sum(nil))

		hashes := vphash_file(args[0])
		for _, hash := range hashes {
			vote_entry(hash, chash, 1)
		}
	},
}

var publishCmd = &cobra.Command{
	Use:   "publish [path]",
	Short: "Published and votes a video file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var hashCmd = &cobra.Command{
	Use:   "hash [path]",
	Short: "Hashes a video file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hashes := vphash_file(args[0])
		for _, hash := range hashes {
			fmt.Println(hash)
		}
	},
}

var findCmd = &cobra.Command{
	Use:   "find [phash]",
	Short: "searches for a phash",
	Args:  cobra.ExactArgs(1),
	Run:   find_phash,
}

var cfindCmd = &cobra.Command{
	Use:   "cfind [chash]",
	Short: "searches for a chash",
	Args:  cobra.ExactArgs(1),
	Run:   find_chash,
}

var deployCmd = &cobra.Command{
	Use:   "deploy [phash algorithm] [chash algorithm]",
	Short: "Deploys a new dictionary service smart contract instance",
	Args:  cobra.RangeArgs(1, 2),
	Run:   depoly_new_contract,
}

var makeWalletCmd = &cobra.Command{
	Use:   "make-wallet",
	Short: "Makes a new wallet",
	Args:  cobra.ExactArgs(0),
	Run:   make_new_wallet,
}

var (
	ContractAddr   string
	EthRpcAddr     string
	EthAccountStr  string
	EthAccountAddr common.Address
	EthClient      *ethclient.Client
	IpfsRpcCAddr   string
	EthWalletPath  string
	EthMnemonic    string
	EthChainId     *big.Int
	EthChainIdStr  string
	EthConn        int
	EthAccountSK   *ecdsa.PrivateKey
)

func init() {
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(hashCmd)
	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(voteCmd)
	rootCmd.AddCommand(voteFileCmd)
	rootCmd.AddCommand(makeWalletCmd)
	rootCmd.AddCommand(findCmd)
	rootCmd.AddCommand(cfindCmd)
	rootCmd.PersistentFlags().StringVar(&ContractAddr, "contract-addr", os.Getenv("VIDHASH_CONTRACT_ADDR"), "address to the smart contract that provides the dictionary service (fallback to VIDHASH_CONTRACT_ADDR)")
	rootCmd.PersistentFlags().StringVar(&IpfsRpcCAddr, "ipfs-rpc", "http://localhost:5001", "IPFS RPC listening address")
	rootCmd.PersistentFlags().StringVar(&EthRpcAddr, "eth-rpc", "/tmp/.geth-test.ipc", "geth RPC listening address")
	rootCmd.PersistentFlags().StringVar(&EthAccountStr, "eth-account", "", "the Etherum account that will pay the fees")
	rootCmd.PersistentFlags().StringVar(&EthMnemonic, "eth-mnemonic", os.Getenv("ETH_MNEMONIC"), "wallet mnemonic (fallback to env var ETH_MNEMONIC)")
	rootCmd.PersistentFlags().StringVar(&EthChainIdStr, "eth-chain-id", "1337", "chain id to use")
}

func main() {
	var err error

	Log, err = logger.New("main", 1, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		Log.Fatal(err)
	}
}
