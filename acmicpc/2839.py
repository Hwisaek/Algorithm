n = int(input())

if n == 1 or n == 2 or n == 4 or n == 7:
    print(-1)
else:
    if (n % 5) % 3 == 0:
        n_5 = n // 5
        n_3 = (n % 5) // 3
    else:
        if (n - (n // 5 - 1) * 5) % 3 == 0:
            n_5 = n // 5 - 1
            n_3 = (n - n_5 * 5) // 3
        elif (n - (n // 5 - 2) * 5) % 3 == 0:
            n_5 = n // 5 - 2
            n_3 = (n - n_5 * 5) // 3
    print(n_5 + n_3)
