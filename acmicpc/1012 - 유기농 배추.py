from collections import deque
import sys

input = sys.stdin.readline

T = int(input())
for _ in range(T):
    M, N, K = map(int, input().split())
    array = [[0] * M for _ in range(N)]
    cabbage = []
    for _ in range(K):
        y, x = map(int, input().split())
        cabbage.append([x, y])
        array[x][y] = 1

    visited = [[False] * M for _ in range(N)]

    dx = [0, 0, 1, -1]
    dy = [1, -1, 0, 0]


    def bfs(array, start, visited):
        if array[start[0]][start[1]] == 1 and not visited[start[0]][start[1]]:
            queue = deque([start])

            visited[start[0]][start[1]] = True

            while queue:
                x, y = queue.popleft()
                for i in range(4):
                    nx = x + dx[i]
                    ny = y + dy[i]
                    if 0 <= nx < N and 0 <= ny < M:
                        if array[nx][ny] == 1 and not visited[nx][ny]:
                            queue.append([nx, ny])
                            visited[nx][ny] = True
            return True
        else:
            return False


    answer = 0
    for start in cabbage:
        if bfs(array, start, visited):
            answer += 1
    print(answer)
