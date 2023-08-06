from setuptools import setup, find_packages

setup(
    name='woning-bricks',
    description='Development workflow made easy',
    version='0.0.4',
    packages=find_packages(),
    install_requires=[
        'click',
        'woning-wattle',
        'colorlog',
    ],
    entry_points={
        'console_scripts': [
            'bricks = bricks.__main__:cli'
        ]
    }
)
