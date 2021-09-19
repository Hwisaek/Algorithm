import sys
from collections import deque

input = sys.stdin.readline

m, n = map(int, input().split())

array = [list(input().split()) for _ in range(n)]

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]

q1 = deque()
q2 = deque()

for i in range(n):
    for j in range(m):
        if array[i][j] == '1':
            q1.append((i, j))

q1_flag = True
answer = 0
while q1 or q2:
    if q1_flag:
        x, y = q1.pop()
    else:
        x, y = q2.pop()

    for i in range(4):
        nx, ny = x + dx[i], y + dy[i]

        if -1 < nx < n and -1 < ny < m:
            if array[nx][ny] == '0':
                array[nx][ny] = '1'
                if q1_flag:
                    q2.append((nx, ny))
                else:
                    q1.append((nx, ny))

    if not q1 and not q2:
        break

    if q1_flag and not q1:
        answer += 1
        q1_flag = False
    elif not q1_flag and not q2:
        answer += 1
        q1_flag = True

for i in range(n):
    for j in range(m):
        if array[i][j] == '0':
            answer = -1
            break

print(answer)
