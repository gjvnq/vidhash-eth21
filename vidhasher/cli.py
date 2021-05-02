import click


@click.group()
def cli():
	"""Command line interface for the greet package"""
	pass


from .video import vp_hash_file

@click.command()
@click.argument('videopath')
def hash_video(videopath):
	hashes = vp_hash_file(videopath)
	for hash_ in hashes:
		print(hash_)
cli.add_command(hash_video)

if __name__ == '__main__':
	cli()