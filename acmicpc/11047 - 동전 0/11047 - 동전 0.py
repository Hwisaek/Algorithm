N, K = map(int, input().split())
coins = []
for _ in range(N):
    coins.append(int(input()))

answer = 0

for i in range(N, 0, -1):
    if K >= coins[i - 1]:
        answer += K // coins[i - 1]
        K %= coins[i - 1]
        if K == 0:
            break

print(answer)
