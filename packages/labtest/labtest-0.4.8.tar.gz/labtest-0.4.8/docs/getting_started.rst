===============
Getting Started
===============

LabTest has three different parts: the laboratory, the project configuration and the client that runs on your machine.

What do you need to do?

- :ref:`I just need to get it working on my machine<getting_started:setting up your local machine>`
- :ref:`I need to get my project ready<getting_started:Getting your code ready>`
- :ref:`I need a laboratory<admin/setting_up_the_lab:Setting up the laboratory>`


Setting up your local machine
=============================

This is a one-time process. It installs the Lab Test package and configures your machine to easily talk to the test server.

Install labtest
---------------

First we :ref:`install <install_stable>` the Lab Test command line package:

.. code-block:: console

    $ pip install labtest

Public key in IAM
-----------------

Make sure your public key was added to your AWS IAM account. Without that, you will not be able to SSH into anything.

Configure SSH
-------------

Let's set up our SSH configuration. We need a few bits of information:

- SSH bastion DNS name or IP address
- The test server IP address (it is a non-routable IP address, like 10.x.x.x)
- Your user name. If your username contains ``+``\ , ``=``\ , ``,``\ , or ``@`` you need to convert a few characters:

  - ``+`` to ``.plus.``
  - ``=`` to ``.equal.``
  - ``,`` to ``.comma.``
  - ``@`` to ``.at.``

For this example:

- **SSH bastion IP address:** ``111.222.111.222``
- **Test server IP address:** ``10.20.3.3``
- **User name:** ``corey.oordt.at.boston.gov`` (converted from ``corey.oordt@boston.gov``\ )

Now we add some lines to our ``~/.ssh/config`` file:

.. code-block:: none
    :caption: The addition to the ``~/.ssh/config`` file.

    Host bastion
    Hostname 111.222.111.222
    Port 22
    User corey.oordt.at.boston.gov
    IdentityFile ~/.ssh/id_rsa

    Host test
    Hostname 10.20.3.3
    User corey.oordt.at.boston.gov
    Port 22
    ProxyCommand ssh -A -T bastion nc %h %p
    IdentityFile ~/.ssh/id_rsa

With that in place, you should be able to :command:`ssh` to the test server:

.. code-block:: console
    :caption: SSH'ing to the test server

    $ ssh test
    Last login: Sun May  6 15:18:17 2018 from ip-10-20-2-195.ec2.internal

           __|  __|_  )
           _|  (     /   Amazon Linux 2 AMI
          ___|\___|___|

    https://aws.amazon.com/amazon-linux-2/
    No packages needed for security; 56 packages available
    Run "sudo yum update" to apply all updates.
    [corey.oordt.at.boston.gov@ip-10-20-10-41 ~]$

You can disconnect by typing :kbd:`control-d` or :kbd:`exit`.


Getting your code ready
=======================

At this point, you may want to read the section on :ref:`how experiments work<admin/how_it_works:experiments>` to understand the process in more depth. Ultimately we will need a Docker image and a configuration when we are done.

Containerize it
---------------

Each experiment runs in a Docker container and is configured via environment variables. If your app doesn't already have a ``Dockerfile`` and a way to build everything as a container, you need to adapt it.

This topic is too broad to go into here, but the :ref:`Tutorial <tutorials/containerizing:containerizing>` demonstrates a very simple conversion. You'll know you are ready when you can run something like:

.. code-block:: console

    $ docker build -t myapp .
    $ docker run --rm -ti myapp

That means your container builds and runs locally.

.. _automating-the-app-build-process:

Automating the app build process
--------------------------------

LabTest doesn't really care how you generate a Docker image. That said, there is a built-in process that will build your app and Docker image on the laboratory server.
