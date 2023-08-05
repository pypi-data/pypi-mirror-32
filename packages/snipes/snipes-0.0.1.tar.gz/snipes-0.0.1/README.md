SNIPES: A system for running tests with SWIFT
=============================================

Running a series of physics-based convergence tests with any code can be very
tedious. This becomes even more difficult when you are throwing around
gigabytes of data, like we do in cosmological simulations.

There are many workflow management solutions available on the market, but
they are often poorly maintained and documented. We wanted something custom
for the hydrodynamical convergence suite in SWIFT, so we built SNIPES,

+ Swift
+ coNvergence
+ physIcs
+ Python
+ Execution
+ System

Okay, it's a bit of a backronym, but _you_ try and come up with an acronym
based on a bird name. It's surprisingly hard...

Requirements
------------

The requirements for _running_ SNIPES can be found in `requirements.txt`. To
run the tests, you will need the `requirements_test.txt`. To install both,
simply run

`pip3 install -r requirements.txt -r requirements_test.txt`

To install a development version of SNIPES, simply run

`pip3 install -e .`

from the main directory.

Current Status
--------------

We have just started! We need to:

+ Implement some classes
  - `Configure` which will configure and make individual binaries with chosen
    physics properties
  - `Runtime` which will be used to generate both the `yaml` configuration file
    and the engine policies that will be used for the run
  - `ICGen` which will generate or find appropriate initial conditions
  - `Runner` which will actually perform the run with the SWIFT binary
  - `Postprocessor` which will do all of the postprocessing for that run.
+ Implement a nice front-end for users
+ Interact with the current hydrodynamical test problems that we have for SWIFT
+ Add interaction with the SLURM batch system

License
-------

SNIPES ships with the same license as SWIFT, the GPLv2.