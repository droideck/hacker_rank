#!/bin/python3

from __future__ import print_function

# One way
print(*range(1, int(input()) + 1), sep='', end='')

# Or another
list(map(lambda x: print(x, end=''), range(1, int(input()) + 1)))
