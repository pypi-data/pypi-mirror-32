# LibSAN

Python library to manage Storage Area Network devices

## Dependencies:
    Make sure you have the following modules installed
        - paramiko

## How to install the module:
    python setup.py install 

## How to uninstall the module
    python setup.py install --force --record files.txt

    cat files.txt | xargs rm -rf

## How to create a tar file
    python setup.py sdist

## How to create an rpm package
    python setup.py bdist --format=rpm

## Usage:
    Before using the modules it is better to copy sample_san_top.conf
    to /etc/san_top.conf (this is the default path to read the config) and
    edit it according to your SAN environment.
