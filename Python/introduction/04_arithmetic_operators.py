#!/bin/python3

a = int(input())
b = int(input())

if (a >= 1) and (a <= 10**10) and \
   (b >= 1) and (b <= 10**10):
    print(str(a + b))
    print(str(a - b))
    print(str(a * b))
