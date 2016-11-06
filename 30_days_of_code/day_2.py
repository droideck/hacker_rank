#!/bin/python3

meal_cost = float(input())
tip_percent = int(input())
tax_percent = int(input())

tip = meal_cost * tip_percent / 100
tax = meal_cost * tax_percent / 100
total_cost = meal_cost + tip + tax

print("The total meal cost is {} dollars.".format(int(round(total_cost))))
