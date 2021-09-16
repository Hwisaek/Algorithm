import sys
from itertools import combinations

input = sys.stdin.readline

array = []

for _ in range(9):
    array.append(int(input()))

combs = combinations(array, 7)

for comb in combs:
    if sum(comb) == 100:
        answer = comb
        break

for c in sorted(list(answer)):
    print(c)
