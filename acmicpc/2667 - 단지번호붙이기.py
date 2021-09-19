import sys
from collections import deque

input = sys.stdin.readline

n = int(input())

array = [list(input()) for _ in range(n)]
visited = [[False] * n for _ in range(n)]

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]

counts = []
q = deque()
for a in range(n):
    for b in range(n):
        if array[a][b] == '1':
            q.append((a, b))
            count = 1 if array[a][b] == '1' and not visited[a][b] else 0
            visited[a][b] = True

            while q:
                x, y = q.pop()

                for i in range(4):
                    nx, ny = x + dx[i], y + dy[i]
                    if -1 < nx < n and -1 < ny < n:
                        if not visited[nx][ny] and array[nx][ny] == '1':
                            q.append((nx, ny))
                            visited[nx][ny] = True
                            count += 1

            if count != 0:
                counts.append(count)

print(len(counts))

for count in sorted(counts):
    print(count)
