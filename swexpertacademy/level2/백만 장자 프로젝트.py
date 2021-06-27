T = int(input())

for test_case in range(1, T + 1):
    n = input()
    price = list(map(int, input().split()))
    profit, last = 0, -1
    for n in reversed(price):
        if n < last:
            profit += last - n
        else:
            last = n
    print('#{} {}'.format(test_case, profit))
