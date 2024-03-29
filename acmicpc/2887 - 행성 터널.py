import sys

input = sys.stdin.readline


# 특정 원소가 속한 집합을 찾기
def find_parent(parent, x):
    # 루트 노드가 아니라면, 루트 노드를 찾을 때까지 재귀적으로 호출
    if parent[x] != x:
        parent[x] = find_parent(parent, parent[x])
    return parent[x]


# 두 원소가 속한 집합을 합치기
def union_parent(parent, a, b):
    a = find_parent(parent, a)
    b = find_parent(parent, b)
    if a < b:
        parent[b] = a
    else:
        parent[a] = b


n = int(input())

# 부모 테이블 초기화하기
# 부모 테이블상에서, 부모를 자기 자신으로 초기화
parent = [i for i in range(n + 1)]

# 모든 간선을 담을 리스트와, 최종 비용을 담을 변수
edges = []
result = 0

# 행성들 좌표
X = []
Y = []
Z = []
for i in range(n):
    x, y, z = list(map(int, input().split()))
    X.append((x, i))
    Y.append((y, i))
    Z.append((z, i))
X.sort()
Y.sort()
Z.sort()

for i in range(n - 1):
    edges.append((X[i + 1][0] - X[i][0], X[i][1], X[i + 1][1]))
    edges.append((Y[i + 1][0] - Y[i][0], Y[i][1], Y[i + 1][1]))
    edges.append((Z[i + 1][0] - Z[i][0], Z[i][1], Z[i + 1][1]))

# 간선을 비용순으로 정렬
edges.sort()

# 간선을 하나씩 확인하며
for edge in edges:
    cost, a, b = edge
    # 사이클이 발생하지 않는 경우에만 집합에 포함
    if find_parent(parent, a) != find_parent(parent, b):
        union_parent(parent, a, b)
        result += cost

print(result)
