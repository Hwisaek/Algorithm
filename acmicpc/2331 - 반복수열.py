from collections import defaultdict

A, P = map(int, input().split())

d = []
d.append(A)

count = defaultdict(int)
count[A] += 1

while True:
    n = d[-1]
    new = sum([int(i) ** P for i in str(n)])
    d.append(new)
    count[new] += 1
    if count[new] == 2:
        break

print(d.index(new))
