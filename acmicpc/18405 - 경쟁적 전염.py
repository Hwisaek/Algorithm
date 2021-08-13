from collections import deque
from collections import defaultdict

dx = [0, 0, 1, -1]
dy = [1, -1, 0, 0]

n, k = map(int, input().split())  # n: 시험관 크기, k: 바이러스 종류
graph = []  # graph: 시험관
virus = defaultdict(list)  # 바이러스 위치
for i in range(n):
    a = list(map(int, input().split()))
    graph.append(a)
    for j in range(n):
        if a[j] != 0:
            virus[a[j]].append([i, j])

s, x, y = map(int, input().split())  # s: 목표 시간 (x, y): 목표 좌표


def bfs(s):
    global n

    q = deque()
    tmp = deque()
    for i in range(1, k + 1):
        q.extend(virus[i])
        virus[i].clear()
    while (q or tmp) and s > 0:

        while q:
            x, y = q.popleft()

            for i in range(4):
                nx = x + dx[i]
                ny = y + dy[i]
                if 0 <= nx < n and 0 <= ny < n:
                    if graph[nx][ny] == 0:
                        tmp.append([nx, ny])
                        graph[nx][ny] = graph[x][y]

        s -= 1
        if s == 0:
            return
        q = tmp
        tmp = deque()


bfs(s)
print(graph[x - 1][y - 1])
