# https://programmers.co.kr/learn/courses/30/lessons/12953
import math


def solution(arr):
    answer = 1
    for n in arr:
        answer = n * answer // math.gcd(answer, n)
    return answer
