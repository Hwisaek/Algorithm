import math


def solution(arr):
    answer = 1
    for n in arr:
        answer = n * answer // math.gcd(answer, n)
    return answer
