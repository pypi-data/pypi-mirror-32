==========================
CosmicDBSemantic Django App
==========================

Install
=======

- Run ``virtualenv demoenv --no-site-packages``
- Run ``demoenv\Scripts\activate``
- Run ``pip install Django``
- Run ``django-admin startproject demo``
- Run ``pip install cosmicdb``


Usage
=====

1. Add ``cosmicdb`` and requirements to your INSTALLED_APPS setting like this (your app must be first to override)::

        INSTALLED_APPS = (
            'YOURAPPHERE',
            'cosmicdb',
            'crispy_forms',
            'sitetree',
            'django_tables2',
            ... (rest of django apps)
        )

2. Add ``cosmicdb.urls`` to your urls.py like this (put cosmicdb urls last)::

        from django.conf.urls import url, include

        urlpatterns = [
            ...
            url(r'^', include('cosmicdb.urls')),
        ]

3. Add cosmicdb settings to your settings.py like this::

        LANGUAGE_CODE = 'en-au'
        COSMICDB_SITE_TITLE = 'Demo Site'
        CRISPY_TEMPLATE_PACK = 'semanticui'
        CRISPY_ALLOWED_TEMPLATE_PACKS = (CRISPY_TEMPLATE_PACK)
        DJANGO_TABLES2_TEMPLATE = 'django_tables2/semantic.html'
        COSMICDB_ALLOW_SIGNUP = False
        AUTH_USER_MODEL = 'cosmicdb.User'
        LOGIN_URL = '/login/'
        EMAIL_USE_TLS = True
        EMAIL_HOST = 'smtp.mailtrap.io'
        EMAIL_PORT = 465
        EMAIL_HOST_USER = '31c8dd7fd64bdd'
        EMAIL_HOST_PASSWORD = 'c11c8370e2408a'
        DEFAULT_FROM_EMAIL = EMAIL_HOST_USER
        DEFAULT_FROM_EMAIL_NAME = COSMICDB_SITE_TITLE
        SITETREE_MODEL_TREE = 'cosmicdb.CosmicDBTree'
        SITETREE_MODEL_TREE_ITEM = 'cosmicdb.CosmicDBTreeItem'

4. Run ``python manage.py migrate``

5. Run ``python manage.py collectstatic``

6. Run ``python manage.py createsuperuser``

Requirements
============

`Django>=2.0.5
<https://github.com/django/django/>`_


Optional
========
Custom Semantic UI Themes
=========================

1.  NodeJS (npm)

2. Gulp
``npm install -g gulp``
``npm install -g gulp-cli``

3. Add NODE_PATH env

4. Semantic UI
``cd PROJECT_DIR\cosmicdb\res\``
``npm install semantic-ui --save``

5. Put semanticui in semantic
``cd semantic/``
``gulp build``

6. Now you can copy your own theme from dist to cosmicdb\static\

``cd PROJECT ROOT``
``cp -rf cosmicdb/res/semanticui/semantic/dist cosmicdb/static/cosmicdb/semantic``

Site Tree
=========

1. Generate sitetree
``python manage.py sitetreedump > treedump.json``


Dev Notes
=========
``adjust cosmicdb/__init__.py for version number``

``python setup.py sdist bdist_wheel``

``replace the following line with version number``

``twine upload dist/cosmicdbsemantic-VERSION_NUMBER*``
``twine upload dist/cosmicdbsemantic-0.0.1* for 0.0.1``
