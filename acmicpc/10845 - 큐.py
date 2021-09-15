import sys
from collections import deque

input = sys.stdin.readline

n = int(input())

q = deque()

answer = []
for _ in range(n):
    opr = input().split()
    if opr[0] == 'push':
        q.append(opr[1])
    elif opr[0] == 'pop':
        answer.append(q.popleft() if q else '-1')
    elif opr[0] == 'size':
        answer.append(str(len(q)))
    elif opr[0] == 'empty':
        answer.append('0' if q else '1')
    elif opr[0] == 'front':
        answer.append(q[0] if q else '-1')
    elif opr[0] == 'back':
        answer.append(q[-1] if q else '-1')

print('\n'.join(answer))
