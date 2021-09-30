# https://programmers.co.kr/learn/courses/30/lessons/76502
from collections import deque


def solution(s):
    answer = 0
    s = deque(s)
    for i in range(len(s)):
        s.rotate()
        if chk(s):
            answer += 1
    return answer


def chk(s):
    d = {")": "(", "]": "[", "}": "{"}
    stack = []
    for c in s:
        stack.append(c)
        if len(stack) > 1:
            if stack[-1] in d and stack[-2] == d[stack[-1]]:
                del stack[-2:]

    return False if stack else True


print(solution("[](){}"))
print(solution("}]()[{"))
print(solution("[)(]"))
print(solution("}}}"))
