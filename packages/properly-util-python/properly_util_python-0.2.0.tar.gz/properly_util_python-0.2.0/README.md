# properly-util-python

##Installation

`pip install -e git+https://github.com/GoProperly/properly-util-python@0.1#egg=properly_util_python`

see:https://stackoverflow.com/questions/15268953/how-to-install-python-package-from-github#comment37317873_15268990

Based on this tutorial:
http://greenash.net.au/thoughts/2015/06/splitting-a-python-codebase-into-dependencies-for-fun-and-profit/

Currently we are hosting our package in Git since it is quick and easy to set up.
However, that has a number of disadvantages, the most visible disadvantage being that pip install will run much slower, because it has to do a git pull every time you ask it to check that foodutils is installed (even if you specify the same commit / tag each time).
http://carljm.github.io/tamingdeps/#33