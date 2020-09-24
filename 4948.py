n = 123456 * 2 + 1

p = [True] * n


for i in range(2, int(n ** 0.5 + 1)):
    if p[i]:
        for j in range(2 * i, n, i):
            p[j] = False

while True:
    n = int(input())

    if n == 0:
        break

    count = 0

    for i in range(n + 1, n * 2 + 1):
        if p[i]:
            count += 1
    print(count)