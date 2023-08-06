DevIdentify Api Wrapper for Python
===============================

A simple python wrapper for dev-identify image grabber.

Usage
-----

-  Clone or download the repository.
-  Install the dependencies with ``pip install -r requirements.txt`` or ``pip install requests``
-  Install the package with ``pip install dev-identify``

Example
-------

.. code:: python

    from dev_identify import DevIdentify

    result = DevIdentify.email('example@gmail.com').save('/path/image_name.jpg').toJson()

    print(result)

Methods
-------

-  ``email()``

   Email method is a static and required method that takes in user email
   as parameter. Email method holds the actual response from devidentify
   server. This method returns an object that can be further chained
   with other methods available.

-  ``save()``

   Save method takes in full directory path as parameter along with
   filename and extension, i.e ``/path/image_name.jpg``. Save method
   will fail if the directory does not exist or has no write permission.
   This method returns an object that can be further chained with other
   methods available.

-  ``toJson()``

   ToJson method simply returns the response in json format. Chain this
   method with either ``email()`` or ``save()`` method to see the actual
   output.

-  ``toDict()``

   ToDict method simply returns the response as an dictionary. Chain this
   method with either ``email()`` or ``save()`` method to see the actual
   output.