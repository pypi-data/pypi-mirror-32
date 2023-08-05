# NotifyExceptions
Get exception on ADMINS mail IDs

About
-----
You can get all occurred exception to ADMIN mail.

Installation
------------
To install NotifyExceptions, simply use pip:

    $ pip install notifyexceptions

Documentation
-------------
You need to add few variables in your settings.py .

ADMINS = [('name', 'person@example.com')]
EMAIL_USE_TLS = True
EMAIL_HOST = '*******'
EMAIL_HOST_USER = '*******'
EMAIL_HOST_PASSWORD = '*******'
EMAIL_PORT = 587

Example
--------

    >>>from notifyexceptions import NotifyException
    >>>try:
            1/0
        except:
            NotifyException()
    
Found a Bug? Something Unsupported?
---------------

Please write me shuklashashank@outlook.com.

    
    
