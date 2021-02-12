t = int(input())

p = [True] * 10001

for i in range(2, int(10001 ** 0.5 + 1)):
    if p[i]:
        for j in range(2 * i, 10001, i):
            p[j] = False

for i in range(t):
    n = int(input())
    for j in range(int(n / 2), 1, -1):
        if p[j]:
            if p[n - j]:
                print(j, n - j)
                break