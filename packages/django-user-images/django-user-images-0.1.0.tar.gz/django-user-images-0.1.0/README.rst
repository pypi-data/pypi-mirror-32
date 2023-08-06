=============================
Django User Images
=============================

.. image:: https://badge.fury.io/py/django-user-images.svg
    :target: https://badge.fury.io/py/django-user-images

Django app for managing user files uploads, permissions and limits


Quickstart
----------

Install Django User Images::

    pip install django-user-images

Add it to your `INSTALLED_APPS`:

.. code-block:: python

    INSTALLED_APPS = (
        ...
        'django_user_images',
        ...
    )

If you want to use it with DRF, use django-user-images.rest:

.. code-block:: python

    INSTALLED_APPS = (
        ...
        'django_user_images.rest',
        ...
    )

Configuration
-------------

* USER_IMAGE_ONE_FILE_LIMIT - limit one file size upload in bytes
* USER_IMAGES_LIMIT - limit total file size per user, static in bytes
* USER_IMAGES_LIMITER - limit total file size per user, dynamic. Should be a string pointing to function, accepting user instance and returning limit in bytes


Credits
-------

Tools used in rendering this package:

*  Cookiecutter_
*  `cookiecutter-djangopackage`_

.. _Cookiecutter: https://github.com/audreyr/cookiecutter
.. _`cookiecutter-djangopackage`: https://github.com/pydanny/cookiecutter-djangopackage
