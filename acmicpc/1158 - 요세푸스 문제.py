from collections import deque

n, k = map(int, input().split())

q = [i + 1 for i in range(n)]

answer = []
idx = 0
while q:
    idx = (idx + k - 1) % len(q)
    answer.append(str(q.pop(idx)))

print('<', ', '.join(answer), '>', sep='')
