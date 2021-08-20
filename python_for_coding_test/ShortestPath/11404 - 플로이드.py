import sys

input = sys.stdin.readline

INF = int(1e9)

n = int(input().rstrip())  # 도시의 개수
m = int(input().rstrip())  # 버스의 개수
graph = [[INF] * (n + 1) for _ in range(n + 1)]

for a in range(1, n + 1):
    for b in range(1, n + 1):
        if a == b:
            graph[a][b] = 0

for _ in range(m):
    a, b, c = map(int, input().rstrip().split())
    graph[a][b] = min(c, graph[a][b])

for k in range(1, n + 1):
    for a in range(1, n + 1):
        for b in range(1, n + 1):
            graph[a][b] = min(graph[a][b], graph[a][k] + graph[k][b])

answer = ""
for a in range(1, n + 1):
    for b in range(1, n + 1):
        if graph[a][b] == INF:
            answer += "0 "
        else:
            answer += f"{graph[a][b]} "
    answer += "\n"

print(answer)
