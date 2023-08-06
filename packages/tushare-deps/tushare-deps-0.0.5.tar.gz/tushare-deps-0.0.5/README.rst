============
tushare-deps
============


.. image:: https://img.shields.io/pypi/v/tushare-deps.svg
        :target: https://pypi.python.org/pypi/tushare-deps

.. image:: https://img.shields.io/travis/starofrainnight/tushare-deps.svg
        :target: https://travis-ci.org/starofrainnight/tushare-deps

.. image:: https://ci.appveyor.com/api/projects/status/github/starofrainnight/tushare-deps?svg=true
        :target: https://ci.appveyor.com/project/starofrainnight/tushare-deps

A library prepare the dependences for tushare library

* License: Apache-2.0
* Documentation: https://tushare-deps.readthedocs.io.

Preface
---------

tushare is an utility for crawling historical and Real-time Quotes data of
China stocks. It's great but not setting install dependences correctly. We will
found only tushare installed but without any dependences installed if we try
install tushare in a virtualenv environment by pip in python3.6 on Ubuntu 18.04.

That's so annoying to install those libraries manually.

Yes, fix this problem in tushare is more simple than just create a new library,
but there have some PRs (#452, #264) related not been pay attention for a long
time ...

Usage
---------

Just place 'tushare-deps' to your install_requires list in setup(), and remove
'tushare' depends, see sample below

.. code ::

    from setuptools import setup, find_packages

    install_requires = [
        'tushare-deps',
    ]

    setup(
        name='testlib',
        version='0.0.1',
        description="Just for test",
        long_description=long_description,
        author="Hong-She Liang",
        author_email='starofrainnight@gmail.com',
        url='https://github.com/starofrainnight/testlib',
        packages=find_packages(),
        include_package_data=True,
        install_requires=install_requires,
        license="Apache Software License",
        zip_safe=False,
        classifiers=[
            'Development Status :: 2 - Pre-Alpha',
            'Intended Audience :: Developers',
            'License :: OSI Approved :: Apache Software License',
            'Natural Language :: English',
            'Programming Language :: Python :: 3',
            'Programming Language :: Python :: 3.4',
            'Programming Language :: Python :: 3.5',
            'Programming Language :: Python :: 3.6',
        ],
    )

Credits
---------

This package was created with Cookiecutter_ and the `PyPackageTemplate`_ project template.

.. _Cookiecutter: https://github.com/audreyr/cookiecutter
.. _`PyPackageTemplate`: https://github.com/starofrainnight/rtpl-pypackage

