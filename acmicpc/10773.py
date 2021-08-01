import sys

K = int(sys.stdin.readline())
total = []
prev = 0
for i in range(K):
    n = int(sys.stdin.readline())
    if n != 0:
        total.append(n)
    else:
        total.pop()

print(sum(total))