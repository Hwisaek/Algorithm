import sys

sys.setrecursionlimit(10 ** 10)

n, m = map(int, sys.stdin.readline().rstrip().split())
graph = []  # 0: 빈칸, 1: 벽, 2: 바이러스
tmp = [[0] * m for _ in range(n)]
for _ in range(n):
    graph.append(list(map(int, sys.stdin.readline().rstrip().split())))

dx = [0, 0, 1, -1]
dy = [1, -1, 0, 0]

result = 0


def dfsV(x, y):
    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        if 0 <= nx < n and 0 <= ny < m:
            tmp[nx][ny] = 2
            dfsV(nx, ny)


def safeSize():
    size = 0
    for i in range(n):
        for j in range(m):
            if tmp[i][j] == 0:
                size += 1
    return size


def dfs(count):
    global result

    if count == 3:
        for i in range(n):
            for j in range(m):
                tmp[i][j] = graph[i][j]

        for i in range(n):
            for j in range(m):
                if tmp[i][j] == 2:
                    dfsV(i, j)

        result = max(result, safeSize())
        return

    for i in range(n):
        for j in range(m):
            if graph[i][j] == 0:
                graph[i][j] = 1
                count += 1
                dfs(count)
                graph[i][j] = 0
                count -= 1


dfs(0)
print(result)
