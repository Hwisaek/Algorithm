import copy
import sys
from collections import defaultdict
from collections import deque

input = sys.stdin.readline


def dfs(tree, visited, n):
    if visited[n]:
        return

    visited[n] = True
    print(n, end=' ')

    if tree[n]:
        for m in tree[n]:
            if not visited[m]:
                dfs(tree, visited, m)


def bfs(tree, visited, n):
    q = deque()
    q.append(n)
    visited[n] = True
    while q:
        next = q.popleft()
        print(next, end=' ')
        if tree[n]:
            for i in tree[next]:
                if not visited[i]:
                    q.append(i)
                    visited[i] = True


n, m, v = map(int, input().split())  # n: 정점의 개수, m: 간선의 개수, v: 탐색을 시작할 정점 번호

tree = defaultdict(set)
visited = [False] * (n + 1)

for i in range(m):
    a, b = map(int, input().split())
    tree[a].add(b)
    tree[b].add(a)

for key in tree:
    tree[key] = sorted(list(tree[key]))

dfs(copy.deepcopy(tree), copy.deepcopy(visited), v)
print()
bfs(copy.deepcopy(tree), copy.deepcopy(visited), v)
