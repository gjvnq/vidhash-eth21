// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HashDictionaryEntry is an auto generated low-level Go binding around an user-defined struct.
type HashDictionaryEntry struct {
	Hpair      HashDictionaryHashPair
	VotesRight *big.Int
	VotesWrong *big.Int
}

// HashDictionaryHashPair is an auto generated low-level Go binding around an user-defined struct.
type HashDictionaryHashPair struct {
	Phash []byte
	Chash []byte
}

// HashDictionaryABI is the input ABI used to generate the binding from.
const HashDictionaryABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"alg_phash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"alg_chash\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"votes_right\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"votes_wrong\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structHashDictionary.Entry\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"AddedEntry\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"}],\"name\":\"addEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chash2idx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chash_alg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entries_list\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"votes_right\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"votes_wrong\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"getEntriesByCHash\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"votes_right\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"votes_wrong\",\"type\":\"int256\"}],\"internalType\":\"structHashDictionary.Entry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"getEntriesByPHash\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"votes_right\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"votes_wrong\",\"type\":\"int256\"}],\"internalType\":\"structHashDictionary.Entry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"}],\"name\":\"getEntryIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getVoteKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"hashpair2idx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"}],\"name\":\"hpair2bytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"phash2idx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phash_alg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"phash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"chash\",\"type\":\"bytes\"}],\"internalType\":\"structHashDictionary.HashPair\",\"name\":\"hpair\",\"type\":\"tuple\"},{\"internalType\":\"enumHashDictionary.VoteChoice\",\"name\":\"new_vote\",\"type\":\"uint8\"}],\"name\":\"voteEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HashDictionaryBin is the compiled bytecode used for deploying new contracts.
var HashDictionaryBin = "0x60806040523480156200001157600080fd5b50604051620025943803806200259483398181016040528101906200003791906200038d565b6000825114156200007f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000076906200044e565b60405180910390fd5b600081511415620000c7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000be9062000470565b60405180910390fd5b8160019080519060200190620000df929190620001da565b508060029080519060200190620000f8929190620001da565b50600060405180606001604052806040518060400160405280604051806020016040528060008152508152602001604051806020016040528060008152508152508152602001600081526020016000815250908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160008201518160000190805190602001906200019a9291906200026b565b506020820151816001019080519060200190620001b99291906200026b565b505050602082015181600201556040820151816003015550505050620006b1565b828054620001e89062000538565b90600052602060002090601f0160209004810192826200020c576000855562000258565b82601f106200022757805160ff191683800117855562000258565b8280016001018555821562000258579182015b82811115620002575782518255916020019190600101906200023a565b5b509050620002679190620002fc565b5090565b828054620002799062000538565b90600052602060002090601f0160209004810192826200029d5760008555620002e9565b82601f10620002b857805160ff1916838001178555620002e9565b82800160010185558215620002e9579182015b82811115620002e8578251825591602001919060010190620002cb565b5b509050620002f89190620002fc565b5090565b5b8082111562000317576000816000905550600101620002fd565b5090565b6000620003326200032c84620004bb565b62000492565b9050828152602081018484840111156200034b57600080fd5b6200035884828562000502565b509392505050565b600082601f8301126200037257600080fd5b8151620003848482602086016200031b565b91505092915050565b60008060408385031215620003a157600080fd5b600083015167ffffffffffffffff811115620003bc57600080fd5b620003ca8582860162000360565b925050602083015167ffffffffffffffff811115620003e857600080fd5b620003f68582860162000360565b9150509250929050565b60006200040f602183620004f1565b91506200041c8262000613565b604082019050919050565b600062000436602183620004f1565b9150620004438262000662565b604082019050919050565b60006020820190508181036000830152620004698162000400565b9050919050565b600060208201905081810360008301526200048b8162000427565b9050919050565b60006200049e620004b1565b9050620004ac82826200056e565b919050565b6000604051905090565b600067ffffffffffffffff821115620004d957620004d8620005d3565b5b620004e48262000602565b9050602081019050919050565b600082825260208201905092915050565b60005b838110156200052257808201518184015260208101905062000505565b8381111562000532576000848401525b50505050565b600060028204905060018216806200055157607f821691505b60208210811415620005685762000567620005a4565b5b50919050565b620005798262000602565b810181811067ffffffffffffffff821117156200059b576200059a620005d3565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f706861736820616c676f726974686d206d75737420626520737065636966696560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b7f636861736820616c676f726974686d206d75737420626520737065636966696560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b611ed380620006c16000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806389804bcb1161008c578063b5fe73ff11610066578063b5fe73ff14610238578063c0946d0614610268578063d8de8a7414610298578063fe09b504146102ca576100cf565b806389804bcb146101ba57806392bff539146101ea578063a6836c881461021a576100cf565b8063264329cf146100d45780633823ea5a1461010457806354aa85fa1461013457806358dca47114610150578063595e79e21461016c578063867274ff1461018a575b600080fd5b6100ee60048036038101906100e99190611418565b6102fa565b6040516100fb9190611a39565b60405180910390f35b61011e6004803603810190610119919061146c565b610341565b60405161012b9190611975565b60405180910390f35b61014e6004803603810190610149919061146c565b61038c565b005b61016a60048036038101906101659190611501565b610629565b005b610174610a6a565b6040516101819190611997565b60405180910390f35b6101a4600480360381019061019f91906113d7565b610af8565b6040516101b19190611a39565b60405180910390f35b6101d460048036038101906101cf9190611418565b610b26565b6040516101e19190611a39565b60405180910390f35b61020460048036038101906101ff919061146c565b610b6d565b6040516102119190611a39565b60405180910390f35b610222610b9d565b60405161022f9190611997565b60405180910390f35b610252600480360381019061024d9190611392565b610c2b565b60405161025f9190611953565b60405180910390f35b610282600480360381019061027d9190611392565b610c5d565b60405161028f9190611953565b60405180910390f35b6102b260048036038101906102ad9190611555565b610c8f565b6040516102c1939291906119fb565b60405180910390f35b6102e460048036038101906102df91906114ad565b610df7565b6040516102f19190611975565b60405180910390f35b600482805160208101820180518482526020830160208501208183528095505050505050818154811061032c57600080fd5b90600052602060002001600091509150505481565b60608180600001906103539190611a54565b8380602001906103639190611a54565b60405160200161037694939291906118d2565b6040516020818303038152906040529050919050565b600061039782610b6d565b9050600081146103dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d3906119b9565b60405180910390fd5b60008054905060036103ed84610341565b6040516103fa9190611909565b908152602001604051809103902081905550600482806000019061041e9190611a54565b60405161042c9291906118b9565b9081526020016040518091039020600080549050908060018154018082558091505060019003906000526020600020016000909190919091505560058280602001906104789190611a54565b6040516104869291906118b9565b9081526020016040518091039020600080549050908060018154018082558091505060019003906000526020600020016000909190919091505560006040518060600160405280846104d790611bd4565b815260200160018152602001600081525090506000819080600181540180825580915050600190039060005260206000209060040201600090919091909150600082015181600001600082015181600001908051906020019061053b929190611126565b506020820151816001019080519060200190610558929190611126565b50505060208201518160020155604082015181600301555050600061057d8433610df7565b905060016006826040516105919190611909565b908152602001604051809103902060006101000a81548160ff021916908360028111156105e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055507f8c0baa86d44d5f30f109c29dc75cea40e1ab88d58b5181ff13c7992f4626af1e8260405161061b91906119d9565b60405180910390a150505050565b600061063483610b6d565b9050600081141561064e576106488361038c565b50610a66565b600061065a8433610df7565b9050600060068260405161066e9190611909565b908152602001604051809103902060009054906101000a900460ff1690508360068360405161069d9190611909565b908152602001604051809103902060006101000a81548160ff021916908360028111156106f3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060016002811115610732577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81600281111561076b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156107d357600083815481106107ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906004020160020160008154809291906107cd90611c29565b91905055505b60028081111561080c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816002811115610845577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156108ad5760008381548110610885577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906004020160030160008154809291906108a790611c29565b91905055505b600160028111156108e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b846002811115610920577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156109885760008381548110610960577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060040201600201600081548092919061098290611cd5565b91905055505b6002808111156109c1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8460028111156109fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610a625760008381548110610a3a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600402016003016000815480929190610a5c90611cd5565b91905055505b5050505b5050565b60028054610a7790611c72565b80601f0160208091040260200160405190810160405280929190818152602001828054610aa390611c72565b8015610af05780601f10610ac557610100808354040283529160200191610af0565b820191906000526020600020905b815481529060010190602001808311610ad357829003601f168201915b505050505081565b6003818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b6005828051602081018201805184825260208301602085012081835280955050505050508181548110610b5857600080fd5b90600052602060002001600091509150505481565b60006003610b7a83610341565b604051610b879190611909565b9081526020016040518091039020549050919050565b60018054610baa90611c72565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd690611c72565b8015610c235780601f10610bf857610100808354040283529160200191610c23565b820191906000526020600020905b815481529060010190602001808311610c0657829003601f168201915b505050505081565b6060610c5560048484604051610c429291906118b9565b9081526020016040518091039020610e50565b905092915050565b6060610c8760058484604051610c749291906118b9565b9081526020016040518091039020610e50565b905092915050565b60008181548110610c9f57600080fd5b906000526020600020906004020160009150905080600001604051806040016040529081600082018054610cd290611c72565b80601f0160208091040260200160405190810160405280929190818152602001828054610cfe90611c72565b8015610d4b5780601f10610d2057610100808354040283529160200191610d4b565b820191906000526020600020905b815481529060010190602001808311610d2e57829003601f168201915b50505050508152602001600182018054610d6490611c72565b80601f0160208091040260200160405190810160405280929190818152602001828054610d9090611c72565b8015610ddd5780601f10610db257610100808354040283529160200191610ddd565b820191906000526020600020905b815481529060010190602001808311610dc057829003601f168201915b505050505081525050908060020154908060030154905083565b6060600082604051602001610e0c919061189e565b6040516020818303038152906040529050610e2684610341565b81604051602001610e38929190611920565b60405160208183030381529060405291505092915050565b60606000828054905067ffffffffffffffff811115610e98577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610ed157816020015b610ebe6111ac565b815260200190600190039081610eb65790505b50905060005b838054905081101561111c576000848281548110610f1e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020015481548110610f60577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906004020160405180606001604052908160008201604051806040016040529081600082018054610f9990611c72565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc590611c72565b80156110125780601f10610fe757610100808354040283529160200191611012565b820191906000526020600020905b815481529060010190602001808311610ff557829003601f168201915b5050505050815260200160018201805461102b90611c72565b80601f016020809104026020016040519081016040528092919081815260200182805461105790611c72565b80156110a45780601f10611079576101008083540402835291602001916110a4565b820191906000526020600020905b81548152906001019060200180831161108757829003601f168201915b5050505050815250508152602001600282015481526020016003820154815250508282815181106110fe577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250808061111490611d1e565b915050610ed7565b5080915050919050565b82805461113290611c72565b90600052602060002090601f016020900481019282611154576000855561119b565b82601f1061116d57805160ff191683800117855561119b565b8280016001018555821561119b579182015b8281111561119a57825182559160200191906001019061117f565b5b5090506111a891906111d3565b5090565b60405180606001604052806111bf6111f0565b815260200160008152602001600081525090565b5b808211156111ec5760008160009055506001016111d4565b5090565b604051806040016040528060608152602001606081525090565b600061121d61121884611ad0565b611aab565b90508281526020810184848401111561123557600080fd5b611240848285611be7565b509392505050565b60008135905061125781611e5f565b92915050565b60008083601f84011261126f57600080fd5b8235905067ffffffffffffffff81111561128857600080fd5b6020830191508360018202830111156112a057600080fd5b9250929050565b600082601f8301126112b857600080fd5b81356112c884826020860161120a565b91505092915050565b6000813590506112e081611e76565b92915050565b6000604082840312156112f857600080fd5b81905092915050565b60006040828403121561131357600080fd5b61131d6040611aab565b9050600082013567ffffffffffffffff81111561133957600080fd5b611345848285016112a7565b600083015250602082013567ffffffffffffffff81111561136557600080fd5b611371848285016112a7565b60208301525092915050565b60008135905061138c81611e86565b92915050565b600080602083850312156113a557600080fd5b600083013567ffffffffffffffff8111156113bf57600080fd5b6113cb8582860161125d565b92509250509250929050565b6000602082840312156113e957600080fd5b600082013567ffffffffffffffff81111561140357600080fd5b61140f848285016112a7565b91505092915050565b6000806040838503121561142b57600080fd5b600083013567ffffffffffffffff81111561144557600080fd5b611451858286016112a7565b92505060206114628582860161137d565b9150509250929050565b60006020828403121561147e57600080fd5b600082013567ffffffffffffffff81111561149857600080fd5b6114a4848285016112e6565b91505092915050565b600080604083850312156114c057600080fd5b600083013567ffffffffffffffff8111156114da57600080fd5b6114e6858286016112e6565b92505060206114f785828601611248565b9150509250929050565b6000806040838503121561151457600080fd5b600083013567ffffffffffffffff81111561152e57600080fd5b61153a858286016112e6565b925050602061154b858286016112d1565b9150509250929050565b60006020828403121561156757600080fd5b60006115758482850161137d565b91505092915050565b600061158a8383611767565b905092915050565b6115a361159e82611b8e565b611d67565b82525050565b60006115b482611b11565b6115be8185611b3f565b9350836020820285016115d085611b01565b8060005b8581101561160c57848403895281516115ed858261157e565b94506115f883611b32565b925060208a019950506001810190506115d4565b50829750879550505050505092915050565b600061162a8385611b72565b9350611637838584611be7565b82840190509392505050565b600061164e82611b1c565b6116588185611b50565b9350611668818560208601611bf6565b61167181611e18565b840191505092915050565b600061168782611b1c565b6116918185611b61565b93506116a1818560208601611bf6565b6116aa81611e18565b840191505092915050565b60006116c082611b1c565b6116ca8185611b72565b93506116da818560208601611bf6565b80840191505092915050565b6116ef81611ba0565b82525050565b6116fe81611ba0565b82525050565b600061170f82611b27565b6117198185611b7d565b9350611729818560208601611bf6565b61173281611e18565b840191505092915050565b600061174a601483611b7d565b915061175582611e36565b602082019050919050565b6000815250565b600060608301600083015184820360008601526117848282611807565b915050602083015161179960208601826116e6565b5060408301516117ac60408601826116e6565b508091505092915050565b600060608301600083015184820360008601526117d48282611807565b91505060208301516117e960208601826116e6565b5060408301516117fc60408601826116e6565b508091505092915050565b600060408301600083015184820360008601526118248282611643565b9150506020830151848203602086015261183e8282611643565b9150508091505092915050565b600060408301600083015184820360008601526118688282611643565b915050602083015184820360208601526118828282611643565b9150508091505092915050565b61189881611bca565b82525050565b60006118aa8284611592565b60148201915081905092915050565b60006118c682848661161e565b91508190509392505050565b60006118df82868861161e565b91506118ea82611760565b6001820191506118fb82848661161e565b915081905095945050505050565b600061191582846116b5565b915081905092915050565b600061192c82856116b5565b915061193782611760565b60018201915061194782846116b5565b91508190509392505050565b6000602082019050818103600083015261196d81846115a9565b905092915050565b6000602082019050818103600083015261198f818461167c565b905092915050565b600060208201905081810360008301526119b18184611704565b905092915050565b600060208201905081810360008301526119d28161173d565b9050919050565b600060208201905081810360008301526119f381846117b7565b905092915050565b60006060820190508181036000830152611a15818661184b565b9050611a2460208301856116f5565b611a3160408301846116f5565b949350505050565b6000602082019050611a4e600083018461188f565b92915050565b60008083356001602003843603038112611a6d57600080fd5b80840192508235915067ffffffffffffffff821115611a8b57600080fd5b602083019250600182023603831315611aa357600080fd5b509250929050565b6000611ab5611ac6565b9050611ac18282611ca4565b919050565b6000604051905090565b600067ffffffffffffffff821115611aeb57611aea611de9565b5b611af482611e18565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611b9982611baa565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611be03683611301565b9050919050565b82818337600083830152505050565b60005b83811015611c14578082015181840152602081019050611bf9565b83811115611c23576000848401525b50505050565b6000611c3482611ba0565b91507f8000000000000000000000000000000000000000000000000000000000000000821415611c6757611c66611d8b565b5b600182039050919050565b60006002820490506001821680611c8a57607f821691505b60208210811415611c9e57611c9d611dba565b5b50919050565b611cad82611e18565b810181811067ffffffffffffffff82111715611ccc57611ccb611de9565b5b80604052505050565b6000611ce082611ba0565b91507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611d1357611d12611d8b565b5b600182019050919050565b6000611d2982611bca565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611d5c57611d5b611d8b565b5b600182019050919050565b6000611d7282611d79565b9050919050565b6000611d8482611e29565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f656e74727920616c726561647920657869737473000000000000000000000000600082015250565b611e6881611b8e565b8114611e7357600080fd5b50565b60038110611e8357600080fd5b50565b611e8f81611bca565b8114611e9a57600080fd5b5056fea26469706673582212208e0e42a74d4c3554a25df6a476be4c33d53b8302f7a232101b17a13982357e0064736f6c63430008040033"

// DeployHashDictionary deploys a new Ethereum contract, binding an instance of HashDictionary to it.
func DeployHashDictionary(auth *bind.TransactOpts, backend bind.ContractBackend, alg_phash string, alg_chash string) (common.Address, *types.Transaction, *HashDictionary, error) {
	parsed, err := abi.JSON(strings.NewReader(HashDictionaryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashDictionaryBin), backend, alg_phash, alg_chash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashDictionary{HashDictionaryCaller: HashDictionaryCaller{contract: contract}, HashDictionaryTransactor: HashDictionaryTransactor{contract: contract}, HashDictionaryFilterer: HashDictionaryFilterer{contract: contract}}, nil
}

// HashDictionary is an auto generated Go binding around an Ethereum contract.
type HashDictionary struct {
	HashDictionaryCaller     // Read-only binding to the contract
	HashDictionaryTransactor // Write-only binding to the contract
	HashDictionaryFilterer   // Log filterer for contract events
}

// HashDictionaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashDictionaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashDictionaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashDictionaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashDictionaryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashDictionaryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashDictionarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashDictionarySession struct {
	Contract     *HashDictionary   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashDictionaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashDictionaryCallerSession struct {
	Contract *HashDictionaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HashDictionaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashDictionaryTransactorSession struct {
	Contract     *HashDictionaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HashDictionaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashDictionaryRaw struct {
	Contract *HashDictionary // Generic contract binding to access the raw methods on
}

// HashDictionaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashDictionaryCallerRaw struct {
	Contract *HashDictionaryCaller // Generic read-only contract binding to access the raw methods on
}

// HashDictionaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashDictionaryTransactorRaw struct {
	Contract *HashDictionaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashDictionary creates a new instance of HashDictionary, bound to a specific deployed contract.
func NewHashDictionary(address common.Address, backend bind.ContractBackend) (*HashDictionary, error) {
	contract, err := bindHashDictionary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashDictionary{HashDictionaryCaller: HashDictionaryCaller{contract: contract}, HashDictionaryTransactor: HashDictionaryTransactor{contract: contract}, HashDictionaryFilterer: HashDictionaryFilterer{contract: contract}}, nil
}

// NewHashDictionaryCaller creates a new read-only instance of HashDictionary, bound to a specific deployed contract.
func NewHashDictionaryCaller(address common.Address, caller bind.ContractCaller) (*HashDictionaryCaller, error) {
	contract, err := bindHashDictionary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashDictionaryCaller{contract: contract}, nil
}

// NewHashDictionaryTransactor creates a new write-only instance of HashDictionary, bound to a specific deployed contract.
func NewHashDictionaryTransactor(address common.Address, transactor bind.ContractTransactor) (*HashDictionaryTransactor, error) {
	contract, err := bindHashDictionary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashDictionaryTransactor{contract: contract}, nil
}

// NewHashDictionaryFilterer creates a new log filterer instance of HashDictionary, bound to a specific deployed contract.
func NewHashDictionaryFilterer(address common.Address, filterer bind.ContractFilterer) (*HashDictionaryFilterer, error) {
	contract, err := bindHashDictionary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashDictionaryFilterer{contract: contract}, nil
}

// bindHashDictionary binds a generic wrapper to an already deployed contract.
func bindHashDictionary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashDictionaryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashDictionary *HashDictionaryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashDictionary.Contract.HashDictionaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashDictionary *HashDictionaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashDictionary.Contract.HashDictionaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashDictionary *HashDictionaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashDictionary.Contract.HashDictionaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashDictionary *HashDictionaryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashDictionary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashDictionary *HashDictionaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashDictionary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashDictionary *HashDictionaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashDictionary.Contract.contract.Transact(opts, method, params...)
}

// Chash2idx is a free data retrieval call binding the contract method 0x89804bcb.
//
// Solidity: function chash2idx(bytes , uint256 ) view returns(uint256)
func (_HashDictionary *HashDictionaryCaller) Chash2idx(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "chash2idx", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Chash2idx is a free data retrieval call binding the contract method 0x89804bcb.
//
// Solidity: function chash2idx(bytes , uint256 ) view returns(uint256)
func (_HashDictionary *HashDictionarySession) Chash2idx(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _HashDictionary.Contract.Chash2idx(&_HashDictionary.CallOpts, arg0, arg1)
}

// Chash2idx is a free data retrieval call binding the contract method 0x89804bcb.
//
// Solidity: function chash2idx(bytes , uint256 ) view returns(uint256)
func (_HashDictionary *HashDictionaryCallerSession) Chash2idx(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _HashDictionary.Contract.Chash2idx(&_HashDictionary.CallOpts, arg0, arg1)
}

// ChashAlg is a free data retrieval call binding the contract method 0x595e79e2.
//
// Solidity: function chash_alg() view returns(string)
func (_HashDictionary *HashDictionaryCaller) ChashAlg(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "chash_alg")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChashAlg is a free data retrieval call binding the contract method 0x595e79e2.
//
// Solidity: function chash_alg() view returns(string)
func (_HashDictionary *HashDictionarySession) ChashAlg() (string, error) {
	return _HashDictionary.Contract.ChashAlg(&_HashDictionary.CallOpts)
}

// ChashAlg is a free data retrieval call binding the contract method 0x595e79e2.
//
// Solidity: function chash_alg() view returns(string)
func (_HashDictionary *HashDictionaryCallerSession) ChashAlg() (string, error) {
	return _HashDictionary.Contract.ChashAlg(&_HashDictionary.CallOpts)
}

// EntriesList is a free data retrieval call binding the contract method 0xd8de8a74.
//
// Solidity: function entries_list(uint256 ) view returns((bytes,bytes) hpair, int256 votes_right, int256 votes_wrong)
func (_HashDictionary *HashDictionaryCaller) EntriesList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hpair      HashDictionaryHashPair
	VotesRight *big.Int
	VotesWrong *big.Int
}, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "entries_list", arg0)

	outstruct := new(struct {
		Hpair      HashDictionaryHashPair
		VotesRight *big.Int
		VotesWrong *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hpair = *abi.ConvertType(out[0], new(HashDictionaryHashPair)).(*HashDictionaryHashPair)
	outstruct.VotesRight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VotesWrong = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EntriesList is a free data retrieval call binding the contract method 0xd8de8a74.
//
// Solidity: function entries_list(uint256 ) view returns((bytes,bytes) hpair, int256 votes_right, int256 votes_wrong)
func (_HashDictionary *HashDictionarySession) EntriesList(arg0 *big.Int) (struct {
	Hpair      HashDictionaryHashPair
	VotesRight *big.Int
	VotesWrong *big.Int
}, error) {
	return _HashDictionary.Contract.EntriesList(&_HashDictionary.CallOpts, arg0)
}

// EntriesList is a free data retrieval call binding the contract method 0xd8de8a74.
//
// Solidity: function entries_list(uint256 ) view returns((bytes,bytes) hpair, int256 votes_right, int256 votes_wrong)
func (_HashDictionary *HashDictionaryCallerSession) EntriesList(arg0 *big.Int) (struct {
	Hpair      HashDictionaryHashPair
	VotesRight *big.Int
	VotesWrong *big.Int
}, error) {
	return _HashDictionary.Contract.EntriesList(&_HashDictionary.CallOpts, arg0)
}

// GetEntriesByCHash is a free data retrieval call binding the contract method 0xc0946d06.
//
// Solidity: function getEntriesByCHash(bytes hash) view returns(((bytes,bytes),int256,int256)[])
func (_HashDictionary *HashDictionaryCaller) GetEntriesByCHash(opts *bind.CallOpts, hash []byte) ([]HashDictionaryEntry, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "getEntriesByCHash", hash)

	if err != nil {
		return *new([]HashDictionaryEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]HashDictionaryEntry)).(*[]HashDictionaryEntry)

	return out0, err

}

// GetEntriesByCHash is a free data retrieval call binding the contract method 0xc0946d06.
//
// Solidity: function getEntriesByCHash(bytes hash) view returns(((bytes,bytes),int256,int256)[])
func (_HashDictionary *HashDictionarySession) GetEntriesByCHash(hash []byte) ([]HashDictionaryEntry, error) {
	return _HashDictionary.Contract.GetEntriesByCHash(&_HashDictionary.CallOpts, hash)
}

// GetEntriesByCHash is a free data retrieval call binding the contract method 0xc0946d06.
//
// Solidity: function getEntriesByCHash(bytes hash) view returns(((bytes,bytes),int256,int256)[])
func (_HashDictionary *HashDictionaryCallerSession) GetEntriesByCHash(hash []byte) ([]HashDictionaryEntry, error) {
	return _HashDictionary.Contract.GetEntriesByCHash(&_HashDictionary.CallOpts, hash)
}

// GetEntriesByPHash is a free data retrieval call binding the contract method 0xb5fe73ff.
//
// Solidity: function getEntriesByPHash(bytes hash) view returns(((bytes,bytes),int256,int256)[])
func (_HashDictionary *HashDictionaryCaller) GetEntriesByPHash(opts *bind.CallOpts, hash []byte) ([]HashDictionaryEntry, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "getEntriesByPHash", hash)

	if err != nil {
		return *new([]HashDictionaryEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]HashDictionaryEntry)).(*[]HashDictionaryEntry)

	return out0, err

}

// GetEntriesByPHash is a free data retrieval call binding the contract method 0xb5fe73ff.
//
// Solidity: function getEntriesByPHash(bytes hash) view returns(((bytes,bytes),int256,int256)[])
func (_HashDictionary *HashDictionarySession) GetEntriesByPHash(hash []byte) ([]HashDictionaryEntry, error) {
	return _HashDictionary.Contract.GetEntriesByPHash(&_HashDictionary.CallOpts, hash)
}

// GetEntriesByPHash is a free data retrieval call binding the contract method 0xb5fe73ff.
//
// Solidity: function getEntriesByPHash(bytes hash) view returns(((bytes,bytes),int256,int256)[])
func (_HashDictionary *HashDictionaryCallerSession) GetEntriesByPHash(hash []byte) ([]HashDictionaryEntry, error) {
	return _HashDictionary.Contract.GetEntriesByPHash(&_HashDictionary.CallOpts, hash)
}

// GetEntryIdx is a free data retrieval call binding the contract method 0x92bff539.
//
// Solidity: function getEntryIdx((bytes,bytes) hpair) view returns(uint256)
func (_HashDictionary *HashDictionaryCaller) GetEntryIdx(opts *bind.CallOpts, hpair HashDictionaryHashPair) (*big.Int, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "getEntryIdx", hpair)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryIdx is a free data retrieval call binding the contract method 0x92bff539.
//
// Solidity: function getEntryIdx((bytes,bytes) hpair) view returns(uint256)
func (_HashDictionary *HashDictionarySession) GetEntryIdx(hpair HashDictionaryHashPair) (*big.Int, error) {
	return _HashDictionary.Contract.GetEntryIdx(&_HashDictionary.CallOpts, hpair)
}

// GetEntryIdx is a free data retrieval call binding the contract method 0x92bff539.
//
// Solidity: function getEntryIdx((bytes,bytes) hpair) view returns(uint256)
func (_HashDictionary *HashDictionaryCallerSession) GetEntryIdx(hpair HashDictionaryHashPair) (*big.Int, error) {
	return _HashDictionary.Contract.GetEntryIdx(&_HashDictionary.CallOpts, hpair)
}

// GetVoteKey is a free data retrieval call binding the contract method 0xfe09b504.
//
// Solidity: function getVoteKey((bytes,bytes) hpair, address voter) pure returns(bytes)
func (_HashDictionary *HashDictionaryCaller) GetVoteKey(opts *bind.CallOpts, hpair HashDictionaryHashPair, voter common.Address) ([]byte, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "getVoteKey", hpair, voter)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetVoteKey is a free data retrieval call binding the contract method 0xfe09b504.
//
// Solidity: function getVoteKey((bytes,bytes) hpair, address voter) pure returns(bytes)
func (_HashDictionary *HashDictionarySession) GetVoteKey(hpair HashDictionaryHashPair, voter common.Address) ([]byte, error) {
	return _HashDictionary.Contract.GetVoteKey(&_HashDictionary.CallOpts, hpair, voter)
}

// GetVoteKey is a free data retrieval call binding the contract method 0xfe09b504.
//
// Solidity: function getVoteKey((bytes,bytes) hpair, address voter) pure returns(bytes)
func (_HashDictionary *HashDictionaryCallerSession) GetVoteKey(hpair HashDictionaryHashPair, voter common.Address) ([]byte, error) {
	return _HashDictionary.Contract.GetVoteKey(&_HashDictionary.CallOpts, hpair, voter)
}

// Hashpair2idx is a free data retrieval call binding the contract method 0x867274ff.
//
// Solidity: function hashpair2idx(bytes ) view returns(uint256)
func (_HashDictionary *HashDictionaryCaller) Hashpair2idx(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "hashpair2idx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hashpair2idx is a free data retrieval call binding the contract method 0x867274ff.
//
// Solidity: function hashpair2idx(bytes ) view returns(uint256)
func (_HashDictionary *HashDictionarySession) Hashpair2idx(arg0 []byte) (*big.Int, error) {
	return _HashDictionary.Contract.Hashpair2idx(&_HashDictionary.CallOpts, arg0)
}

// Hashpair2idx is a free data retrieval call binding the contract method 0x867274ff.
//
// Solidity: function hashpair2idx(bytes ) view returns(uint256)
func (_HashDictionary *HashDictionaryCallerSession) Hashpair2idx(arg0 []byte) (*big.Int, error) {
	return _HashDictionary.Contract.Hashpair2idx(&_HashDictionary.CallOpts, arg0)
}

// Hpair2bytes is a free data retrieval call binding the contract method 0x3823ea5a.
//
// Solidity: function hpair2bytes((bytes,bytes) hpair) pure returns(bytes)
func (_HashDictionary *HashDictionaryCaller) Hpair2bytes(opts *bind.CallOpts, hpair HashDictionaryHashPair) ([]byte, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "hpair2bytes", hpair)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Hpair2bytes is a free data retrieval call binding the contract method 0x3823ea5a.
//
// Solidity: function hpair2bytes((bytes,bytes) hpair) pure returns(bytes)
func (_HashDictionary *HashDictionarySession) Hpair2bytes(hpair HashDictionaryHashPair) ([]byte, error) {
	return _HashDictionary.Contract.Hpair2bytes(&_HashDictionary.CallOpts, hpair)
}

// Hpair2bytes is a free data retrieval call binding the contract method 0x3823ea5a.
//
// Solidity: function hpair2bytes((bytes,bytes) hpair) pure returns(bytes)
func (_HashDictionary *HashDictionaryCallerSession) Hpair2bytes(hpair HashDictionaryHashPair) ([]byte, error) {
	return _HashDictionary.Contract.Hpair2bytes(&_HashDictionary.CallOpts, hpair)
}

// Phash2idx is a free data retrieval call binding the contract method 0x264329cf.
//
// Solidity: function phash2idx(bytes , uint256 ) view returns(uint256)
func (_HashDictionary *HashDictionaryCaller) Phash2idx(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "phash2idx", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Phash2idx is a free data retrieval call binding the contract method 0x264329cf.
//
// Solidity: function phash2idx(bytes , uint256 ) view returns(uint256)
func (_HashDictionary *HashDictionarySession) Phash2idx(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _HashDictionary.Contract.Phash2idx(&_HashDictionary.CallOpts, arg0, arg1)
}

// Phash2idx is a free data retrieval call binding the contract method 0x264329cf.
//
// Solidity: function phash2idx(bytes , uint256 ) view returns(uint256)
func (_HashDictionary *HashDictionaryCallerSession) Phash2idx(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _HashDictionary.Contract.Phash2idx(&_HashDictionary.CallOpts, arg0, arg1)
}

// PhashAlg is a free data retrieval call binding the contract method 0xa6836c88.
//
// Solidity: function phash_alg() view returns(string)
func (_HashDictionary *HashDictionaryCaller) PhashAlg(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HashDictionary.contract.Call(opts, &out, "phash_alg")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PhashAlg is a free data retrieval call binding the contract method 0xa6836c88.
//
// Solidity: function phash_alg() view returns(string)
func (_HashDictionary *HashDictionarySession) PhashAlg() (string, error) {
	return _HashDictionary.Contract.PhashAlg(&_HashDictionary.CallOpts)
}

// PhashAlg is a free data retrieval call binding the contract method 0xa6836c88.
//
// Solidity: function phash_alg() view returns(string)
func (_HashDictionary *HashDictionaryCallerSession) PhashAlg() (string, error) {
	return _HashDictionary.Contract.PhashAlg(&_HashDictionary.CallOpts)
}

// AddEntry is a paid mutator transaction binding the contract method 0x54aa85fa.
//
// Solidity: function addEntry((bytes,bytes) hpair) returns()
func (_HashDictionary *HashDictionaryTransactor) AddEntry(opts *bind.TransactOpts, hpair HashDictionaryHashPair) (*types.Transaction, error) {
	return _HashDictionary.contract.Transact(opts, "addEntry", hpair)
}

// AddEntry is a paid mutator transaction binding the contract method 0x54aa85fa.
//
// Solidity: function addEntry((bytes,bytes) hpair) returns()
func (_HashDictionary *HashDictionarySession) AddEntry(hpair HashDictionaryHashPair) (*types.Transaction, error) {
	return _HashDictionary.Contract.AddEntry(&_HashDictionary.TransactOpts, hpair)
}

// AddEntry is a paid mutator transaction binding the contract method 0x54aa85fa.
//
// Solidity: function addEntry((bytes,bytes) hpair) returns()
func (_HashDictionary *HashDictionaryTransactorSession) AddEntry(hpair HashDictionaryHashPair) (*types.Transaction, error) {
	return _HashDictionary.Contract.AddEntry(&_HashDictionary.TransactOpts, hpair)
}

// VoteEntry is a paid mutator transaction binding the contract method 0x58dca471.
//
// Solidity: function voteEntry((bytes,bytes) hpair, uint8 new_vote) returns()
func (_HashDictionary *HashDictionaryTransactor) VoteEntry(opts *bind.TransactOpts, hpair HashDictionaryHashPair, new_vote uint8) (*types.Transaction, error) {
	return _HashDictionary.contract.Transact(opts, "voteEntry", hpair, new_vote)
}

// VoteEntry is a paid mutator transaction binding the contract method 0x58dca471.
//
// Solidity: function voteEntry((bytes,bytes) hpair, uint8 new_vote) returns()
func (_HashDictionary *HashDictionarySession) VoteEntry(hpair HashDictionaryHashPair, new_vote uint8) (*types.Transaction, error) {
	return _HashDictionary.Contract.VoteEntry(&_HashDictionary.TransactOpts, hpair, new_vote)
}

// VoteEntry is a paid mutator transaction binding the contract method 0x58dca471.
//
// Solidity: function voteEntry((bytes,bytes) hpair, uint8 new_vote) returns()
func (_HashDictionary *HashDictionaryTransactorSession) VoteEntry(hpair HashDictionaryHashPair, new_vote uint8) (*types.Transaction, error) {
	return _HashDictionary.Contract.VoteEntry(&_HashDictionary.TransactOpts, hpair, new_vote)
}

// HashDictionaryAddedEntryIterator is returned from FilterAddedEntry and is used to iterate over the raw logs and unpacked data for AddedEntry events raised by the HashDictionary contract.
type HashDictionaryAddedEntryIterator struct {
	Event *HashDictionaryAddedEntry // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HashDictionaryAddedEntryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDictionaryAddedEntry)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HashDictionaryAddedEntry)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HashDictionaryAddedEntryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDictionaryAddedEntryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDictionaryAddedEntry represents a AddedEntry event raised by the HashDictionary contract.
type HashDictionaryAddedEntry struct {
	Arg0 HashDictionaryEntry
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedEntry is a free log retrieval operation binding the contract event 0x8c0baa86d44d5f30f109c29dc75cea40e1ab88d58b5181ff13c7992f4626af1e.
//
// Solidity: event AddedEntry(((bytes,bytes),int256,int256) arg0)
func (_HashDictionary *HashDictionaryFilterer) FilterAddedEntry(opts *bind.FilterOpts) (*HashDictionaryAddedEntryIterator, error) {

	logs, sub, err := _HashDictionary.contract.FilterLogs(opts, "AddedEntry")
	if err != nil {
		return nil, err
	}
	return &HashDictionaryAddedEntryIterator{contract: _HashDictionary.contract, event: "AddedEntry", logs: logs, sub: sub}, nil
}

// WatchAddedEntry is a free log subscription operation binding the contract event 0x8c0baa86d44d5f30f109c29dc75cea40e1ab88d58b5181ff13c7992f4626af1e.
//
// Solidity: event AddedEntry(((bytes,bytes),int256,int256) arg0)
func (_HashDictionary *HashDictionaryFilterer) WatchAddedEntry(opts *bind.WatchOpts, sink chan<- *HashDictionaryAddedEntry) (event.Subscription, error) {

	logs, sub, err := _HashDictionary.contract.WatchLogs(opts, "AddedEntry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDictionaryAddedEntry)
				if err := _HashDictionary.contract.UnpackLog(event, "AddedEntry", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddedEntry is a log parse operation binding the contract event 0x8c0baa86d44d5f30f109c29dc75cea40e1ab88d58b5181ff13c7992f4626af1e.
//
// Solidity: event AddedEntry(((bytes,bytes),int256,int256) arg0)
func (_HashDictionary *HashDictionaryFilterer) ParseAddedEntry(log types.Log) (*HashDictionaryAddedEntry, error) {
	event := new(HashDictionaryAddedEntry)
	if err := _HashDictionary.contract.UnpackLog(event, "AddedEntry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
