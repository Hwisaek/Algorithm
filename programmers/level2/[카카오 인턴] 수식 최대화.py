# https://programmers.co.kr/learn/courses/30/lessons/67257
import re
from itertools import permutations


def solution(expression):
    answer = 0
    num = re.split('\+|\-|\*', expression)
    opr = re.split('[0-9]+', expression)[1:-1]
    for order in permutations("+-*", 3):
        nc = num.copy()
        oc = opr.copy()
        for o in order:
            idx = 0
            while idx < len(oc):
                if oc[idx] == o:
                    nc[idx] = str(eval(''.join((nc[idx], o, nc[idx + 1]))))
                    nc.pop(idx + 1)
                    oc.pop(idx)
                    idx -= 1
                idx += 1
        answer = max(answer, abs(int(nc[0])))
    return answer


print(solution("100-200*300-500+20"))
print(solution("50*6-3*2"))
