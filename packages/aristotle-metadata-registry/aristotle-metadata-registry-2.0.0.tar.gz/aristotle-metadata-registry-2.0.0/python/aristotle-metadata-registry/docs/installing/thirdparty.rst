Configuring third-party apps
============================

Aristotle takes care of most of the work of getting a registry setup with the settings import::

    from aristotle_mdr.required_settings import *

but there are few areas for customisation or tweaking.

Django
------

Every django setting can be overridden, but the ones that will be most important when configuring Aristotle-MDR are:

* ``DATABASE`` - By default Aristotle will configure a SQLite file-based database.
  While this is fine for very small low-traffic registries, configuring Django to use a
  fully-fledged relational database management system like PostgreSQL or MySQL will
  be better for larger, high-traffic sites.
* ``ROOT_URLCONF`` - This is the python library that will be used to define the
  settings for django to resolve URLs. If you aren't using any extensions, you can
  just leave this as the default which points to the Aristotle URLs file - ``aristotle_mdr.urls``.
  If you are using extensions, you'll need to point this at the URLs file that you have created to
  handle all of the different URL configuration files for each extension.
* ``WSGI_APPLICATION`` - This points to the file and WSGI application that you have created
  to if you are intending to `deploy via a WSGI server <https://docs.djangoproject.com/en/1.8/howto/deployment/wsgi/>`_.

Haystack
--------

For search to work, Haystack is required to be installed. There are no options to disable this,
as without search a registry is quite useless. However you can change some settings.

* ``HAYSTACK_SEARCH_RESULTS_PER_PAGE`` - Self explanatory,  this defaults to 10 items per page.
* ``HAYSTACK_CONNECTIONS`` - This define which search indexers are being used and how they are
  connected. By default this uses the `Whoosh Engine <https://pypi.python.org/pypi/Whoosh/>`_,
  which is quite fast and because its a Pure-Python implementation reduces the complexity in getting it setup.
  `For more advanced usage, read the Haystack documentation <http://django-haystack.readthedocs.org/en/latest/tutorial.html#configuration>`_.
* ``HAYSTACK_SIGNAL_PROCESSOR`` - Included for completion, this defaults to ``aristotle_mdr.contrib.help.signals.AristotleHelpSignalProcessor``.
  This is a custom signal processor that performs real-time, status-aware changes to the index and monitors for changes to Help Pages.
  The alternative recommended option is ``aristotle_mdr.signals.AristotleSignalProcessor``, which  only monitors changes to metadata items.
  **Read the warnings below for why you probably only want to use these options.**

Warnings about Haystack
+++++++++++++++++++++++
* Always make sure ``haystack`` is included **once and only once** in ``INSTALLED_APPS``,
  otherwise your installation will throw errors.
* Make sure ``haystack`` is included in ``INSTALLED_APPS`` *before* ``aristotle_mdr``.
* Be aware that Haystack will only update search indexes when told, Aristotle includes a
  ``SignalProcessor`` that performs registation status-aware real-time updates to the index.
  **Switching this for another processor may expose private information** through search results,
  *but will not allow unauthorised users to access the complete item*.

LESS Compilation
----------------

Aristotle-MDR includes a number of uncompiled LESS files that need to be compiled by
django-static-precompiler. By default Aristotle-MDR uses the Python-based ``lesscpy``
compiler for this which is approximately compatible, but slower than, to the Node ``lessc`` compiler.
If you have complex requirements in your custom LESS files, want a faster compile time
or wish to use another CSS precompile type, override the following setting in your ``settings.py``::

    STATIC_PRECOMPILER_COMPILERS = (
      ('static_precompiler.compilers.LESS', {"executable": "lesscpy"}),
    )

In production, its advisable to compile the LESS files *once* and cache these withother static files.
This makes the choice of precompiler less of an issue for production environments.