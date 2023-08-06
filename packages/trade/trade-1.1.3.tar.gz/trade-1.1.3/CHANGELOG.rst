Change Log
==========

1.1.3 (2018-06-02)
------------------
* fix: Removed unused *copy* import in *subject.py*.

1.1.2 (2018-03-19)
------------------
* Updated docs.

1.1.0 (2017-02-27)
------------------
* "details" as a optional argument for Subject creation. Empty dict assumed as default.

1.0.0 (2017-02-27)
------------------
* Core framework apart from example implementations.

0.4 (2017-12-03)
----------------

* Accumulator no longer as a external dependency

0.3 (2016-06-01)
----------------

* OperatioContainer now runs everything as tasks
* All tasks are defined in container_tasks.py
* Events are now defined as subclasses of Operation


0.2.9 (2016-05-29)
------------------

* Options and events are no longer part of the core framework. An example
  of an app developed on top of the **trade** framework is now on the *sample*
  folder and contains the old implementations for calls and puts, as well
  as the events representing stock splits, reverse stock splits and bonus shares.
* The JSON interface now must be initialized with the types of assets
  that will be present in the json (like "Option", "Asset"), and also with the
  tasks for the OperationContainer to run (check the tests to see how this works).


0.2.8 (2015-11-17)
------------------

* Prototype for json options and option exercises support
* Faster day trade identification
