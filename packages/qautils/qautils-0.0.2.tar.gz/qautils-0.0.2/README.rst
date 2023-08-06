QAutils ( *Quality Assurance development, utils module* )
=========================================================

.. image:: https://img.shields.io/github/issues/netzulo/qautils.svg
  :alt: Issues on Github
  :target: https://github.com/netzulo/qautils/issues

.. image:: https://img.shields.io/github/issues-pr/netzulo/qautils.svg
  :alt: Pull Request opened on Github
  :target: https://github.com/netzulo/qautils/issues

.. image:: https://img.shields.io/github/release/netzulo/qautils.svg
  :alt: Release version on Github
  :target: https://github.com/netzulo/qautils/releases/latest

.. image:: https://img.shields.io/github/release-date/netzulo/qautils.svg
  :alt: Release date on Github
  :target: https://github.com/netzulo/qautils/releases/latest

Code Metrics by sonarqube
-------------------------

.. image:: http://qalab.tk:82/api/badges/gate?key=qautils
  :alt: Quality Gate
  :target: http://qalab.tk:82/api/badges/gate?key=qautils
.. image:: http://qalab.tk:82/api/badges/measure?key=qautils&metric=lines
  :alt: Lines
  :target: http://qalab.tk:82/api/badges/gate?key=qautils
.. image:: http://qalab.tk:82/api/badges/measure?key=qautils&metric=bugs
  :alt: Bugs
  :target: http://qalab.tk:82/api/badges/gate?key=qautils
.. image:: http://qalab.tk:82/api/badges/measure?key=qautils&metric=vulnerabilities
  :alt: Vulnerabilities
  :target: http://qalab.tk:82/api/badges/gate?key=qautils
.. image:: http://qalab.tk:82/api/badges/measure?key=qautils&metric=code_smells
  :alt: Code Smells
  :target: http://qalab.tk:82/api/badges/gate?key=qautils
.. image:: http://qalab.tk:82/api/badges/measure?key=qautils&metric=sqale_debt_ratio
  :alt: Debt ratio
  :target: http://qalab.tk:82/api/badges/gate?key=qautils
.. image:: http://qalab.tk:82/api/badges/measure?key=qautils&metric=comment_lines_density
  :alt: Comments
  :target: http://qalab.tk:82/api/badges/gate?key=qautils


Python tested versions
----------------------

+-------------------+-------------------+-------------------+-------------------+-------------------+-------------------+
|  **3.6**          |  **3.5**          |  **3.4**          |  **3.3**          |  **3.2**          |  **2.7**          |
+===================+===================+===================+===================+===================+===================+
|    *Supported*    |    *Supported*    |    *Supported*    |  *Not Supported*  |  *Not Supported*  |    *Supported*    |
+-------------------+-------------------+-------------------+-------------------+-------------------+-------------------+


How to install ?
----------------

+ 1. Install from PIP file : ``pip install qautils``

+ 1. Install from setup.py file : ``python setup.py install``


How to exec tests ?
-------------------

+ 1. Tests from setup.py file : ``python setup.py test``


Getting Started
^^^^^^^^^^^^^^^

*Just starting example of usage before read* `Usage Guide`_.

.. code:: python


    from qautils.files import settings


    # file_path = './' by default
    SETTINGS = settings(
        file_path="/home/user/config/dir/",
        file_name="settings.json"
    )
    KEY_TO_CHECK = "some_json_key_name"


    try:
        print(SETTINGS[KEY_TO_CHECK])
    except Exception as err:
        print("ERROR: {}".format(err))
    finally:
        bot.close()



Contributing
~~~~~~~~~~~~

We welcome contributions to **qautils**! These are the many ways you can help:

* Submit patches and features
* Make **qautils** ( *new updates for community* )
* Improve the documentation for qautils_
* Report bugs 
* And donate_ !

Please read our **documentation** to get started. Also note that this project
is released with a code-of-conduct_ , please make sure to review and follow it.


.. _qautils: https://netzulo.github.io/qautils
.. _donate: https://opencollective.com/qautils
.. _code-of-conduct: https://github.com/netzulo/qalab/blob/master/CODE_OF_CONDUCT.rst
.. _Usage Guide: USAGE.rst