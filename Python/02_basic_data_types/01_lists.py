#!/bin/python3

n = input()
ls = []

for i in range(int(n)):
    cmd = input()
    if cmd == "print":
        print(ls)
    else:
        cmd_ls = cmd.split()
        if len(cmd_ls) == 1:
            eval("ls.{}()".format(cmd_ls[0]))
        if len(cmd_ls) == 2:
            eval("ls.{}({})".format(cmd_ls[0], cmd_ls[1]))
        elif len(cmd_ls) == 3:
            eval("ls.{}({}, {})".format(cmd_ls[0], cmd_ls[1], cmd_ls[2]))
