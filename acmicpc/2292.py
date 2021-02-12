n = int(input())

i = 1
n_min = 1
n_max = 1

if n == 1:
    print(1)
else:
    while True:
        n_min += 6 * (i - 1)
        n_max += 6 * i
        if n >= n_min and n <= n_max:
            print(i + 1)
            break
        i += 1
