import heapq
import sys

input = sys.stdin.readline
INF = int(1e9)  # 무한을 의미하는 값으로 10억을 설정


def solution(n, edge):
    answer = 0

    start = 1

    # 각 노드에 연결되어 있는 노드에 대한 정보를 담는 리스트를 만들기
    graph = [[] for i in range(n + 1)]
    # 최단 경로 테이블을 모두 무한으로 초기화
    distance = [INF] * (n + 1)

    # 모든 간선 정보를 입력받기
    for a, b, in edge:
        # a번 노드에서 b번 노드로 가는 비용이 c라는 의미
        graph[a].append((b, 1))
        graph[b].append((a, 1))

    def dijkstra(start):
        q = []
        # 시작 노드로 가기 위한 최단 경로는 0으로 설정하여, 큐에 삽입
        heapq.heappush(q, (0, start))
        distance[start] = 0
        while q:  # 큐가 비어있지 않다면
            # 가장 최단 거리가 짧은 노드에 대한 정보 꺼내기
            dist, now = heapq.heappop(q)
            # 현재 노드가 이미 처리된 적이 있는 노드라면 무시
            if distance[now] < dist:
                continue
            # 현재 노드와 연결된 다른 인접한 노드들을 확인
            for i in graph[now]:
                cost = dist + i[1]
                # 현재 노드를 거쳐서, 다른 노드로 이동하는 거리가 더 짧은 경우
                if cost < distance[i[0]]:
                    distance[i[0]] = cost
                    heapq.heappush(q, (cost, i[0]))

    # 다익스트라 알고리즘을 수행
    dijkstra(start)

    max = 0
    # 모든 노드로 가기 위한 최단 거리를 출력
    for i in range(1, n + 1):
        if distance[i] != INF:
            if distance[i] > max:
                max = distance[i]
                answer = 1
            elif distance[i] == max:
                answer += 1

    return answer


print(solution(6, [[3, 6], [4, 3], [3, 2], [1, 3], [1, 2], [2, 4], [5, 2]]))
