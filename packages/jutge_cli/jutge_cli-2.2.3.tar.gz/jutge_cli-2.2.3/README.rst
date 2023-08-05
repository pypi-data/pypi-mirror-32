jutge\_cli: a command line interface for `jutge.org`_
=========================================================================

#. `Intro`_
#. `Installation`_

    #. `Global installation (root)`_
    #. `Archlinux PKGBUILD`_
    #. `Installation using virtualenv (no root)`_

#. `Usage`_
#. `Configuration`_

    #. `Basic options`_
    #. `Problem sets`_
    #. `Login credentials`_

#. `Authenticate to jutge`_

    #. `login command`_
    #. `cookie command`_
    #. `cookie flag`_

#. `Commands`_

    #. `add-test (add)`_
    #. `archive`_
    #. `check`_
    #. `cookie`_
    #. `download (down)`_
    #. `login`_
    #. `new`_
    #. `show`_
    #. `test`_
    #. `import`_
    #. `upload (up)`_

#. `License`_

Intro
-----

``jutge_cli`` is a python3 console application that aims to automate common
tasks when working with `jutge.org`_ problems. Those tasks
include:

* Creating new files named after problem title given the problem code
* Displaying statement and public test cases of a given problem
* Compiling and testing a program against public test cases
* Uploading program solutions to `jutge.org`_
* Checking `jutge.org`_ results for last submissions or for
  a specific problem.
* Adding ant testing against custom test cases to a problem
* Batch uploading problems from a given problem set
* Batch creating new files of a given problem set
* Extract and rename problem solutions from a `jutge.org`_ zip file export
  to a specific folder.


Installation
------------

``jutge_cli`` is included in `pypi <https://pypi.org/project/jutge_cli/>`_
and as such, it can be installed through ``pip3``.


Global installation (root)
~~~~~~~~~~~~~~~~~~~~~~~~~~

To install python module run:

.. code:: sh

    sudo pip3 install jutge_cli

This should install all dependencies and create an executable named
``jutge`` in ``/usr/bin/jutge``.


Archlinux PKGBUILD
~~~~~~~~~~~~~~~~~~

There is also a ``PKGBUILD`` included in the repository for arch linux users.


Installation using virtualenv (no root)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

You can install this program inside of a python3
`virtualenv <https://virtualenv.pypa.io/>`_:

.. code:: sh

    virtualenv -p /usr/bin/python3 jutge_cli_venv && cd jutge_cli_venv
    source bin/activate
    pip3 install jutge_cli

Once the above commands complete successfully, the ``jutge`` will be installed
inside the ``bin`` folder of the virtualenv. It is recommended to link it to
the user ``bin`` folder and add it to your ``$PATH``.

.. code:: sh

    mkdir ~/bin
    ln -s bin/jutge ~/bin/jutge

Remember to add bin to your path by adding the following line to ``.bashrc``
or equivalent:

.. code:: sh

    export PATH=$PATH:~/bin


Usage
-----

To use ``jutge_cli`` run the command ``jutge`` followed by the
subcommand you want to execute:

.. code:: sh

    jutge [SUBCOMMAND] [SUBCOMMAND_OPTIONS]

For the program to work you will have to either specify the code of the
problem you want to test (``-c`` flag) or rename the program file so
that it begins with the code. The code must match the following regular
expression: ``[PGQX]\d{5}_(ca|en|es)`` (note that the code includes the
language although it is not mandatory).


Configuration
-------------

You can configure default parameters through the YAML configuration file:
``~/.jutge_cli.yaml``.

Basic options
~~~~~~~~~~~~~

The following example lists all the basic options and
their default values:

.. code:: yaml

    database : ~/Documents/jutge/DB
    regex : '[PGQX]\d{5}_(ca|en|es)'
    diff-prog : diff
    diff-flags : -y
    inp-suffix : inp
    cor-suffix : cor
    folder : ~/Documents/jutge/Done

Those options can also be specified using the flags with the same name.

Although you can change the regex it is **not** recommended to do so since bad
regex may break correct functionality.


Problem sets
~~~~~~~~~~~~

You can also add problem sets with the ``problem_sets`` group. These will
make the commands ``new``, ``archive`` and ``update`` classify problems into
folders:

.. code:: yaml

    problem_sets:
        P1: [P19724, P34279, P37297, P37469, P42042, P51126, P51352, P61634, P66529, P67171, P70955, P82374, P89265, P92351, P98960, P99182, X54725, X59678, X64734, X89070]
        P2: [P27341, P28754, P29448, P32046, P34451, P35547, P37500, P55622, P59539, P59875, P60816, P64976, P65171, P74398, P79784, P85370, P97156, X30229, X32391, X80452]
        P3: [P13623, P19991, P29973, P32533, P61061, P79817, P80660, P87323, P96767, X01646, X08783, X26853, X29759, X59091, X84338, X98097]

The above configuration file will save problems ``P19724, P34279...`` into
folder ``P1``, problems ``P27341, P28754...`` into ``P2`` and so on.


Login credentials
~~~~~~~~~~~~~~~~~

You can also provide login credentials in the configuration file inside
the group ``login``:

.. code:: yaml

    login:
        email: myemail@mydomain.com
        password: mypassword

You can omit either email, password or both and the login command will
prompt the user for input when issued.


Authenticate to jutge
---------------------

To upload problem solutions or to access private problems (the ones which code
starts with ``X``) you must be logged in into `jutge.org`_.
The preferred method to login is through the ``jutge login`` command although
there are 2 more methods involving cookies.


login command
~~~~~~~~~~~~~

Issuing the command ``jutge login`` will prompt the user for their email and
password and save the session cookie for next use. If email or
password are already provided in `Login credentials`_ it will not prompt the
user to input them. For more details on the command see `login`_


cookie command
~~~~~~~~~~~~~~

The command ``jutge cookie`` accepts a cookie as a parameter and will
store it for next use. For more details on the command see `cookie`_


cookie flag
~~~~~~~~~~~

You can also explicitly provide a cookie to each subcommand call through the
``--cookie`` flag:

.. code:: sh

    jutge --cookie MY_COOKIE download -c X00000


Global flags
------------

Most of the flags depend on the subcommands, but there are some global
flags that effect all subcommands. Those are:

-  ``--regex MY_REGEX`` regular expression used to extract codes from filenames
-  ``--cookie MY_COOKIE`` Cookie used to connect to `jutge.org`_
-  ``--database FOLDER`` Change database location
-  ``--no-download`` If this flag is provided, ``jutge_cli`` will not attempt
   to connect to the internet


Commands
--------

#. `add-test (add)`_
#. `archive`_
#. `check`_
#. `cookie`_
#. `download (down)`_
#. `login`_
#. `new`_
#. `show`_
#. `test`_
#. `import`_
#. `upload (up)`_

add-test (add)
~~~~~~~~~~~~~~

This command adds a custom test case into the database. A test case consists
of two files, the input that will be feed to the program and the expected
output or solution. Those files can be provided through the flags ``-i``
(input) and ``-o`` (expected output) or if omitted the user will be prompted to
input them manually through stdin.

The following example will add the contents of files ``inp`` and ``cor`` to
the database as test cases for the problem ``P00001_ca``

.. code:: sh

    # Add the contents of inp and cor to the database for problem P00001_ca:
    jutge add-test -i inp -o cor P00001_ca_prog.cpp

    # Prompt the user to enter the input and expected output and add them to
    # the database for problem P00001_ca:
    jutge add-test P00001_ca_prog.cpp


archive
~~~~~~~

Move problem file to the archive folder. This folder can be
changed through the ``-f`` flag. To overwrite files already in the folder
use the ``--overwrite`` flag.

The default behaviour is to move the file to the folder, if you want to copy
it instead use the ``--copy`` flag.

The following example will move the file ``P00001_ca_prog.cpp`` to the
folder ``Accepted`` and overwrite if already in the folder.

.. code:: sh

    jutge archive --folder Accepted/ P00001_ca_prog.cpp --overwrite


check
~~~~~

Checks submissions to `jutge.org`_
and displays them in the terminal. The program will return 0 if the last
submission's verdict is ``AC`` or ``100/100``  and 1 otherwise.

This subcommand accepts 3 flags:

* ``--last`` show only the last submission
* ``--reverse`` order the output so that the last submission is on top
* ``--code`` check if a given problem code is accepted, rejected or not done
  yet


cookie
~~~~~~

Add cookie provided as first argument to a temporary directory so it is used
for next commands. If the first argument is ``delete`` the current cookie
will be deleted and if the argument is ``print`` or ``show`` it will
output the current saved cookie.

The command will check that the provided cookie is valid before saving the
value, to skip this check use the flag ``--skip-check``.


download (down)
~~~~~~~~~~~~~~~

This command will attempt to download the html page and zip file corresponding
to the given problem from `jutge.org`_ and add them to the
database. Either a code flag (``-c``) or a program file (``-p``) must be
provided.

Note that other commands that depend on the database files will
automatically try to download them if they don't exist and therefore
this command is only useful when populating the database in advance.

The following example will populate the local database for problem
``P00001_en``:

.. code:: sh

    jutge download P00001_en


login
~~~~~

Prompt the user to input their credentials and login to `jutge.org`_. If
credentials are already specified in the configuration file (`Login
credentials`_) it will not prompt for them.

The flags ``--email`` and ``--password`` can be used to specify the credentials
without prompting and to override the ones specified in the configuration file.


new
~~~

This command must be followed by a code. It will fetch the problem title
from the code and create a new file whose name is the code followed by
the title. The ``--extension`` or ``-e`` flag can be used to specify the
extension of the file (defaults to ``cpp``).

If flag ``--problem-set`` is provided, all programs in the specified problem
set will be created inside a folder named after the problem set.

The following example will populate create a new python file named
``P87523_ca_-_Hola-adéu.py``

.. code:: sh

    jutge new P87523_ca --extension py


show
~~~~

This command provides 3 sub commands to print information about the problem.
Those are:

-  ``title`` print problem title
-  ``stat`` print statement
-  ``cases`` print test cases in database

By default ``stat`` will parse the problem statement through ``pypandoc`` to
optimize the output for terminal if you prefer raw HTML or ``pypandoc`` takes
to much time to parse the output you can use the flag ``--html``.

The following example will print all cases in the database for the problem
``P87523_ca`` (if any).

.. code:: sh

    jutge show cases P87523_ca


test
~~~~

This is the most useful command in the tool set. It allows to test your
code against all the test cases found in the database and output side by
side differences using ``diff``.

The command takes a file that can be either an executable or source file or
script of a supported language executable file as parameter and tests it
against the test cases in the database folder. Note that if the program if a
source file that needs to be compiled, ``jutge_cli`` will compile it to
a file named after the original name with extension ``.x``.

You can specify an other program to act as ``diff`` (such as ``colordiff``) and
its flags (separated by commas) through ``--diff-prog`` and ``--diff-flags``.

The following example will test the executable ``P87523_ca_prog.x`` against
the test cases for problem P87523\_ca. The expected output and the output of
the program will be shown side by side using ``colordiff``.

.. code:: sh

    jutge test P87523_ca_prog.x --diff-prog colordiff


import
~~~~~~

This command extracts all accepted submissions from a `jutge.org`_ zip file,
renames them according to their title and adds them to the archive folder
that can be specified through the ``-f`` flag or in the main configuration
file. Note that the zip file must be the one downloaded from your
`jutge.org`_ profile.

.. code:: sh

    jutge import problems.zip


upload (up)
~~~~~~~~~~~

This command uploads a file to `jutge.org`_ to be
evaluated. Note that you must have a valid cookie previously saved by ``jutge
cookie PHPSSID`` or you can provide it through the ``--cookie`` flag. As of
now, the program cannot report if the upload was successful so you will have to
check your submissions page manually. The compiler to use will be determined by
the filename extension but you can specify another one through the
``--compiler`` flag.

.. code:: sh

    jutge upload P00001_ca_prog.cpp --compiler 'G++'

If the flag ``--problem-set`` the command will upload all problems from the
specified set found in the current working directory or in the set folder in
the current working directory. (Keep in mind that `jutge.org`_ limits the
number of submissions to 20 per hour so it is discouraged to use this flag
with large problem sets)

By default upload will test all problems against public test cases in the
database (not including custom ones). You can skip those checks with the flag
``--skip-test``

If you want to check the submitted problem verdict directly after upload, use
the flag ``--check`` which will wait for the judge verdict and output it.

License
-------

This software is licensed under the `GPL v3 license
<http://www.gnu.org/copyleft/gpl.html>`_.

.. _jutge.org: https://jutge.org
