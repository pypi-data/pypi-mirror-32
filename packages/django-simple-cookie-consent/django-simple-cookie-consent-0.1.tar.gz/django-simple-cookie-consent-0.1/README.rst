=====
Django Simple Cookie Consent
=====

Django Simple Cookie Consent is a simple Django app to display a custom cookie message.
Each visitor to the site will get a popup to say that the site uses cookies.

Detailed documentation is in the "docs" directory.

Quick start
-----------

1. Add "django_simple_cookie_consent.apps.DjangoSimpleCookieConsentConfig" to your INSTALLED_APPS setting like this::

    INSTALLED_APPS = [
        ...
        #Django >= 1.7
        'django_simple_cookie_consent.apps.DjangoSimpleCookieConsentConfig',
        #Django > 1.7
        'django_simple_cookie_consent',
    ]

2. If using Django South place the following in your SOUTH_MIGRATION_MODULES like this::

    SOUTH_MIGRATION_MODULES = {
        'django_simple_cookie_consent': 'django_simple_cookie_consent.south_migrations.0001_initial.py',
    }

3. Add "{% load django_simple_cookie_consent_tags %}" to your base template then place "{% display_cookie_consent %}" somewhere in the head tag like this::

    ...other loads
    {% load django_simple_cookie_consent_tags %}
    <head>
    ...other links/scripts
    {% display_cookie_consent %}
    </head>

4. Run `python manage.py migrate` to create the CookieConsentSettings model.

5. Start the development server and visit http://127.0.0.1:8000/admin/
   to create a settings instance (you'll need the Admin app enabled).

6. Visit http://127.0.0.1:8000/ to see the cookie message.


This package is using https://github.com/insites/cookieconsent but does not support all features at the current time (this may happen later on if the other features are needed).

