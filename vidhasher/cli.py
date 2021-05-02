import click


@click.group()
def cli():
	"""Command line interface for the greet package"""
	pass


from .video import vp_hash_file
from .eth import search_for_phash

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

if __name__ == '__main__':
	cli(auto_envvar_prefix='VIDHASHER')