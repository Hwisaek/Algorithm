N = int(input())
array = []
for _ in range(N):
    # t: 상담 소요 기간, p: 상담 금액
    array.append(list(map(int, input().split())))

money = [0] * (N + 1)

for idx, arr in enumerate(array):
    t, p = arr
    money[idx] = max(money[idx], money[idx - 1])
    if idx + t < (N + 1):
        money[idx + t] = max(money[idx + t], money[idx] + p)

print(max(money))
