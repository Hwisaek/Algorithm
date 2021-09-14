import sys

input = sys.stdin.readline

n, h = map(int, input().split())  # 길이, 높이

high = [0] * h  # 종유석
low = [0] * h  # 석순

for j in range(n // 2):
    i = int(input())
    high[i - 1] += 1
    i = int(input())
    low[h - i] += 1
if n % 2 == 1:
    i = int(input())
    low[h - i] += 1

for i in range(1, h):
    high[h - i - 1] += high[h - i]
    low[i] += low[i - 1]
answer = [high[i] + low[i] for i in range(h)]

min = 200001
count = -1
for idx in range(len(answer)):
    if answer[idx] < min:
        min = answer[idx]
        count = 1
    elif answer[idx] == min:
        count += 1

print(min, count)
