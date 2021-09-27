import sys

input = sys.stdin.readline
t = int(input())
for _ in range(t):
    s = input()
    print(int(s[0]) + int(s[2]))
