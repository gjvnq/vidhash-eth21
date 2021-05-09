import click

# TODO
# * Deploy contract
# * Test with Arbitrum
# * Vote
# * Publish on IPFS?

@click.group()
def cli():
	"""Command line interface for the greet package"""
	pass


from .video import vp_hash_file
from .eth import search_for_phash, deploy_contract

@click.command()
@click.argument('videopath')
def hash_video(videopath):
	hashes = vp_hash_file(videopath)
	for hash_ in hashes:
		print(hash_)
cli.add_command(hash_video)

@click.command()
@click.argument('phash')
@click.option('--ipc-addr', required=True, envvar='VIDHASHER_IPC_ADDR')
@click.option('--contract-addr', required=True, envvar='VIDHASHER_CONTRACT_ADDR')
def search(phash, contract_addr, ipc_addr):
	search_for_phash(phash, contract_addr, ipc_addr)
cli.add_command(search)

@click.command()
@click.argument('phash_alg')
@click.argument('chash_alg')
@click.option('--ipc-addr', required=True, envvar='VIDHASHER_IPC_ADDR')
@click.option('--account-key', required=True, envvar='VIDHASHER_ACCOUNT')
def deploy(phash_alg, chash_alg, account_key, ipc_addr):
	deploy_contract(phash_alg, chash_alg, account_key, ipc_addr)
cli.add_command(deploy)

if __name__ == '__main__':
	cli(auto_envvar_prefix='VIDHASHER')