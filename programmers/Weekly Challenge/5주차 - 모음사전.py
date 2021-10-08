# https://programmers.co.kr/learn/courses/30/lessons/84512
from itertools import product


def solution(word):
    vowel = ['A', 'E', 'I', 'O', 'U']
    l = []
    for i in range(1, 6):
        l.extend(map(''.join, product(vowel, repeat=i)))
    l.sort()
    return l.index(word) + 1


print(solution("AAAAE"))
print(solution("AAAE"))
print(solution("I"))
print(solution("EIO"))
