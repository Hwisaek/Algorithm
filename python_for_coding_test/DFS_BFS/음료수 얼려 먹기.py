from collections import deque


def solution(n, m, frame):  # n: 행, m:열, frame: 얼음틀
    answer = 0

    def dfs(x, y):
        if x < 0 or x >= n or y < 0 or y >= m:
            return False
        if frame[x][y] == 0:
            frame[x][y] = 1
            dfs(x + 1, y)
            dfs(x - 1, y)
            dfs(x, y + 1)
            dfs(x, y - 1)
            return True
        return False

    # for i in range(n):
    #     for j in range(m):
    #         if dfs(i, j):
    #             answer += 1

    def bfs(x, y):
        q = deque()
        q.append([x, y])
        frame[x][y] = 1

        while q:
            v = q.popleft()
            if v[0] + 1 < n and frame[v[0] + 1][v[1]] == 0:
                q.append([v[0] + 1, v[1]])
                frame[v[0] + 1][v[1]] = 1
            if v[0] - 1 >= 0 and frame[v[0] - 1][v[1]] == 0:
                q.append([v[0] - 1, v[1]])
                frame[v[0] - 1][v[1]] = 1
            if v[1] + 1 < m and frame[v[0]][v[1] + 1] == 0:
                q.append([v[0], v[1] + 1])
                frame[v[0]][v[1] + 1] = 1
            if v[1] - 1 >= 0 and frame[v[0]][v[1] - 1] == 0:
                q.append([v[0], v[1] - 1])
                frame[v[0]][v[1] - 1] = 1

    for i in range(n):
        for j in range(m):
            if frame[i][j] == 0:
                bfs(i, j)
                answer += 1
    return answer


print(solution(4, 5, [[0, 0, 1, 1, 0],
                      [0, 0, 0, 1, 1],
                      [1, 1, 1, 1, 1],
                      [0, 0, 0, 0, 0]]))
print(solution(15, 14, [[0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0],
                        [1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0],
                        [1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0],
                        [1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0],
                        [1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
                        [1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0],
                        [1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1],
                        [0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
                        [0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1],
                        [0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0],
                        [0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0],
                        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0],
                        [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1],
                        [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1],
                        [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1]]))
