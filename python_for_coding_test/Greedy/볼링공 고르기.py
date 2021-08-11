from itertools import combinations
from collections import Counter
import time


def solution(n, m, k):  # n:볼링공의 개수, m: 볼링공 최대 무게, k: 각 볼링공의 무게
    start = time.time()
    answer = 0

    comb = list(combinations(k, 2))

    answer += len(comb)

    for a, b in comb:
        if a == b:
            answer -= 1

    end = time.time()
    print("수행시간: ", end - start)
    return answer


print(solution(5, 3, [1, 3, 2, 3, 2]))  # 8
print(solution(8, 5, [1, 5, 4, 3, 2, 4, 5, 2]))  # 25


def solution2(n, m, k):  # n:볼링공의 개수, m: 볼링공 최대 무게, k: 각 볼링공의 무게
    start = time.time()
    answer = 0

    count = Counter(k)  # 리스트 내의 원소의 개수를 자동으로 계산
    values = list(count.values())
    # 볼링공을 무게 별로 나눠서 조합의 수 계산
    ## 볼링공 무게별 개수: {1:1, 2:2, 3:2}
    ## 무게 1부터 고르기 시작.
    ## 무게 1을 고를 때 경우의 수: 2 + 2 = 4
    ## 무게 2를 고를 때 경우의 수: 2 * 2 = 4
    ## 총 8

    ## 볼링공 무게별 개수: {1:1, 2:2, 3:1, 4:2, 5:2}
    ## 무게 1부터 선택 시작
    ## 무게 1: 1 * (2 + 1 + 2 + 2) = 7
    ## 무게 2: 2 * (1 + 2 + 2) = 10
    ## 무게 3: 1 * (2 + 2) = 4
    ## 무게 4: 2 * 2= 4
    ## 총 25
    for i in range(len(values)):
        answer += values[i] * sum(values[i + 1:])

    end = time.time()
    print("수행시간: ", end - start)
    return answer


print(solution2(5, 3, [1, 3, 2, 3, 2]))  # 8
print(solution2(8, 5, [1, 5, 4, 3, 2, 4, 5, 2]))  # 25
