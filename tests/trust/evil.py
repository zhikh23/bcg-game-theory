#!/usr/bin/python3

from random import random

while True:
    role = input()
    if role == "I": # Investor
        m = int(input())
        x = m // 3 + int(random() * m // 3)
        print(x)
        returned = input()
    else: # Trustee
        x = int(input())
        print(0)
