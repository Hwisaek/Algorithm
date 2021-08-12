from collections import deque

# 동 남 북 서
dx = [1, 0, 0, -1]
dy = [0, 1, -1, 0]


def solution(n, m, map):  # n:행, m:열, map:미로 정보

    def bfs(x, y):
        q = deque()

        q.append([x, y])
        while q:
            x, y = q.popleft()
            for i in range(4):
                nx = x + dx[i]
                ny = y + dy[i]
                # 범위 벗어나면 스킵
                if not 0 <= nx < n or not 0 <= ny < m:
                    continue
                # 벽이면 스킵
                if map[nx][ny] == 0:
                    continue
                if (nx != 0 or ny != 0) and map[nx][ny] == 1:
                    map[nx][ny] = map[x][y] + 1
                    q.append([nx, ny])

        return map[n - 1][m - 1]

    return bfs(0, 0)


print(solution(5, 6, [[1, 0, 1, 0, 1, 0],
                      [1, 1, 1, 1, 1, 1],
                      [0, 0, 0, 0, 0, 1],
                      [1, 1, 1, 1, 1, 1],
                      [1, 1, 1, 1, 1, 1]]))
