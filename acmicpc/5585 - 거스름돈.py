money = 1000 - int(input())

coins = [500, 100, 50, 10, 5, 1]

answer = 0

for coin in coins:
    count, money = divmod(money, coin)
    answer += count
    if money == 0:
        break

print(answer)
