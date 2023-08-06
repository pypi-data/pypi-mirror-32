========================
Experiment Configuration
========================

In order for LabTest to create an experiment, it needs a little bit of information about your project. LabTest :ref:`requires <reference/configuration:required configuration options>` very little configuration, but allows for lots of :ref:`customization <reference/configuration:optional configuration options>`\ .


Automatic configuration files
=============================

Most code repositories have enough configuration files, and we didn't want to add *another* configuration file. Lab Test will automatically look for its configuration information  in several files already in your repository, under a ``labtest`` section:

- ``.labtest.yml``
- ``setup.cfg``
- ``package.json``

Alternate configuration files
=============================

You can alternatively pass the configuration file to Lab Test at the command line with the ``--config`` option. Lab Test supports ``.ini``, ``.yml/.yaml``, and ``.json`` formats.


.. literalinclude:: /../tests/fixtures/config.ini
   :language: ini
   :caption: Example ``.ini`` configuration


.. literalinclude:: /../tests/fixtures/config.yml
   :language: yaml
   :caption: Example ``.yaml`` configuration


.. literalinclude:: /../tests/fixtures/config.json
   :language: json
   :caption: Example ``.json`` configuration

.. _configuration_required_options:

Required configuration options
==============================

There are several options that are required in order for Lab Test to work correctly.

.. _configuration:host:

``host``
--------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``None``
    * - Required:
      - ``True``
    * - Acceptable values:
      - String or IP Address

The DNS name, IP address or SSH config ``Host`` of the test server. You need this to connect to the laboratory.


.. _configuration:test_domain:

``test_domain``
---------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``None``
    * - Required:
      - ``True``
    * - Acceptable values:
      - String

The DNS subdomain in which the test server lives. This is the wildcard DNS name without the ``*.``\ , like ``test.example.com``\ . This is used with :ref:`configuration:host_name_pattern` to create the virtual host name.

.. _configuration_optional_options:

Optional configuration options
==============================

.. contents::
    :local:


.. _configuration:app_build_command:

``app_build_command``
---------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``None``
    * - Required:
      - ``False``
    * - Acceptable values:
      - Strings

The script or command to run to build the app within the Docker image. This is required if you want to build the application on the test server. Also set the :ref:`configuration:app_build_image` option.

This command is executed from within the container and in the project's directory.

For example:

.. code-block:: yaml

    app_build_command: npm run build

If you require several commands, you will need to create a script in your repository that we can run:

.. code-block:: yaml

    app_build_command: ./bin/build_my_app

If the execute bit is not set you must include the name of the program to execute the script, for example:

.. code-block:: yaml

    app_build_command: python ./config/build.py

or:

.. code-block:: yaml

    app_build_command: bash bin/build_my_app


.. _configuration:app_build_image:

``app_build_image``
-------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``None``
    * - Required:
      - ``False``
    * - Acceptable values:
      - String

The Docker image to use to build the app. `Shippable`_ has some great `images publicly available`_\ . Here is their `docker page`_\ . This is required if you want to build the application on the test server. Also set the :ref:`configuration:app_build_command` option.

.. _shippable: https://www.shippable.com/
.. _images publicly available: http://docs.shippable.com/platform/runtime/machine-image/ami-overview/
.. _docker page: https://hub.docker.com/u/drydock/


.. _configuration:app_name:

``app_name``
------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - The name of the project directory
    * - Required:
      - ``False``
    * - Acceptable values:
      - Strings

The name of the application. Ideally this should be a URL-friendly value. In order to get the default value, the labtest command must be made from within a Git repository.


.. _configuration:before_start_command:

``before_start_command``
------------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``None``
    * - Required:
      - ``False``
    * - Acceptable values:
      - Strings

The script or command to run from within the newly-built Docker image before the service is started.

This command is executed from within the container as the experiment will run; you will have access to all backing services. Its working directory will depend on your project's Dockerfile and its `WORKDIR`_ setting.


.. _workdir: https://docs.docker.com/engine/reference/builder/#workdir

.. warning:: Any scripts you need to call must be in the container. The project directory from the host will *not* be mounted.

For example:

.. code-block:: yaml

    app_build_command: ./manage.py migrate

If you require several commands, you will need to create a script in your repository that we can run:

.. code-block:: yaml

    app_build_command: ./bin/update_database

If the execute bit is not set you must include the name of the program to execute the script, for example:

.. code-block:: yaml

    app_build_command: python ./config/migrate.py

or:

.. code-block:: yaml

    app_build_command: bash bin/update_my_app



.. _configuration:build_provider:

``build_provider``
------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``local``
    * - Required:
      - ``False``
    * - Acceptable values:
      - ``local``

This is how your application and Docker image are built. Currently only ``local`` is supported.

You must also set :ref:`configuration:code_repo_url`, :ref:`configuration:app_build_image`, :ref:`configuration:app_build_command`. If the default for :ref:`configuration:container_build_command` doesn't work for your project, set that too.


.. _configuration:code_repo_url:

``code_repo_url``
-----------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``None``
    * - Required:
      - ``False``
    * - Acceptable values:
      - URL

The URL of the code repository to check out.


.. _configuration:container_build_command:

``container_build_command``
---------------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - (reformatted for clarity)

        .. code-block:: bash

            docker build \
                -t $APP/$INSTANCE \
                --build-arg RELEASE=$RELEASE \
                --build-arg APP_NAME=$APP \
                --build-arg BRANCH_NAME=$BRANCH \
                --build-arg INSTANCE_NAME=$INSTANCE \
                .
    * - Required:
      - ``False``
    * - Acceptable values:
      - Docker build command


This is the command to use to build the container for your app.

Lab Test appends this command to a script that sets the following variables: ``$APP_NAME``, ``$INSTANCE_NAME``, ``$BRANCH_NAME``, and ``$RELEASE``.

If you override the docker build command, you *must* still tag it with ``$APP/$INSTANCE`` or the remaining commands will fail.

.. note::

    If your ``Dockerfile`` doesn't use the default ``--build-arg``\ s passed, they are ignored.


.. _configuration:container_provider:

``container_provider``
----------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``local``
    * - Required:
      - ``False``
    * - Acceptable values:
      - ``local`` or ``aws``

This is to extend how Lab Test can talk to different Docker container repositories. Currently only ``local`` and ``aws`` are supported.


.. _configuration:docker_image_pattern:

``docker_image_pattern``
------------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``%(APP_NAME)s/%(INSTANCE_NAME)s:latest``
    * - Required:
      - ``False``
    * - Acceptable values:
      - String with placeholders

The string pattern to use when dynamically determining which image to use to build the container. Allows `Python string interpolation formatting`_\ , with ``APP_NAME`` and ``INSTANCE_NAME`` in the context.

The value of this option depends on how your Docker images are built. (See `docker pull documentation`_ for more information about specifying images) If they are built using the :ref:`built-in <reference/builtin_build_process:Built-in build process>` Lab Test method (the default), then the images will be local to the test server and can use a simple name. If the Docker images are built using an external process and in a private repo, the name will look like a URL, without the ``https://``\ .

.. _docker pull documentation: https://docs.docker.com/engine/reference/commandline/pull/



.. _configuration:environment:

``environment``
---------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``[]``
    * - Required:
      - ``False``
    * - Acceptable values:
      - A sequence of strings

A list of environment variable strings to include in the Docker container.

.. code-block:: yaml
   :caption: Example in YAML

   labtest:
     environment:
       - FOO=bar
       - TEST=true
       - DEBUG=true

.. code-block:: json
   :caption: Example in JSON

   {
     "labtest": {
       "environment": [
         "FOO=bar",
         "TEST=true",
         "DEBUG=true"
       ]
     }
   }

.. code-block:: ini
   :caption: Example in INI

   [labtest]
   environment = FOO=bar,TEST=true,DEBUG=true


.. _configuration:host_name_pattern:

``host_name_pattern``
---------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``%(APP_NAME)s-%(INSTANCE_NAME)s``
    * - Required:
      - ``False``
    * - Acceptable values:
      - String with placeholders suitable for a URL

The pattern to use to generate the host part of the DNS address. This is used with :ref:`configuration:test_domain` to generate the virtual host name.

This pattern may contain `Python string interpolation formatting`_. The context passed to this includes:

- ``APP_NAME``
- ``INSTANCE_NAME``
- ``BRANCH_NAME``

.. _python string interpolation formatting: https://docs.python.org/3/library/stdtypes.html#old-string-formatting


.. _configuration:use_ssh_config:

``use_ssh_config``
------------------

.. list-table::
    :class: uk-table uk-table-striped uk-table-small
    :widths: 33 64
    :stub-columns: 1

    * - Default:
      - ``False``
    * - Required:
      - ``False``
    * - Acceptable values:
      - ``True`` or ``False``

Use your local SSH config when connecting. For example, if you set this option to ``True`` and add these items (changing the various ``User`` and ``Hostname`` values) to your ``~/.ssh/config``\ :

.. code-block:: none

   Host bastion
   Hostname 111.222.111.222
   Port 22
   User monty.python.at.boston.gov
   IdentityFile ~/.ssh/id_rsa

   Host test
   Hostname 10.20.10.5
   User monty.python.at.boston.gov
   Port 22
   ProxyCommand ssh -A -T bastion nc %h %p
   IdentityFile ~/.ssh/id_rsa

You can now set the :ref:`configuration:host` configuration to ``test`` and it will route everything through the SSH bastion in the test environment. You can even ``ssh test`` from the command line.

