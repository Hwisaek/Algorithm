t = int(input())


for i in range(t):
    k = int(input())
    n = int(input())
    p = [i for i in range(1, n + 1)]
    for _ in range(k):
        for j in range(1, n):
            p[j] += p[j - 1]
    print(p[-1])