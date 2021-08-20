# 서로소 집합 알고리즘 - 그래프

# O(VM)
# V: 노드의 개수, M: find 혹은 union 연산의 수

# 부모 노드가 연결된 순서를 확인 가능

import sys

input = sys.stdin.readline


# 특정 원소가 속한 집합을 찾기
# x의 루트 노드를 찾아서 반환
def find_parent(parent, x):
    # 루트 노드가 아니라면, 루트 노드를 찾을 때까지 재귀적으로 호출
    if parent[x] != x:
        return find_parent(parent, parent[x])
    return x


# 두 원소가 속한 집합을 합치기
def union_parent(parent, a, b):
    a = find_parent(parent, a)
    b = find_parent(parent, b)
    if a < b:
        parent[b] = a
    else:
        parent[a] = b


# 노드의 개수와 간선(Union 연산)의 개수 입력 받기
v, e = map(int, input().split())

# 부모 테이블 초기화하기
# 부모 테이블상에서, 부모를 자기 자신으로 초기화
parent = [i for i in range(v + 1)]

# Union 연산을 각각 수행
for i in range(e):
    a, b = map(int, input().split())
    union_parent(parent, a, b)

# 각 원소가 속한 집합 출력하기
print('각 원소가 속한 집합: ', end='')
for i in range(1, v + 1):
    print(find_parent(parent, i), end=' ')

print()

# 부모 테이블 내용 출력하기
print('부모 테이블: ', end='')
for i in range(1, v + 1):
    print(parent[i], end=' ')
