.. _api:

API Reference
=============

.. currentmodule:: gordon

main
----

.. automodule:: gordon.main

.. autofunction:: gordon.main.setup

router
------

.. automodule:: gordon.router
.. autoclass:: gordon.router.GordonRouter


plugins_loader
--------------

.. automodule:: gordon.plugins_loader

.. autofunction:: gordon.plugins_loader.load_plugins


interfaces
----------

.. autointerface:: gordon.interfaces.IEventMessage
    :members:

.. autointerface:: gordon.interfaces.IEventConsumerClient
    :members:

.. autointerface:: gordon.interfaces.IEnricherClient
    :members:

.. autointerface:: gordon.interfaces.IPublisherClient
    :members:

.. autointerface:: gordon.interfaces.IGenericPlugin
    :members:

.. autointerface:: gordon.interfaces.IMetricRelay
    :members:
