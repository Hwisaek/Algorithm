from collections import deque
import sys

input = sys.stdin.readline
# 노드의 개수와 간선의 개수를 입력 받기
v = int(input())

# 모든 노드에 대한 진입차수는 0으로 초기화
indegree = [0] * (v + 1)
# 각 노드에 연결된 간선 정보를 담기 위한 연결 리스트 초기화
graph = [[] for i in range(v + 1)]

time = [0] * (v + 1)
# 방향 그래프의 모든 간선 정보를 입력 받기
for i in range(1, v + 1):
    data = list(map(int, input().split()))
    time[i] = data[0]

    # 방향 그래프의 모든 간선 정보를 입력 받기
    for j in data[1:-1]:
        graph[j].append(i)  # 정점 A에서 B로 이동 가능
        # 진입 차수를 1 증가
        indegree[i] += 1


# 위상 정렬 함수
def topology_sort():
    pre = [0] * (v + 1)
    q = deque()  # 큐 기능을 위한 deque 라이브러리 사용

    # 처음 시작할 때는 진입차수가 0인 노드를 큐에 삽입
    for i in range(1, v + 1):
        if indegree[i] == 0:
            q.append(i)

    # 큐가 빌 때까지 반복
    while q:
        # 큐에서 원소 꺼내기
        now = q.popleft()
        time[now] += pre[now]
        # 해당 원소와 연결된 노드들의 진입차수에서 1 빼기
        for i in graph[now]:
            indegree[i] -= 1
            pre[i] = max(pre[i], time[now])
            # 새롭게 진입차수가 0이 되는 노드를 큐에 삽입
            if indegree[i] == 0:
                q.append(i)

    for _ in time[1:]:
        print(_)


topology_sort()
