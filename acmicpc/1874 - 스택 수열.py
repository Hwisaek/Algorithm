import sys

input = sys.stdin.readline

n = int(input())

stack = [0]
count = 1

answer = []
for _ in range(n):
    a = int(input())

    while stack[-1] < a:
        stack.append(count)
        count += 1
        answer.append('+')
    if a == stack.pop():
        answer.append('-')
    else:
        answer = ['NO']
        break

print('\n'.join(answer))
