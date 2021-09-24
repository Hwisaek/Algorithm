from collections import deque

dx = [-2, -1, 1, 2]
dy = [1, 2, 2, 1]

n, m = map(int, input().split())

array = [[False] * m for _ in range(n)]

q = deque()
q.append((n-1, 0))
array[n-1][0] = True
count = 1

while q:
    x, y = q.popleft()
    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        if -1 < nx < n and -1 < ny < m:
            if not array[nx][ny]:
                q.append((nx, ny))
                array[nx][ny] = True
                count += 1

print(count)
