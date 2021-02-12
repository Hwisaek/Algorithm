m, n = map(int, input().split())

for i in range(m, n + 1):
    p = 1
    if i == 1:
        continue
    n = int(i ** 0.5)
    for j in range(2, n + 1):
        if i % j == 0:
            p = 0
            break
    if p == 0:
        continue
    print(i)