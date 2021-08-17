import sys
import heapq

input = sys.stdin.readline
n = int(input())

card = []
for i in range(n):
    heapq.heappush(card, int(input()))

answer = 0
while len(card) != 1:
    a, b = heapq.heappop(card), heapq.heappop(card)
    now = a + b
    answer += a + b
    heapq.heappush(card, now)

print(answer)
