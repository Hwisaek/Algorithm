import sys

input = sys.stdin.readline

n = int(input())
a = set(map(int, input().split()))
m = int(input())
b = list(map(int, input().split()))

answer = ''
for i in b:
    if i in a:
        answer += '1\n'
    else:
        answer += '0\n'

print(answer)
