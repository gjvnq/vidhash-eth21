package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gjvnq/go-logger"
	"github.com/spf13/cobra"
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
	Use:   "vote [phash/path] [chash]",
	Short: "Votes on an equivalence entry",
	Long:  "If only one argument is present, it will be interpreted as the file path.\nIf two are present, they will be interpreted as phash and chash.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var voteFileCmd = &cobra.Command{
	Use:   "vote-file [path]",
	Short: "Votes on an equivalence entry based on a local file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
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
		// Do Stuff Here
	},
}

var findCmd = &cobra.Command{
	Use:   "find [phash]",
	Short: "searches for a phash",
	Args:  cobra.ExactArgs(1),
	Run:   find_phash,
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
	rootCmd.AddCommand(makeWalletCmd)
	rootCmd.AddCommand(findCmd)
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
