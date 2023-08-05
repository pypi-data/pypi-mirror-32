from setuptools import setup

setup(
    name='udn-cis-client',
    version='1.3.1',
    packages=[
        'cis_client',
        'cis_client.commands',
        'cis_client.lib',
        'cis_client.lib.aaa',
        'cis_client.lib.cis_north',
        'cis_client.lib.cis_gateway',
    ],
    include_package_data=True,
    install_requires=[
        'click',
        'terminaltables',
        'requests',
        'progressbar2',
        'paramiko'
    ],
    entry_points='''
        [console_scripts]
        udn-cis-client=cis_client.cli:cli
    ''',
    zip_safe=False,
    script_name='setup.py',
    author='Marian Horban',
    author_email='mhorban@vidscale.com'
)
