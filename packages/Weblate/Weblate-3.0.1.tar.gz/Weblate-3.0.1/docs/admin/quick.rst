Quick setup guide
=================

.. note::

    This is just a quick guide for installing and starting to use Weblate for
    testing purposes. Please check :ref:`install` for more real world setup
    instructions.

Installing from sources
-----------------------

#. Install all required dependencies, see :ref:`requirements`.

#. Grab Weblate sources (either using Git or download a tarball) and unpack
   them, see :ref:`install-weblate`.

#. Copy :file:`weblate/settings_example.py` to :file:`weblate/settings.py` and
   adjust it to match your setup. You will at least need to configure the database
   connection (possibly adding user and creating the database). Check
   :ref:`config` for Weblate specific configuration options.

#. Create the database which will be used by Weblate, see :ref:`database-setup`.

#. Build Django tables, static files and initial data (see
   :ref:`tables-setup` and :ref:`static-files`):

   .. code-block:: sh

        ./manage.py migrate
        ./manage.py collectstatic
        ./scripts/generate-locales # If you are using Git checkout

#. Configure webserver to serve Weblate, see :ref:`server`.


Installing using Docker
-----------------------

#. Clone weblate-docker repo:

   .. code-block:: sh

        git clone https://github.com/WeblateOrg/docker.git weblate-docker
        cd weblate-docker

#. Start Weblate containers:

   .. code-block:: sh

        docker-compose up

.. seealso::

    See :ref:`docker` for more detailed instructions and customization options.

Installing on OpenShift 2
-------------------------

#. You can install Weblate on OpenShift PaaS directly from its Git repository using the OpenShift Client Tools:

   .. parsed-literal::

        rhc -aweblate app create -t python-2.7 --from-code \https://github.com/WeblateOrg/weblate.git --no-git

#. After installation everything should be preconfigured and you can immediately start to add a translation
   project as described below. 
   
.. seealso::
   
    For more information, including on how to retrieve the generated admin password, see :ref:`openshift`.

Adding translation
------------------

#. Open admin interface (http://localhost/admin/) and create project you
   want to translate. See :ref:`project` for more details.

   All you need to specify here is project name and its website.

#. Create component which is the real object for translating - it points to
   VCS repository and selects which files to translate. See :ref:`component`
   for more details.

   The important fields here being component name, VCS repository address and
   mask for finding translatable files. Weblate supports a wide range of formats
   including Gettext PO files, Android resource strings, OS X string properties,
   Java properties or Qt Linguist files, see :ref:`formats` for more details.


#. Once the above is completed (it can be lengthy process depending on size of
   your VCS repository and number of messages to translate), you can start
   translating.
