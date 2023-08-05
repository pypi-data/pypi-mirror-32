Plone Mosaic
============

**Plone Mosaic** is a site builder and layout solution for Plone.

Read this introduction and `the package documentation`__ for more details how to use this package.

__  http://plone-app-mosaic.s3-website-us-east-1.amazonaws.com/latest/

.. image:: https://secure.travis-ci.org/plone/plone.app.mosaic.png?branch=master
    :alt: Travis CI badge
    :target: http://travis-ci.org/plone/plone.app.mosaic

.. image:: https://coveralls.io/repos/plone/plone.app.mosaic/badge.png?branch=master
    :alt: Coveralls badge
    :target: https://coveralls.io/r/plone/plone.app.mosaic

..  image:: https://www.herokucdn.com/deploy/button.png
    :target: https://heroku.com/deploy?template=https://github.com/plone/plone.app.mosaic

Concepts
--------

Mosaic, Blocks_ and Tiles_ provide a simple, yet powerful way to manage the pages on your Plone website.
At their core, they rely on semantic HTML and resources with valid, publishable URLs.

**Mosaic Editor** editor is a visual editor for pages rendered using Blocks.
It relies on a grid system to place tiles onto a page in an intuitive, WYSIWYG, drag-and-drop manner.
Using Mosaic Editor, it is easy to compose pages with complex, balanced and visually appealing layouts.

Currently, the Mosaic Editor is activated, when any content with *Mosaic layout* view active is being edited.
(Mosaic layout is available for any content with *Layout support* behavior enabled.)

**Blocks** is a rendering algorithm based on HTML markup conventions.
A page managed by Mosaic Editor is stored as a simple HTML document.
It is representing the actual content of that page as a standalone, publishable resource devoid of any site layout content (e.g. global navigation elements).
This is referred to as *content layout*.

**Tiles** represent the dynamic portions of a page.
At its most basic level, a tile is simply an HTML document with a publishable URL.

In practice, tiles are usually implemented as browser views deriving from the ``Tile`` base class and registered with the ``<plone:tile />`` ZCML directive.
This allows tiles to have some basic metadata and automatically generated edit forms for any configurable aspects, which Mosaic will expose to users.
See `plone.tiles`_ for examples.

When work with tiles in Mosaic Editor, there are three types of tiles:

Text tiles
    Static HTML markup (WYSIWYG-edited text) placed into the content or site layout.
    Strictly speaking, text tiles are not tiles in that they do not involve any tile fetching or merging - instead they are stored as part of the page or site layout.
    To the user, however, a text tile can be moved around and managed like any other.

Field tiles
    Render the value of a metadata field such as the title or description.
    The values of field tiles may be edited in-place in the page,
    but the value is stored in the underlying field and can be indexed in the catalog, used for navigation and so on.
    In practice, a field tile is an instance of the special tile ``plone.app.standardtiles.fields`` with the field name passed as a parameter.

App tiles
    Any other type of dynamic tile. Examples may include a folder listing, a media player, a poll or pretty much anything else you can think of.

..  _Blocks: https://pypi.python.org/pypi/plone.app.blocks
..  _Tiles: https://pypi.python.org/pypi/plone.app.tiles
..  _plone.tiles: https://pypi.python.org/pypi/plone.tiles


Advanced Editor usage
---------------------

Advanced mode
    If you press the "alt" key you will be shown the layout structure, labels for your tiles and css classes for rows.

Custom classes on rows
    Also in the advanced mode, you're able to add custom classes on rows by double clicking the displayed row class.

Subcolumns
    In order to nest columns inside a cell, drag a tile, then press the "ctrl" key and drop the tile close to an existing one, either before or after it, in accordance to the shown insert marker.

Fluid rows
    For fluid (full width) rows select any tile in the row and choose "Fluid" from the "Format" menu.
    Fluid row styles only make sense on pages without portlets. In Plone 5.1.3 we can check that automatically (with plone.app.layout 2.8.0) and those styles are only active if no portlet columns are shown.


Installation
------------

**Plone Mosaic** is installed by building a Plone site with package
**plone.app.mosaic** and activating its **Plone Mosaic** add-on.

The dependencies are already version pinned in Plones ecosystem.

However, if a newer version of mosaic is needed,
the good known set for the version can be found at Github, Mosaic Code repository, in the file `versions.cfg <https://github.com/plone/plone.app.mosaic/blob/master/versions.cfg>`_.

After the add-on activation, the new content layout and editor support can be
enabled for any content type by enabling behaviors **Layout support** and
**Drafting support**.


Status
------

**Plone Mosaic** is considered to be in early production phase.
All feature are working and Mosaic is used in production.

But `there may still be bugs, which should be reported.`__

Translations are still missing.

Not all the features of Plone Mosaic have yet easily accessible UI (e.g.
layouts can be created into *portal_resources* and bound to content types as
named views only through Zope Management Interface, ZMI).

__ https://github.com/plone/plone.app.mosaic/milestones/11


Backend development
-------------------

Plone 5 version of Plone Mosaic is available for development using
``plips/plip-mosaic.cfg`` at Plone 5 coredev-buildout.

Plone 4 version of Plone Mosaic can be developed by cloning the product
directly.

Clone and build:

..  code:: bash

    $ git clone https://github.com/plone/plone.app.mosaic
    $ cd plone.app.mosaic
    $ python bootstrap.py  # clean python 2.7 virtualenv recommended
    $ bin/buildout

Startup:

..  code:: bash

    $ bin/demo

Get started:

* open a browser at ``http://localhost:55001/plone/++add++Document``
* login as ``admin`` with password ``secret``
* save the new page
* from the *Display*-menu, select the new entry *Mosaic layout*
* click *Edit* to see the new *Mosaic Editor*

Alternatively you can also use ``bin/instance fg``.

.. For impatient types, there is also an online demo installation available:
   http://plone-app-mosaic.herokuapp.com. It needs about 60 seconds to spin up and
   it will purge all changes after about an hour of non-usage.


Frontend development
--------------------

Plone Mosaic requires javascript and css bundles,
which must be manually updated for Plone 4.3.x with:

.. code:: bash

   $ make install
   $ make clean all mode=release

The bundle can also be built with source maps and watched for changes with:

.. code:: bash

   $ npm install
   $ make clean all watch


Webpack based frontent development
----------------------------------

Plone Mosaic can be developed with Webpack running:

.. code:: bash

   $ make watch_theme

or starting the instances either manually or with ``make watch_instance`` and starting the Webpack development server with:

.. code:: bash

   $ make watch_webpack

Once you have activated theme called **Plone Mosaic**,
the editor will be reloaded and rebuilt by Webpack development server after each filesystem change.


Documentation screenshots
-------------------------

To script screenshots into the Sphinx documentation, use the development buildout:

..  code:: bash

    $ git clone https://github.com/plone/plone.app.mosaic
    $ cd plone.app.mosaic
    $ python bootstrap.py  # clean python 2.7 virtualenv recommended
    $ bin/buildout -c develop.cfg

To speed up your iterations, before compiling the docs, start the robot server with:

..  code:: bash

    $ bin/robot-server plone.app.mosaic.testing.PLONE_APP_MOSAIC_ROBOT -v

With robot-server running, you can re-build the docs' screenshots relatively fast with:

..  code:: bash

    $ bin/robot-sphinx docs html

Or simply run the embedded screenshots as robot tests from a single document with:

..  code:: bash

    $ bin/robot docs/getting-started.rst

or with phantomjs:

..  code:: bash

    $ bin/robot -v BROWSER:phantomjs docs/getting-started.rst

and open ``./report.html`` to view the test report.

Just add ``Debug`` keyword anywhere to pause the robot in the middle of the screenshot script and drop you into a Robot Framework REPL.
