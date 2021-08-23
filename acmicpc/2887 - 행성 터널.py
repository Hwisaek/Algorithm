import sys

input = sys.stdin.readline


# 두 행성간 터널 비용 구하기
def getCost(a, b):
    return min(abs(a[0] - b[0]), abs(a[1] - b[1]), abs(a[2] - b[2]))


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

# 행성들 좌표
planets = []
for _ in range(n):
    planets.append(list(map(int, input().split())))

# 부모 테이블 초기화하기
# 부모 테이블상에서, 부모를 자기 자신으로 초기화
parent = [i for i in range(n + 1)]

# 모든 간선을 담을 리스트와, 최종 비용을 담을 변수
edges = []
result = 0

for i in range(n):
    for j in range(i + 1, n):
        edges.append((getCost(planets[i], planets[j]), i, j))

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
