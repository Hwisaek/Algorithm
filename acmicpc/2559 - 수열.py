import sys

input = sys.stdin.readline

n, k = map(int, input().split())

array = list(map(int, input().split()))

cum_sum = [0] * (n + 1)

for i in range(1, n + 1):
    cum_sum[i] = cum_sum[i - 1] + array[i - 1]

max_temp = -100000001
for i in range(n - k + 1):
    max_temp = max(max_temp, cum_sum[i + k] - cum_sum[i])

print(max_temp)
