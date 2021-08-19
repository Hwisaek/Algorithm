N = int(input())
array = list(map(int, input().split()))
array.reverse()

d = [1] * N

for i in range(1, N):
    for j in range(i):
        if array[j] < array[i]:
            d[i] = max(d[i], d[j] + 1)

print(N - max(d))
