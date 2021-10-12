# https://programmers.co.kr/learn/courses/30/lessons/12938
from itertools import permutations


def solution(n, s):
    answer = []
    p = list(filter(lambda x: sum(x) == s, permutations([i for i in range(1, s)], n)))
    m = 1
    if len(p) > s or not p:
        return [-1]
    else:
        for arr in p:
            tmp = 1
            for n in arr:
                tmp *= n

            if tmp > m:
                m = tmp
                answer = arr
        return sorted(answer)


print(solution(2, 9), [4, 5])
print(solution(2, 1), [-1])
print(solution(2, 8), [4, 4])
