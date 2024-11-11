#!/usr/bin/python3

from random import random

while True:
    whoami = input()
    summ = int(input())
    if whoami == "A":
        print(int(summ * random()))
    else:
        offer = int(input())
        if offer / summ < 0.5:
            print("N")
        else:
            print("Y")
