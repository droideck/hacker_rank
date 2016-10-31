#!/bin/python3

N = int(input().strip())

if (N >= 1) and (N <= 100):
    if (N % 2 != 0) or (N >= 6) & (N <= 20):
        print("Weird")
    else:
        print("Not Weird")
