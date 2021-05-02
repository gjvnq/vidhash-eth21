import web3
import web3.eth
from hexbytes import HexBytes
from web3.middleware import geth_poa_middleware
import importlib.resources as pkg_resources

W3 = None
CONTRACT_ABI = pkg_resources.read_text('vidhasher', 'contract.abi')

def dial(ipc_addr):
	global W3
	if W3 is None:
		W3 = web3.Web3(web3.Web3.IPCProvider(ipc_addr))
		W3.middleware_onion.inject(geth_poa_middleware, layer=0)

def search_for_phash(phash, contract_addr, ipc_addr):
	dial(ipc_addr)
	phash = HexBytes(phash)
	contract = W3.eth.contract(address=contract_addr, abi=CONTRACT_ABI)

	chash_alg = contract.functions.chash_alg().call()
	phash_alg = contract.functions.phash_alg().call()

	print(f'address   = {contract_addr}')
	print(f'phash alg = {phash_alg}')
	print(f'chash alg = {chash_alg}')
	print()

	ans = contract.functions.getEntriesByPHash(phash).call()
	for entry in ans:
		(phash, chash), right, wrong = entry
		print(f'chash = 0x{chash.hex()}')
		print(f'votes = {right} YES / {wrong} NO')
		print()
