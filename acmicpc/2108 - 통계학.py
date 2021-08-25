from collections import Counter
import sys

input = sys.stdin.readline

N = int(input())
array = []
for _ in range(N):
    array.append(int(input()))

array.sort()
count = Counter(array)
mode = count.most_common(2)
result = mode[0][0]
if N > 1 and mode[1][1] == mode[0][1]:
    result = mode[1][0]

print(round(sum(array) / N))  # 산술평균
print(array[N // 2])  # 중앙값
print(result)  # 최빈값
print(array[-1] - array[0])  # 범위
