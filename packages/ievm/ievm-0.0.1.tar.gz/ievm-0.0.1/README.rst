ievm is a tool to download Microsoft IE virtual machine images.
Currently nothing more than download and checksum.
Create this since the old ievm.sh script is never timely updated.

Usage
-----

Basic usage::

    $ ievm [OPTIONS]

``OPTIONS `` includes the following required items:

- `-t TYPE`, `--type TYPE`:
  VM types, valid types are `hyperv`, `vagrant`, `virtualbox`, and `vmware`

- `-v VERSION`, `--version VERSION`:
  IE versions, valid version are `ie11-win7`, `msedge-win10`, `ie11-win81`,
  `ie9-win7`, and `ie10-win7`

- `-p PATH`, `--path PATH`:
  path to save downloaded file

- `-h`, `--help`:
  show help messages


Installation
------------

Install from PyPI::

    $ pip install ievm


License
-------

ievm is released under Apache License Version 2.0. See the LICENSE_ file for more details.

.. _LICENSE: https://github.com/zenixls2/ievm/blob/master/LICENSE
