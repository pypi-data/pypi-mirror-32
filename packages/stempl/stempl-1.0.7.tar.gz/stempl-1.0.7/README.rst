stempl
======

|travis|

HTML templates with just Python.

Installing
----------

Use `pip`_ to install the package:

.. code-block:: text

    pip install stempl

Example
-------

For more examples checkout the `tests`_ file.

.. code-block:: python

    from flask import Flask
    from stempl import *


    app = Flask(__name__)


    def render_head():
        with Head() as t:
            t << meta(charset='utf-8')
            t << meta(name='viewport', content='width=device-width, initial-scale=1, shrink-to-fit=no')
            t << title('Hello World')
            t << link(href='https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css', rel='stylesheet')
            t << script(src='https://code.jquery.com/jquery-3.3.1.slim.min.js')
            t << script(src='https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js')
            t << script(src='https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js')

        return t()


    def render_navbar():
        with Ul(_class='navbar-nav') as u:
            u << li(a('Home', _class='nav-link', href='#'), _class='nav-item active')
            u << li(a('Features', _class='nav-link', href='#'), _class='nav-item active')
            u << li(a('Pricing', _class='nav-link', href='#'), _class='nav-item active')
            u << li(a('Disabled', _class='nav-link disabled', href='#'), _class='nav-item active')

        with Nav(_class='navbar navbar-expand-lg navbar-light bg-light') as n:
            n << a('Navbar', _class='navbar-brand', href='#')
            with Button(_class='navbar-toggler', type='button', data_toggle='collapse', data_target='#navbarNav') as b:
                b << span(_class='navbar-toggler-icon')
            n << b()
            n << div(u(), _class='collapse navbar-collapse', id='navbarNav')

        return n()


    def layout(content):
        with Html() as h:
            h << render_head()
            with Body() as b:
                b << render_navbar()
                b << div(content, _class='container', id='main')
            h << b()

        return doctype() + h()


    def index_template():
        with Div(id='index-page') as d:
            d << h1('Hello World')

        return layout(d)


    @app.route('/', methods=['GET'])
    def index():
        return index_template()


    if __name__ == '__main__':
        app.run(host='0.0.0.0', port=3000, debug=True)



.. _pip: https://pip.pypa.io/en/stable/quickstart/
.. _tests: stempl/tests.py
.. |travis| image:: https://travis-ci.org/gabrielhora/stempl.svg?branch=master
    :target: https://travis-ci.org/gabrielhora/stempl
    :alt: Travis CI status
