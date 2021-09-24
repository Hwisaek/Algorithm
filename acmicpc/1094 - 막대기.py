import sys

input = sys.stdin.readline

x = int(input())

n = [64]

count = 1
while sum(n) > x:
    a = n.pop()
    n.extend([a / 2, a / 2])
    count += 1
    if sum(n[:-1]) >= x:
        n.pop()
        count -= 1

print(count)
