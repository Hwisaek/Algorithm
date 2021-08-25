import copy
from collections import defaultdict
from collections import deque


def solution(tickets):
    answer = []

    graph = defaultdict(list)

    for ticket in tickets:
        graph[ticket[0]].append([ticket[1], False])

    for k in graph.keys():
        a = dfs(graph, k)
        if len(a) > len(answer):
            answer = a
        elif len(a) == len(answer) and ''.join(a) < ''.join(answer):
            answer = a

    return answer


def dfs(graph, start, fro):
    graph = copy.deepcopy(graph)
    max = []
    a = []
    for i in graph[start]:
        if not i[1]:
            i[1] = True
            a = dfs(graph, i[0])
            if len(a) > len(max):
                max = a
            elif len(a) == len(max) and ''.join(a) < ''.join(max):
                max = a

    result = [start]
    result.extend(max)

    return result


# print(solution([["ICN", "JFK"], ["HND", "IAD"], ["JFK", "HND"]]))
print(solution([["ICN", "SFO"], ["ICN", "ATL"], ["SFO", "ATL"], ["ATL", "ICN"], ["ATL", "SFO"]]))
