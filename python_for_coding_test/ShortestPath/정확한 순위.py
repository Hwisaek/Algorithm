from collections import defaultdict

INF = int(1e9)

n, m = map(int, input().split())

graph = [[INF] * (n + 1) for _ in range(n + 1)]
for a in range(1, n + 1):
    for b in range(1, n + 1):
        if a == b:
            graph[a][b] = 0

for _ in range(m):
    a, b = map(int, input().rstrip().split())
    graph[a][b] = 1

for k in range(1, n + 1):
    for a in range(1, n + 1):
        for b in range(1, n + 1):
            graph[a][b] = min(graph[a][b], graph[a][k] + graph[k][b])

answer = defaultdict(int)
for a in range(1, n + 1):
    for b in range(1, n + 1):
        if graph[a][b] != INF:
            answer[graph[a][b]] += 1

count = 0
for v in answer.values():
    if v == 1:
        count += 1

print(count)
print(answer)
