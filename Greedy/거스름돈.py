coins = [500, 100, 50, 10]

change = int(input('거스름돈을 입력하세요 >>'))

# 동전개수
count = 0
for coin in coins:
    count += change // coin
    change %= coin

print('동전의 최소 개수: ', count)
