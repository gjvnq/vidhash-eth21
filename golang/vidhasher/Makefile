.PHONY: all run

all: vidhasher

contract.bin: contract.sol
	solc --bin contract.sol | tail -n +4 > contract.bin
contract.abi: contract.sol
	solc --abi contract.sol | tail -n +4 > contract.abi
contract.go: contract.bin contract.abi
	abigen --abi contract.abi --pkg main --type HashDictionary --out contract.go --bin contract.bin

vidhasher: go.mod go.sum *.go contract.go
	go fmt
	go build

run: vidhasher
	./vidhasher

test-net:
	# WARING: does NOT work with Geth 1.10.2 (use 1.10.1) see https://github.com/ethereum/go-ethereum/issues/22792
	cd ../../ && geth --nodiscover --ipcpath /tmp/.geth-test.ipc --datadir test-net  --dev -http --http.corsdomain="http://remix.ethereum.org" --http.api web3,eth,debug,personal,net --vmdebug