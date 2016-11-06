#!/bin/python3

n = int(input())
n_list = input()
tp = tuple(map(int, n_list.split()))

print(hash(tp))
