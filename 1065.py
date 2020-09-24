n = int(input())

a = 99

if n < 100:
    print(n)
else:
    for i in range(100, n + 1):
        if int(i / 100) - int((i % 100) / 10) == int(int(i % 100) / 10) - (i % 10):
            a += 1

    print(a)