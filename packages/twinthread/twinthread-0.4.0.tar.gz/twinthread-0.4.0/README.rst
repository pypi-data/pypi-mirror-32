==========
twinthread
==========



Twinthread data access for data science exploration



Installation
------------

::

    git clone {this repo}
    cd project_directory
    # activate your venv i.e.
    pipenv shell
    pip install -e .[dev]

Usage
---------

Use Pipenv to install new modules so they're tracked correctly in the Pipfile::

    pipenv install new-package

To run the app locally and/or deploy to cloudfoundry::

    twinthread runserver
    twinthread dev deploy
