#!/bin/python3


def is_leap(year):
    leap = False

    assert (year >= 1990) and (year <= 10**5)
    if year % 4 == 0:
        if year % 100 == 0:
            if year % 400 == 0:
                leap = True
        else:
            leap = True

    return leap

if __name__ == "__main__":
    assert is_leap(1991) is False
    assert is_leap(1990) is False
    assert is_leap(2000) is True
    assert is_leap(2004) is True
