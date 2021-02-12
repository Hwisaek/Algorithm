import sys

t = int(sys.stdin.readline().rstrip())

for x in range(t):
    a, b = map(int, sys.stdin.readline().split())
    print(a + b)
