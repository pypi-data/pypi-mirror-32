Tempmail2
=========

Python API Wrapper version 2 for `temp-mail.org <https://temp-mail.org/>`_ service. Temp-mail.org is a service which lets you use anonymous emails for free. You can view full API specification in `api.temp-mail.org <http://api.temp-mail.org/>`_.

Requirements
------------

`requests <https://crate.io/packages/requests/>`_ - required.

Install ::

 $ pip install requests

Installation
------------

Installing with pip::

    $ pip install tempmail2

Usage
-----

Get all emails from given email login and domain::

    from tempmail2 import TempMail

    tm = TempMail(login='denis', domain='@gnail.pw')
    print tm.get_mailbox()  # list of emails in denis@gnail.pw

Generate email address and get emails from it::

    from tempmail2 import TempMail

    tm = TempMail()
    email = tm.get_email_address()  # v5gwnrnk7f@gnail.pw
    print tm.get_mailbox(email)  # list of emails

