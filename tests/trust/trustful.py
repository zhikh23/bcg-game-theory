#!/usr/bin/python3

from random import random

while True:
    role = input()
    if role == "I": # Investor
        m = int(input())
        x = m // 2 + int(random() * m // 2)
        print(x)
        returned = input()
    else: # Trustee
        x = int(input())
        y = x + int(random() * x)   # y < 3x
        print(y)
