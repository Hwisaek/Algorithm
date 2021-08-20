INF = int(999)


def solution(n, results):
    graph = [[INF] * (n + 1) for _ in range(n + 1)]
    for a in range(1, n + 1):
        for b in range(1, n + 1):
            if a == b:
                graph[a][b] = 0

    for a, b in results:
        graph[a][b] = 1

    # 점화식에 따라 플로이드 워셜 알고리즘을 수행
    for k in range(1, n + 1):
        for a in range(1, n + 1):
            for b in range(1, n + 1):
                graph[a][b] = min(graph[a][b], graph[a][k] + graph[k][b])

    result = 0
    # 각 학생을 번호에 따라 한 명씩 확인하며 도달 가능한지 체크
    for i in range(1, n + 1):
        count = 0
        for j in range(1, n + 1):
            if graph[i][j] != INF or graph[j][i] != INF:
                count += 1
        if count == n:
            result += 1

    return result


print(solution(5, [[4, 3], [4, 2], [3, 2], [1, 2], [2, 5]]))
