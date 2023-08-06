rucio arc
=========

A rucio extension to the ARC data delivery service.

How to use
----------

    $ rucio-arc --help
    usage: rucio-arc [-h] [--debug] {ping,submit,query,cancel} ...

    positional arguments:
      {ping,submit,query,cancel}
        ping                Ping data delivery service.
        submit              Submit transfer.
        query               Query transfers.
        cancel              Cancel transfers.

    optional arguments:
      -h, --help            show this help message and exit
      --debug               print debug messages to stderr.

    $ rucio-arc ping
    Pinging https://localhost:8443/datadeliveryservice
    {'ResultCode': 'OK', 'LoadAvg': '0.47', 'AllowedDirs': ['/wlcg/cache']}

    $ rucio-transfertool-arc submit  --source gsiftp://hostname:2222/xxxx/zzzz  --destination /wlcg/cache/`uuidgen` --size 371272707 --checksum adler32:7a797445
    9afb332d-3aa3-4282-813e-c09b01d79243:

    $ rucio-arc query 9afb332d-3aa3-4282-813e-c09b01d79243
    9afb332d-3aa3-4282-813e-c09b01d79243:
	    ResultCode: TRANSFERRED

How to install
--------------

    pip install arc_rucio
