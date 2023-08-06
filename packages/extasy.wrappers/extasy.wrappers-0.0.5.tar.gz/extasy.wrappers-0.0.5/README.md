# ExTASY Wrappers #

A library to simplify the construction of ExTASY simulation-analysis workflows.

### What is this repository for? ###

* *ExTASY wrappers* provide a consistent API for launching command line applications (such as MD simulations or simulation analysis jobs) from python scripts. Though designed originally with ExTASY simulation-analysis workflows in mind, they are extensible to many other possible applications. 

In a nutshell, the construction and execution of a complete simulation or analysis job is reduced to a python function call of the form:
```
#!python
results = kernel.run(inputs)
```
where "kernel" might be pmemd, grompp, pyCoCo, LSDMap, etc.

See the Wiki for more details.

* Version 0.0.1

### How do I get set up? ###

* Clone the repository and then run setup.py:
```
git clone https://bitbucket.org/extasy-project/wrappers.git
cd wrappers
python setup.py install --user
```

* Dependencies: *ExTASY wrappers* itself has no dependencies outside the standard python libraries. To run the examples though will require:

1. An MD code: Amber (or Ambertools) and/or GROMACS
2. A sampling analysis package: [ExTASY CoCo](https://bitbucket.org/extasy-project/coco) and/or [ExTASY-compliant LSDMap](https://github.com/ClementiGroup/LSDMap)

Running the parallel examples will additionally require:
[Dask](https://github.com/blaze/dask) and/or [mpipool](https://github.com/adrn/mpipool/tree/master/mpipool)

* How to run tests: See the README files in the folders under ./examples.


### Who do I talk to? ###

* Charlie Laughton