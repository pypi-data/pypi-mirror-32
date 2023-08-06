===============================
Welcome to python-quoine v0.1.4
===============================

.. image:: https://img.shields.io/pypi/v/python-quoine.svg
    :target: https://pypi.python.org/pypi/python-quoine

.. image:: https://img.shields.io/pypi/l/python-quoine.svg
    :target: https://pypi.python.org/pypi/python-quoine

.. image:: https://img.shields.io/travis/sammchardy/python-quoine.svg
    :target: https://travis-ci.org/sammchardy/python-quoine

.. image:: https://img.shields.io/coveralls/sammchardy/python-quoine.svg
    :target: https://coveralls.io/github/sammchardy/python-quoine

.. image:: https://img.shields.io/pypi/wheel/python-quoine.svg
    :target: https://pypi.python.org/pypi/python-quoine

.. image:: https://img.shields.io/pypi/pyversions/python-quoine.svg
    :target: https://pypi.python.org/pypi/python-quoine

This is an unofficial Python wrapper for the `Quoinex and Qryptos exchanges REST API v2 <https://developers.quoine.com/v2>`_. I am in no way affiliated with Quoine, Quoinex or Qryptos, use at your own risk.

PyPi
  https://pypi.python.org/pypi/python-quoine

Source code
  https://github.com/sammchardy/python-quoine

Documentation
  https://python-quoine.readthedocs.io/en/latest/


Features
--------

- Implementation of all REST endpoints.
- Simple handling of authentication
- Response exception handling
- Simple market and limit buy functions
- Margin orders for Quoinex

TODO
----

- Websocket implementation

Quick Start
-----------

Register an account with `Quoinex <https://accounts.quoinex.com/sign-up?affiliate=PAxghztC67615>`_
or `Qryptos <https://accounts.qryptos.com/sign-up?affiliate=PAxghztC67615>`_.

Generate an API Key and assign relevant permissions.

.. code:: bash

    pip install python-quoine


.. code:: python

    from quoine.client import Quoinex
    client = Quoinex(api_key, api_secret)

    # get products
    products = client.get_products()

    # get market depth
    depth = client.get_order_book(product_id=products[0]['id'])

    # place market buy order
    order = client.create_market_buy(
        product_id=products[0]['id'],
        quantity='100',
        price_range='0.01')

    # get list of filled orders
    filled_orders = client.get_orders(status=client.STATUS_FILLED)


For more `check out the documentation <https://python-quoine.readthedocs.io/en/latest/>`_.

Donate
------

If this library helped you out feel free to donate.

- ETH: 0xD7a7fDdCfA687073d7cC93E9E51829a727f9fE70
- NEO: AVJB4ZgN7VgSUtArCt94y7ZYT6d5NDfpBo
- LTC: LPC5vw9ajR1YndE1hYVeo3kJ9LdHjcRCUZ
- BTC: 1Dknp6L6oRZrHDECRedihPzx2sSfmvEBys

Other Exchanges
---------------

If you use `Binance <https://www.binance.com/?ref=10099792>`_ check out my `python-binance <https://github.com/sammchardy/python-binance>`_ library.

If you use `Kucoin <https://www.kucoin.com/#/?r=E42cWB>`_ check out my `python-kucoin <https://github.com/sammchardy/python-kucoin>`_ library.

If you use `Allcoin <https://www.allcoin.com/Account/RegisterByPhoneNumber/?InviteCode=MTQ2OTk4MDgwMDEzNDczMQ==>`_ check out my `python-allucoin <https://github.com/sammchardy/python-allcoin>`_ library.

If you use `IDEX <https://idex.market>`_ check out my `python-idex <https://github.com/sammchardy/python-idex>`_ library.

If you use `BigONE <https://big.one>`_ check out my `python-bigone <https://github.com/sammchardy/python-bigone>`_ library.

.. image:: https://analytics-pixel.appspot.com/UA-111417213-1/github/python-quoine?pixel