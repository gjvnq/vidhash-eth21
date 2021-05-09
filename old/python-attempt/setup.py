from setuptools import setup

setup(
    name='vidhasher',
    version=0.1,
    packages=['vidhasher'],
    install_requires=['click', 'opencv-python', 'scikit-image', 'numpy', 'bitstring', 'imagehash', 'web3', 'hexbytes'],
    entry_points={'console_scripts': ['vidhasher = vidhasher.cli:cli']},
    package_data={'': ['LICENSE', 'vidhasher/contract.abi', 'vidhasher/contract.sol', 'vidhasher/contract.bin']},
    include_package_data=True,
)