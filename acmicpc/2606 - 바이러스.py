from collections import defaultdict
from collections import deque

n = int(input())  # 컴퓨터의 수
d = int(input())  # 간선의 수
graph = defaultdict(list)  # 그래프 연결 정보
visited = [False] * (n + 1)
for _ in range(d):
    a, b = map(int, input().split())
    graph[a].append(b)
    graph[b].append(a)


def bfs(x):
    count = 0
    q = deque()

    q.append(x)
    visited[x] = True

    while q:
        x = q.popleft()

        for i in graph[x]:
            if visited[i]:
                continue
            q.append(i)
            visited[i] = True
            count += 1

    return count


print(bfs(1))
