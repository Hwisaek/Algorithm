from collections import deque

n, k = map(int, input().split())

q = deque([i + 1 for i in range(n)])

answer = '<'
while q:
    q.rotate(-(k - 1))
    answer += str(q.popleft()) + ', '

answer = answer[:-2] + '>'

print(answer)
