from setuptools import setup

setup(
    name='vidhasher',
    version=0.1,
    packages=['vidhasher'],
    install_requires=['click', 'opencv-python', 'scikit-image', 'numpy', 'bitstring', 'imagehash'],
    entry_points={'console_scripts': ['vidhasher = vidhasher.cli:cli']}
)