import sys

input = sys.stdin.readline


# 특정 원소가 속한 집합을 찾기
# x의 루트 노드를 찾아서 반환
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


answer = 0

# 탑승구의 수
g = int(input())

# 비행기의 수
p = int(input())

# 부모 테이블 초기화하기
# 부모 테이블상에서, 부모를 자기 자신으로 초기화
parent = [i for i in range(g + 1)]

for _ in range(p):
    root = find_parent(parent, int(input()))
    if root == 0:
        break
    else:
        union_parent(parent, root, root - 1)
        answer += 1

print(answer)
