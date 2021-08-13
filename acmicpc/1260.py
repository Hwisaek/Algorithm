from collections import defaultdict
from collections import deque


def dfs(n):
    global visited
    global tree
    if visited[n]:
        return
    print(n, end=' ')
    visited[n] = True

    if tree[n]:
        for m in tree[n]:
            dfs(m)


def bfs(n):
    global visited
    global tree
    global q

    q.append(n)
    visited[n] = True
    while q:
        next = q.popleft()
        print(next, end=' ')
        for i in tree[next]:
            if visited[i]:
                continue
            q.append(i)
            visited[i] = True


n, m, v = map(int, input().split())  # n: 정점의 개수, m: 간선의 개수, v: 탐색을 시작할 정점 번호

tree = defaultdict(list)
visited = defaultdict(bool)

for i in range(m):
    a, b = map(int, input().split())
    tree[a].append(b)

for key in tree:
    tree[key].sort()
    visited[key] = False

dfs(v)

print()
q = deque()
for key in visited:
    visited[key] = False

bfs(v)
