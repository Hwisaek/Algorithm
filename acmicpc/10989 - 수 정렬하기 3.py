import sys
from collections import defaultdict

input = sys.stdin.readline

N = int(input())

d = defaultdict(int)
for _ in range(N):
    n = int(input())
    d[n] += 1

for k in sorted(list(d.keys())):
    for _ in range(d[k]):
        sys.stdout.write(f'{k}\n')
