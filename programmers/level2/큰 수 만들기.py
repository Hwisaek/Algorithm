import itertools


def solution(number, k):
    answer = 0

    number = list(map(int, number))

    max_n = 0
    idx = -1
    for i in range(k):
        if number[i] > max_n:
            idx = i
            max_n = number[i]

    number = number[idx:]
    k -= idx

    return answer


print(solution("1924", 2))
print(solution("1231234", 3))
print(solution("4177252841", 4))
