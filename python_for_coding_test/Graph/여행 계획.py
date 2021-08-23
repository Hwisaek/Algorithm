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


def solution(v, m, array, plan):
    # 부모 테이블 초기화하기
    # 부모 테이블상에서, 부모를 자기 자신으로 초기화
    parent = [i for i in range(v + 1)]

    for i in range(v):
        for j in range(i, v):
            if array[i][j] == 1:
                union_parent(parent, i, j)

    answer = True
    root = parent[plan[0]]
    for i in plan:
        if root != parent[i]:
            answer = False
            break

    if answer:
        print("YES")
    else:
        print("NO")


solution(5, 4,
         [[0, 1, 0, 1, 1],
          [1, 0, 1, 1, 0],
          [0, 1, 0, 0, 0],
          [1, 1, 0, 0, 0],
          [1, 0, 0, 0, 0]],
         [2, 3, 4, 3])
