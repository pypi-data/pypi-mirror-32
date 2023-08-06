YSchema
=======

YSchema is a terse and simple schema format with a reference validator
implementation for YAML, JSON and other dictionary based data structures.

YSchema is quite minimal (in terms of lines of code) and is continuously tested
against a set of of valid and invalid example data (see the ``examples``
directory). The code works nicely for its intended purpose, but may not be the
most powerful or popular, even if it does what it was intended for very well.
The main assumption (at least for now) is that all keys are strings without
whitespace.

YSchema is written in Python (v. 3) and validates dictionaries containing basic
datatypes like strings, ints, floats, lists and nested dictionaries. The schema
is also a dictionary, so both the data and the schema can be written in Python,
JSON, YAML, TOML, ... formats. YSchema cannot validate all possible YAML / JSON
data, in fact it cannot even validate its own schema files since those use
significant white space in dictionary keys to describe expected data types and
whether the data is required or not.

To install the YSchema Python library along with the ``yschema`` command line
program run:

.. code:: bash

    python3 -m pip install -U yschema

Consider using a virtual environment or adding ``--user`` to the ``pip`` command
if you do not want to install into the system's site-packages directory. PS: You
may also want to look at older and more established schema and validators such
as Yamale_ or json-schema_ in case those serve your needs better.

.. _Yamale: https://github.com/23andMe/Yamale
.. _json-schema: http://json-schema.org


.. contents::


Introduction to YSchema
-----------------------

A simple example schema:

.. code:: yaml

    # There must be a key "name" that maps to a string
    required name: str
    
    # There can be an integer age, but it is not required
    optional age: int
    
    # The optional height must be above 0
    optional height: float(min_val=0)

To validate this, first load the schema above into a dictionary, then load the
data to validate into another dictionary, and finally run:

.. code:: python

    import yschema
    
    # possibly loaded from json or yaml or just a plain old dict
    schema = my_load_schema_function()
    data_dict = {'name': 'Tormod'}
    
    yschema.validate(data_dict, schema_dict)

If the function does not raise ``yschema.ValidationError`` then the data is
valid according to the given schema. You can also use the ``yschema`` command
to validate YAML files from the command line.

A more complicated example, showing constants and nested dictionaries:

.. code:: yaml

    # Example of a constant that can be used in validation functions
    constant minimum_string_length: 5
    
    # A sub-dictionary
    type Whale: 
        # The name is a string of a given minimum length
        required name: str(min_len=minimum_string_length)
        
        # The length must be between 0 and 500 meters
        optional length: float(min_val=0, max_val=500.0)
    
    required whales: list(type=Whale)

The above schema validates data like this:

.. code:: yaml

    whales:
      - name: Unknown Whale
      - name: Enormous Whale
        length: 200.0

Note that when working with aliases and types the order of the keys in the
dictionary starts to matter. Either use a Python 3.6 or later, or load your
schema into an OrderedDict. YSchema contains a helper function for ordered safe
loading of YAML files:

.. code:: python

    with open(schema_file_name, 'rt') as yml:
        schema = yschema.yaml_ordered_load(yml)


More advanced features
----------------------

**Built in types**: the following types are implemented. Optional parameters
are listed below each type:

* Any
* bool
* str
    - min_len
    - max_len
    - equals - e.g. ``str(equals='Hi!')`` or matching one of several
      pissibilities with ``str(equals=('a', 'b', 'c'))``
    - prefix
* int
    - min_val
    - max_val
    - equals - e.g. ``int(equals=3)`` or ``int(equals=(2, 4, 6))``
* float
    - min_val
    - max_val
    - equals - e.g. ``float(equals=3.2)`` or ``float(equals=(2.1, 4.4))``
* list
    - min_len
    - max_len
    - type - e.g. ``list(type=int)`` or ``list(type=Whale)``
* one_of
    - types - e.g. ``one_of(types=(int, str))`` or
      ``one_of(types=(str(prefix='Moby'), Whale))``
* any_of
    - types - see ``one_of`` (``any_of`` matches if any of the types match, 
      ``one_of`` requires exactly one match)

**Alias**: you can give an alias to avoid typing the same type definition over
and over again:

.. code:: yaml

    alias Cat: one_of(types=(HouseCat, Tiger, Lynx))
    alias Cats: list(type=Cat)

**Glob**: you can allow undefined keys by using a glob. The following will
validate OK for all documents

.. code:: yaml

    optional *: Any

**Inherit**: a sub-schema introduced by ``type`` can contain a key ``inherit``
with the name of a previously defined sub-schema to avoid repeating 
definitions that are shared among several types:

.. code:: yaml

    type MeshBase:
        optional move: list(type=str)
        optional sort_order: list(type=int)
        optional mesh_file: str
    type MeshDolfinFile:
        inherit: MeshBase
        required type: str(equals=('XML', 'XDMF', 'HDF5'))
        required mesh_file: str
        optional facet_region_file: str
    type MeshMeshio:
        inherit: MeshBase
        required type: str(equals='meshio')
        required mesh_file: str
        optional meshio_type: str
    required mesh: one_of(types=(MeshMeshio, MeshDolfinFile))
 

Releases
--------

Version 1.0.2 - June 11. 2018

Improve error messages and add convinience function to safe-load YAML into an
OrderedDict

Version 1.0.1 - June 7. 2018
............................

Completed v 1.0 implementation goals. The YSchema language is powerful enough to
express most of what I wanted for validating Ocellaris_ input files. The code
base is decently tested (using the fantastic CircleCI service) and a command
line tool is also included for validating YAML files from the shell.

There may not be a large number of additional releases if no more features are
found to be necessary for the author's uses. It is relatively easy to add new
type validators from user code, but feel free to submit a pull request if you
are finding YSchema useful and have implemented some general purpose validators.
YSchema does not intend to compete with complex and more fully featured schema
languages like json-schema_.

.. _Ocellaris: https://bitbucket.org/trlandet/ocellaris

Copyright and license
---------------------

YSchema is copyright Tormod Landet, 2018. YSchema is licensed under the Apache
2.0 license, a permissive free software license compatible with version 3 of
the GNU GPL. See the file LICENSE for the details.
