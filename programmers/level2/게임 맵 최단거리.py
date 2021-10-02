# https://programmers.co.kr/learn/courses/30/lessons/1844
from collections import deque

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]
INF = int(1e9)


def solution(maps):
    # 2차원 리스트(그래프 표현)를 만들고, 모든 값을 무한으로 초기화
    n = len(maps)
    m = len(maps[0])
    cost = [[INF] * m for _ in range(n)]

    # 큐(Queue) 구현을 위해 deque 라이브러리 사용
    q = deque([(0, 0)])
    # 현재 노드를 방문 처리
    cost[0][0] = 1

    # 큐가 빌 때까지 반복
    while q:
        x, y = q.popleft()
        for i in range(4):
            nx, ny = x + dx[i], y + dy[i]
            if -1 < nx < n and -1 < ny < m:
                if cost[nx][ny] == INF and maps[nx][ny] == 1:
                    q.append((nx, ny))
                    cost[nx][ny] = cost[x][y] + 1

    return cost[n - 1][m - 1] if cost[n - 1][m - 1] != INF else -1


print(solution([[1, 0, 1, 1, 1], [1, 0, 1, 0, 1], [1, 0, 1, 1, 1], [1, 1, 1, 0, 1], [0, 0, 0, 0, 1]]))
print(solution([[1, 0, 1, 1, 1], [1, 0, 1, 0, 1], [1, 0, 1, 1, 1], [1, 1, 1, 0, 0], [0, 0, 0, 0, 1]]))
