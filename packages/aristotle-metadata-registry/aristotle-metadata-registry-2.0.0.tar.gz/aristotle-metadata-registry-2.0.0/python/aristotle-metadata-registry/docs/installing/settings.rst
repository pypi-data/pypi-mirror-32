Configuring the behavior of Aristotle-MDR
=========================================

Environment variables
---------------------

The default django settings file for Aristotle-MDR looks for a number of enviromnet variables for
storing files or configuring your webapp. These are all prefixed with ``aristotlemdr:``.

``BASE_DIR``
    Defaults to the path of where Aristotle is installed. Its highly adviced this is changed.
``SECRET_KEY``
    Defaults to a very insecure value - you MUST change this before going into production.
    `From Django settings documentation <https://docs.djangoproject.com/en/1.10/ref/settings/#std:setting-SECRET_KEY>`__:
    
        A secret key for a particular Django installation. This is used to provide cryptographic signing,
        and should be set to a unique, unpredictable value.
``STATIC_ROOT``
    Defaults to the value of ``BASE_DIR + "/static"``
    `From Django settings documentation <https://docs.djangoproject.com/en/1.10/ref/settings/#std:setting-STATIC_ROOT>`__:
    
        A secret key for a particular Django installation. This is used to provide cryptographic signing,
        and should be set to a unique, unpredictable value.
``MEDIA_ROOT``
    Defaults to the value of ``BASE_DIR + "/media"``
    `From Django settings documentation <https://docs.djangoproject.com/en/1.10/ref/settings/#std:setting-MEDIA_ROOT>`__:
    
        Absolute filesystem path to the directory that will hold user-uploaded files.


``ARISTOTLE_SETTINGS``
----------------------

The following are required within a dictionary in the settings for the configured Django project.

``CONTENT_EXTENSIONS``
    A list of the *namespaces* used to add additional content types,
    these are used when discovering the available extensions for about pages -
    required format a ``list`` of ``strings``.
``BULK_ACTIONS``
    A list of fully-qualified python paths to the bulk action form classes that 
    provide the action. :doc:`More information on configuring 
    bulk actions is available here <../extensions/bulk_actions>`.
``PDF_PAGE_SIZE``
    The default page size to deliver PDF downloads if a page size is not specified in the URL
``SEPARATORS``
    A key:value set that describes the separators to be used for name suggestions in the
    admin interface. These are set by specifying the key as the django model name for
    a given model, and the value as the separator.
    When a value for a model isn't stated in this field it defaults to a hyphen ``-``.
    The default settings in ``required_settings.py`` set additional defaults and
    specify the separator for "DataElements" as a comma with a single space ``, ``
    and the separator for "DataElementConcepts" as an em-dash ``–``.
``SITE_NAME``
    The main title for the site - required format ``string`` or ``unicode``.
``SITE_BRAND``
    A URL to the logo to use for the site, this can be relative or absolute.
``SITE_INTRO``
    The introductory text use on the home page as a prompt for users -
    required format ``string`` or ``unicode``.
``WORKGROUP_CHANGES``
    An array that specified which classes of user can move items between workgroups.
    Possible options include ``'admin'``, ``'manager'`` or ``'submitter'``.
``DOWNLOADERS``
    A list of download options - explained below:

``ARISTOTLE_SETTINGS.DOWNLOADERS``
**********************************
This is a **list of tuples** that define the different download options that will
be made available to users. This tuple defines in order:

* filetype used in the URL to catch this download type - must match the regex ``[a-zA-Z0-9\-\.]+``.
* the name presented on the download menu
* The `Font-Awesome <http://fortawesome.github.io/Font-Awesome/icons/#file-type>`_ icon used in the download menu
* the python module that includes the ``downloader.py`` file for handling this filetype

For example::

    ('pdf',"PDF","fa-file-pdf-o","aristotle_mdr")

Menu options are only given if a template for that download file type exists for
a given object. The first (filetype) setting is used when catching URLs for downloads, so that
when resolving URLs the filetype is used in the URL in the following way::

    /download/<download-file-type>/<item-id>

This file type is also passed to the download manager for this filetype, so that multiple
file types can be handled by the same extension.

For example, if an object class had a PDF template, based on the above
configuration the menu below would be accessible:

 .. image:: /_static/download_menu.png
    :alt: An example download menu with one option for a PDF download link.

And clicking this would access the following relative URL::

    /download/pdf/<object_class_id>

For more information on creating additional download extensions consult the guide on
:doc:`Adding new download formats </extensions/downloads>`.

Sample settings
---------------

Below is the ``ARISTOTLE_SETTINGS`` used on the hosted
Aristotle example::

    ARISTOTLE_SETTINGS = {
       # 'The main title for the site.'
        'SITE_NAME': 'Aristotle Metadata Registry',
       # URL for the Site-wide logo
        'SITE_BRAND': '/static/aristotle_mdr/images/aristotle_small.png',
       # 'Intro text use on the home page as a prompt for users.'
        'SITE_INTRO': 'Use Aristotle Metadata to search for metadata...',
       # Extensions that add additional object types for search/display.
        'CONTENT_EXTENSIONS' : [ 'comet' ],
       # Separators for auto-generating the names of constructed items.
        'SEPARATORS': { 'DataElement':',',
                    'DataElementConcept':'–'},
        'DOWNLOADERS': [
            ('pdf','PDF','fa-file-pdf-o','aristotle_pdf'),
        ]
      }
