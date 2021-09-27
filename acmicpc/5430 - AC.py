import sys
from collections import deque

input = sys.stdin.readline
t = int(input().rstrip())

for _ in range(t):
    p = input().rstrip().replace('RR', '')
    n = input().rstrip()
    array = input().rstrip()[1: -1].split(',')
    if array[0] != '':
        array = deque(map(int, array))
    else:
        array.pop()

    error = False
    order = 1
    for opr in p:
        if opr == 'R':
            order *= -1
        elif opr == 'D':
            if not array:
                error = True
                print('error')
                break
            elif order == 1:
                array.popleft()
            else:
                array.pop()

    if not error:
        print('[', end='')
        answer = ''
        for i in range(len(array))[::1 * order]:
            answer += str(array[i]) + ','
        print(answer[:-1], end='')
        print(']')
