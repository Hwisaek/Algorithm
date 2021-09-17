import sys

input = sys.stdin.readline

n = input()
a = set(input().split())
m = input()
b = input().split()

answer = ''
for i in b:
    answer += '1\n' if i in a else '0\n'

print(answer)
