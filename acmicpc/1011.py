t = int(input())

for _ in range(t):
    x, y = map(int, input().split())

    d = y - x

    n = (d ** 0.5) * 2 - 1

    if n != int(n):
        n = int(n) + 1
    else:
        n = int(n)
    print(n)
