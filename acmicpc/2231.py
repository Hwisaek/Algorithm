n = int(input())

if n == 1:
    print(0)

for i in range(1, n):
    m = i

    sum = i

    while True:
        if m / 10 == 0:
            break
        sum += m % 10
        m = m // 10
    if sum == n:
        print(i)
        break
    if i == n - 1:
        print(0)
        break