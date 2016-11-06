#!/bin/python3

N = int(input())

if (N >= 1) and (N <= 20):
    for i in range(N):
        if i >= 0:
            print(str(i**2))
