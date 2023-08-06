NPR
===

This module provides a simple framework for working with NPR's cloud services.

You can install this module via:

.. code-block:: python

  pip install npr

Setup:
------

Begin by authenticating your app.  Auth will walk you through key creation.

.. code-block:: python

  import npr
  npr.auth()

**output**:

.. code-block:: bash

  To authenticate your app:
    1. LOGIN to http://dev.npr.org (if it's your first time, you'll need to register.)
    2. Open the dev console (drop down in the top right corner of dev center)
    3. Create a new application
    4. Select that application and enter your credentials below
        Application ID:

Fill in your application ID and secret at the prompts.  Once verified, you must login:

.. code-block:: python

  npr.login()

**output**:

.. code-block:: bash 

  Go to https://secure.npr.org/device login and enter:
  Z3SDM6

The script will poll the npr auth server every 5 seconds until you login and it gets a token.  
Then it will store your token and you shant (SHANT!) have to do this again.

Common variables:
-----------------
The most common variables for many classes have already been loaded into the namespace, 
and you can access these in the asset dictionary:

.. code-block:: python

  stations = npr.Stations('boston')
  stations.a

**output**:

.. code-block:: bash 
    {'id': '330',
     'mp3': 'https://icecast-stream.wbur.org/wbur_nprorg',
     'name': 'WBUR',
     'station': [{'id': '330',
       'mp3': 'https://icecast-stream.wbur.org/wbur_nprorg',
       'name': 'WBUR',
       'stream': 'https://icecast-stream.wbur.org/wbur.aac'},
      {'id': '396',
       'mp3': 'https://streams.audio.wgbh.org:8200/wgbh',
       'name': 'WGBH Radio',
       'stream': 'https://streams.audio.wgbh.org:8200/wgbh'},
      {'id': '168809220', 'name': 'WGBH'}],
     'stream': 'https://icecast-stream.wbur.org/wbur.aac'}

Because they are in the namespace, you can use dot notation to access any of the first-level 
variables:

.. code-block:: python

  stations.stream

**output**:

.. code-block:: bash 

  'https://icecast-stream.wbur.org/wbur.aac'


Custom variables:
-----------------

You can also use a reverse lookup to find the keys to your own variables:

.. code-block:: python

  search = npr.Search('Hidden Brain')
  search.pretty()

**output**:

.. code-block:: bash 

  "audioTitle": "Ep. 64: I'm Right, You're Wrong",
  "date": "2017-03-13T21:00:19-04:00",
  "description": "There are some topics
    "items": [],
    "links": {
      "audio": [
        {
          "content-type": "audio/mp3",
          "href": "https://play.podtrac.com/npr-510308...
	. . . 

And, using the above output, query to find the key to **Ep. 64: I'm Right, You're Wrong**

.. code-block:: python

  search.find("Ep. 64: I'm Right, You're Wrong")

**output**:

.. code-block:: bash

  Ep. 64: I'm Right, You're Wrong .response['items'][0]['items'][2]['attributes']['audioTitle']

And now you can loop through all the recent episodes:

.. code-block:: python

  for episode in search.response['items'][0]['items']:
    print(episode['attributes']['audioTitle'])

**output**:

.. code-block:: bash

  Ep. 66: Liar, Liar
  Episode 65: Tunnel Vision
  Ep. 64: I'm Right, You're Wrong

To grab **more than the last three episodes** from this aggregation, you'll need to lookup 
the affiliate code and pass it to the **Agg class**:

.. code-block:: python

  hiddenBrain = Agg('510308')
  hiddenBrain.pretty()

Build an NPR One app:
---------------------

This won't help you play audio through a speaker, but it'll get you the data you need.  First, initialize your player:

.. code-block:: python

  player = npr.One()

Now pass the title of the story to your display and the story audio to your player, use:

.. code-block:: python

  player.title
  player.audio

To get the next segment, use:

.. code-block:: python

  player.skip()

or

.. code-block:: python

  player.complete()

...depending on the user action.  Then you call player.audio to play the next segment.

Explore Tab:
------------

The channel endpoint just lets you know what collections are available.  You'll need a distinct call for each row (collection) in the explore tab.  So to initialize the explore object and see all the stories in the third row, use:

.. code-block:: python

  explore = npr.Channels()
  explore.fetch(2)
  explore.row.pretty()

Authentication functions:
-------------------------

	| **npr.auth()** - authenticates your app with your developer credentials from dev.npr.org
	| **npr.login()** - returns a short code your user can enter at secure.npr.org/device, which will deliver a bearer token to your app
	| **npr.logout()** - removes the user's bearer token from your app.  Remember to logout before distributing your app.
	| **npr.deauth()** - removes your developer credentials from the app by deleting the npr.conf file

Endpoint classes:
-----------------

	| **npr.Station(orgId)** - returns metadata about an NPR station, where 'orgId' is the orgId of the station.
	| **npr.Stations('query')** - returns metadata about NPR stations that match a query (call letters, zip code, city, or any indexed value)
	| **npr.Stations(lat,lon)** - returns metadata about NPR stations at a location (lon should be negative, because all our stations are west of the meridian)
	| **npr.Search('query')** - returns programs or episode titles with a term that matches your 'query'
	| **npr.searchall('query')** - returns any story with a term that matches your 'query'
	| **npr.User()** - returns data (including content preferences) about the logged in user
	| **npr.Recommend()** - returns a list of recommended audio for the logged in user.
	| **npr.One()** - Like recommend, except you can advance to the next segment via skip() and complete()
	| **npr.Agg()** - returns audio segments from the selected aggregation (aka affiliation)
	| **npr.Channels()** - returns channels from the explore tab, which, along with fetch(row) will also return segments.

Endpoint helper functions:
--------------------------

	| **npr.docs()** - Lists example endpoint calls
	| <YOUR OBJECT NAME> **.a** - Lists variables loaded into the namespace of the current object.
	| <YOUR OBJECT NAME> **.response** - the json response from the endpoint
	| <YOUR OBJECT NAME> **.pretty()** - prints the json output in human-readable form
	| <YOUR OBJECT NAME> **.find('your json value')** - returns the json key path for the value you entered

Full endpoint documentation is available at http://dev.npr.org

Packaging for PyPI:
-------------------

- from npr/npr, type the command:

.. code-block:: bash

  pasteurize -w __init__.py

- open npr/tests/test.ipynb in jupyter and run some of the tests
- increment the version number in npr/setup.py and add any new dependencies

.. code-block:: bash

  version='0.1.2',
  install_requires=[
    'requests','future','requests[security];python_version<"2.9"',
  ],

- push new code to github
- from repo root (npr) build the package:

.. code-block:: bash

  python setup.py sdist bdist_wheel

- update twine (optional) and upload it to PyPI:

.. code-block:: bash

  pip install --upgrade twine
  twine upload dist/* --skip-existing

- uninstall and reinstall npr on your machine.

.. code-block:: bash

  pip uninstall npr
  pip install npr

(pat yourself on the back)


