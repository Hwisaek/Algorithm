# https://programmers.co.kr/learn/courses/30/lessons/81302
from collections import deque

dx = [0, 0, 1, -1]
dy = [1, -1, 0, 0]


def solution(places):
    answer = []

    for place in places:
        now = 1
        for i in range(5):
            for j in range(5):
                if place[i][j] == 'P':
                    if isDistanced(place, (i, j, 2)):
                        now = 0
                        break
            if now == 0:
                break
        answer.append(now)

    return answer


# BFS 함수 정의
def isDistanced(graph, start):
    visited = [[False] * 5 for _ in range(5)]

    # 큐(Queue) 구현을 위해 deque 라이브러리 사용
    queue = deque()
    queue.append(start)

    visited[start[0]][start[1]] = True

    # 큐가 빌 때까지 반복
    while queue:
        # 큐에서 하나의 원소를 뽑아 출력
        x, y, time = queue.popleft()
        # 해당 원소와 연결된, 아직 방문하지 않은 원소들을 큐에 삽입
        for i in range(4):
            nx, ny = x + dx[i], y + dy[i]
            if -1 < nx < 5 and -1 < ny < 5:
                if not visited[nx][ny]:
                    if graph[nx][ny] == 'P':
                        return True
                    if graph[nx][ny] != 'X':
                        visited[nx][ny] = True
                        if time > 1:
                            queue.append((nx, ny, time - 1))
    return False


print(solution([["POOOP",
                 "OXXOX",
                 "OPXPX",
                 "OOXOX",
                 "POXXP"],
                ["POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"],
                ["PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"],
                ["OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"],
                ["PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"]]),
      solution([["POOOP", "OXXOX", "OPXPX", "OOXOX", "POXXP"],
                ["POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"],
                ["PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"],
                ["OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"],
                ["PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"]]) == [1, 0, 1, 1, 1])
