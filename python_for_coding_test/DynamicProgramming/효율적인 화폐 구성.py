N, M = map(int, input().split())

d = [-1] * 10101

coins = []
for _ in range(N):
    coin = int(input())
    coins.append(coin)
    d[coin] = 1

for i in range(min(coins), M + 1):
    for c in coins:
        d[i + c] = min(d[i + c], d[i] + 1) if d[i + c] != -1 else d[i] + 1

print(d[M])
