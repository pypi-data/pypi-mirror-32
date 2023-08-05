===============================
Twilio Messages Scenario
===============================

Imports::

    >>> import os
    >>> import datetime
    >>> from decimal import Decimal
    >>> import twilio
    >>> from proteus import Model, Wizard
    >>> from trytond.tests.tools import activate_modules
    >>> from trytond.modules.company.tests.tools import create_company, \
    ...     get_company

Install twilio_messages::

    >>> config = activate_modules('twilio_messages')

Create company::

    >>> Company = Model.get('company.company')
    >>> _ = create_company()
    >>> company = get_company()
