from setuptools import setup

setup(
    name='oversee',
    packages=['oversee'],
    version='0.22',
    package_data={'oversee': 'oversee/*.yaml'},
    include_package_data=True,
    description='A python utility to help manage your Ubuntu OS!',
    author='Jacob Smith',
    author_email='jacob.smith@unb.ca',
    url='http://github.com/jacsmith21/oversee',
    install_requires=[
        'click',
        'python-elevate',
        'PyYaml',
        'beautifulsoup4',
        'elevate'
    ],
    entry_points={
        'console_scripts': ['oversee=oversee.main:main']
    }
)
