import copy
from collections import defaultdict
import heapq
from collections import deque

N = int(input())

time = 0  # 움직인 시간


# BFS 함수 정의
def bfs(graph, start, target, size):
    global N
    global time, x, y

    map = copy.deepcopy(graph)
    map[start[0]][start[1]] = 0

    # 큐(Queue) 구현을 위해 deque 라이브러리 사용
    queue = deque([start])

    # 현재 노드를 방문 처리
    visited = [[False] * N for _ in range(N)]
    visited[start[0]][start[1]] = True

    # 큐가 빌 때까지 반복
    while queue:
        # 큐에서 하나의 원소를 뽑아 출력
        a, b = queue.popleft()

        # 해당 원소와 연결된, 아직 방문하지 않은 원소들을 큐에 삽입
        for i in range(4):
            nx = a + dx[i]
            ny = b + dy[i]
            if 0 <= nx < N and 0 <= ny < N:
                if not visited[nx][ny] and graph[nx][ny] <= size:
                    queue.append((nx, ny))
                    visited[nx][ny] = True
                    map[nx][ny] = map[a][b] + 1
                    if 0 < graph[nx][ny] < size:  # 해당 물고기를 먹을 수 있으면 True 반환
                        time += map[nx][ny]
                        graph[nx][ny] = 0
                        x, y = nx, ny
                        return True
    return False  # 해당 물고기를 먹을 수 없으면 False


array = []

size = 2  # 아기 상어의 크기
x, y = -1, -1  # 아기 상어의 좌표

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]

fish = []  # 물고기
for i in range(N):
    data = list(map(int, input().split()))
    array.append(data)
    for j in range(N):
        if 0 < data[j] < 9:  # 물고기면 추가
            fish.append((data[j], i, j))
        if data[j] == 9:  # 아기 상어의 위치
            x = i
            y = j
fish.sort(reverse=True)

count = 0
while fish:
    if fish[-1][0] < size:  # 물고기를 잡아 먹을 수 있으면 해당 물고기 탐색
        f = fish.pop()
        if bfs(array, (x, y), f, size):
            count += 1
            if count == size:
                size += 1
                count = 0
    else:  # 더이상 먹을 수 없으면 종료
        break

print(time)
