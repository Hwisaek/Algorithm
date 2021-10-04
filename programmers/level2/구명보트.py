# https://programmers.co.kr/learn/courses/30/lessons/42885
from collections import Counter


def solution(people, limit):
    answer = 0
    count = Counter(people)
    now = 0
    for kg in sorted(count.keys()):
        p = min((limit - now) // kg, count[kg])
        now += kg * p
        count[kg] -= p

    return answer


print(solution([70, 50, 80, 50], 100))
print(solution([70, 80, 50], 100))
