vmshepherd-zookeeper-driver
===========================

Introduction
------------

Provides plugin for ``VmShepherd`` - driver allows to store runtime data and lock management in Zookeeper.

Installation
------------

Simply use ``pip``.

:: 

    pip install vmshepherd-zookeeper-driver


Library requires (as well as VmShepherd itself) python 3.6 or later.

Usage
-----

Install package (in the same environment as VmShepherd) and configure ``VmShepherd`` like:

::

    # ...

    runtime:
      driver: ZookeeperDriver
      servers:
       - some.zk.host
      working_path: /vmshepherd
      addauth:
        auth: vmshepherduser:password


    # ...

Available config options
~~~~~~~~~~~~~~~~~~~~~~~~

.. csv-table::
   :header: "Name", "Type", "Description", "Default value"
   :widths: 15, 10, 40, 10

   "servers", "list", "Zookeeper hosts", ""
   "working_path", "string", "Base path where vmshepherd will read/write/create/deletes its nodes. A cdrwa permissions must be set for this path either to provided auth otherwise to anyone/world", "/vmshepherd"
   "addauth", "object", "Authentication options. If not provided or `null` no auth assumed.", "null"
   "addauth.scheme", "string", "Zookeeper's auth scheme (eg. digest sasl).", "digest"
   "addauth.auth", "string", "Auth data specific to given scheme (eg. user:password for digest)","vmshepherd:vmshepherd"

License
-------

MIT
