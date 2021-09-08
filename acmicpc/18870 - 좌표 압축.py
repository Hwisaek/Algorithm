import sys

input = sys.stdin.readline

N = int(input())
array = list(map(int, input().split()))
s = sorted(list(set(array)))

d = {}
for idx, value in enumerate(s):
    d[value] = idx

for n in array:
    print(d[n], end=' ')
