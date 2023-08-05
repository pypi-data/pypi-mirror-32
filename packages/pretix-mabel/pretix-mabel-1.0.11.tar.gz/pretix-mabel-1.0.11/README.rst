Pretix Mabel
==========================

This is a plugin for `pretix`_. 

Development setup
-----------------

1. Make sure that you have a working `pretix development setup`_.

2. Clone this repository, eg to ``local/pretix-mabel``.

3. Activate the virtual environment you use for pretix development.

4. Install mabel's dependencies (``pip install -r requirements.txt``)
 
5. Install our forked version of ``python-raven``: ``pip install git+https://github.com/mabelticketing/python-raven``

6. Create the database tables we need: ``python manage.py migrate``

7. Execute ``python setup.py develop`` within this directory to register this application with pretix's plugin registry.

8. Execute ``make`` within this directory to compile translations.

9. Restart your local pretix server. You can now use the plugin from this repository for your events by enabling it in
   the 'plugins' tab in the settings.


License
-------

Copyright 2017 Christopher Little

Released under the terms of the Apache License 2.0


.. _pretix: https://github.com/pretix/pretix
.. _pretix development setup: https://docs.pretix.eu/en/latest/development/setup.html
