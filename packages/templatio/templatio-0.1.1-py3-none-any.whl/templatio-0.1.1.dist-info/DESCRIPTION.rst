templatio
===============
.. image:: https://travis-ci.org/marcobellaccini/templatio.svg?branch=master
    :target: https://travis-ci.org/marcobellaccini/templatio

About templatio
--------------------------
templatio is a Python 3 utility that uses `TextFSM`_ and `Jinja2`_ to 
convert text files based on input and output templates.

It can be used by developers as a module or by users as a script.

It supports `TextFSM syntax`_ for input templates and 
`Jinja2 syntax`_ for output templates.

templatio is Free Software, released under the `Apache License, Version 2.0`_.

templatio was written by Marco Bellaccini (marco.bellaccini(at!)gmail.com) after a fine lunch.

QuickStart
--------------------------
Suppose that we want to convert a simple formatted text file 
to a nice html file.

Here is the input text file.

*infile.txt*:

.. code:: bash

    Name: John
    Surname: Doe
    Country: Atlantis

First, we need to create an input template using `TextFSM syntax`_.

*intempl.txt*:

.. code:: bash

    Value Name (\S+)
    Value Surname (\S+)
    Value Country (\S+)

    Start
      ^Name: ${Name}
      ^Surname: ${Surname}
      ^Country: ${Country}

Then, we need to write an output template using `Jinja2 syntax`_.

*outtempl.html*:

.. code:: html

    <!DOCTYPE html>
    <html>
      <head>
        <title>Hello</title>
      </head>
      <body>
        <p>Hi, I'm {{ Name }} {{ Surname }} and I come from {{ Country }}.</p>
      </body>
    </html>

Now, we can convert the file by running templatio as a script:

	templatio intempl.txt outtempl.html infile.txt outfile.html

And here is what we get.

*outfile.html*:

.. code:: html

    <!DOCTYPE html>
    <html>
      <head>
        <title>Hello</title>
      </head>
      <body>
        <p>Hi, I'm John Doe and I come from Atlantis.</p>
      </body>
    </html>

It's also possible to use templatio as a Python module:

.. code:: python

    import templatio

    inData = """Name: John
    Surname: Doe
    Country: Atlantis"""

    inTemplate = """Value Name (\S+)
    Value Surname (\S+)
    Value Country (\S+)

    Start
      ^Name: ${Name}
      ^Surname: ${Surname}
      ^Country: ${Country}"""

    outTemplate = """<!DOCTYPE html>
    <html>
      <head>
        <title>Hello</title>
      </head>
      <body>
        <p>Hi, I'm {{ Name }} {{ Surname }} and I come from {{ Country }}.</p>
      </body>
    </html>"""

    outData = templatio.parseInToOut(inTemplate, outTemplate, inData)

Let's get more
--------------------------
Here is a slightly more complex example, that gives you an idea of how you
can leverage `TextFSM`_ and `Jinja2`_ templates to 
perform advanced conversions.

Assume that we want to generate a json drive usage report from the output of
the `df`_ command on a system with 2 drives.

Json objects associated with the drives should have an *alarm* value set 
to *true* if disk usage is over 80%.

Here are the input and template files.

*infile.txt*:

.. code:: bash

    Filesystem     1K-blocks    Used Available Use% Mounted on
    udev             2014208       0   2014208   0% /dev
    tmpfs             405100    5848    399252   2% /run
    /dev/sda1       16447356 4893016  10699148  32% /
    /dev/sda2        1017324  893016    934423  96% /mnt/foo
    tmpfs            2025484  222424   1803060  11% /dev/shm
    tmpfs               5120       4      5116   1% /run/lock
    tmpfs            2025484       0   2025484   0% /sys/fs/cgroup
    tmpfs             405096      56    405040   1% /run/user/1000

*intempl.txt*:

.. code:: bash

    Value Drive1 (\S+)
    Value Drive1Usage (\d+)
    Value Drive2 (\S+)
    Value Drive2Usage (\d+)

    # start state
    Start
      # after parsing drive1 data, switch to Drive1parsed state
      ^/dev/${Drive1} .* ${Drive1Usage}% -> Drive1parsed

    # drive 1 parsed state
    Drive1parsed
      ^/dev/${Drive2} .* ${Drive2Usage}%

*outtempl.json*:

.. code:: bash

    {% macro checkusage(usage) -%}
        {% if usage > 80 -%}true{% else %}false{% endif %}
    {%- endmacro -%}
    {
      "drives": {
        "drive1": {
          "name": "{{ Drive1 }}",
          "usage": "{{ Drive1Usage }}",
          "alarm": {{ checkusage(Drive1Usage | int) }}
        },
        "drive2": {
          "name": "{{ Drive2 }}",
          "usage": "{{ Drive2Usage }}",
          "alarm": {{ checkusage(Drive2Usage | int) }}
        }
      }
    }

We run templatio (in this example, we use it as a script):

	templatio intempl.txt outtempl.json infile.txt report.json

And we get this nice result.

*report.json*:

.. code:: json

    {
      "drives": {
        "drive1": {
          "name": "sda1",
          "usage": "32",
          "alarm": false
        },
        "drive2": {
          "name": "sda2",
          "usage": "96",
          "alarm": true
        }
      }
    }

Much more!
--------------------------
These were just toy examples: both `TextFSM`_ and `Jinja2`_ have powerful 
template syntaxes.

After reading their documentation, 
you'll be able to perform really cool conversions!

.. _TextFSM: https://github.com/google/textfsm
.. _Jinja2: http://jinja.pocoo.org/
.. _TextFSM syntax: https://github.com/google/textfsm/wiki/TextFSM
.. _Jinja2 syntax: http://jinja.pocoo.org/docs/latest/templates/
.. _Apache License, Version 2.0: https://www.apache.org/licenses/LICENSE-2.0
.. _df: https://www.gnu.org/software/coreutils/manual/html_node/df-invocation.html


