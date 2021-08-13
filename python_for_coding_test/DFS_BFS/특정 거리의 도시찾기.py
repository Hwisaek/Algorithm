import sys
from collections import deque
from collections import defaultdict

n, m, k, x = map(int, sys.stdin.readline().rstrip().split())  # n:도시의 개수, m: 도로의 개수, k:거리 정보, x:출발 도시
cities = {x: -1 for x in range(n + 1)}
graph = defaultdict(list)
for _ in range(m):
    a, b = map(int, sys.stdin.readline().rstrip().split())
    graph[a].append(b)

answer = []

q = deque()

q.append(x)
cities[x] = 0
while q:
    x = q.popleft()

    for i in graph[x]:
        if cities[i] != -1:  # 방문한 노드면 생략
            continue
        q.append(i)
        cities[i] = cities[x] + 1
        if cities[i] == k:
            answer.append(i)

answer.sort()

if answer:
    for i in answer:
        print(i)
else:
    print(-1)
