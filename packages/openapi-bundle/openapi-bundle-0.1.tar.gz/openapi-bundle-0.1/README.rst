OpenAPI bundle
==============

Combine multiple json or yaml files into a single readable OpenAPI specification.

Why bother ?
------------

I am new to `OpenAPI (swagger) <http://swagger.io/>`_ . It's pain for me to write a very large yaml file, I follow `the post <https://azimi.me/2015/07/16/split-swagger-into-smaller-files.html>`_ to break the yaml into files.

"Most resolvers will resolve remote references first and the resolve local references." as it said in the post, but saddly json-refs gives me an error for local ref.

So I start a new tool and I believe local ref should be left untouched to improve readability.

Install
-------

::

  pip install openapi-bundle

Usage
-----

Basically::

   openapi-bundle /path/to/your/root/api/spec.yaml > /path/to/your/combined/api/spec.yaml

For other options, please run::

  openapi-bundle --help
