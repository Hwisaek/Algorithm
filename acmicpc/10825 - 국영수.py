import sys

input = sys.stdin.readline

n = int(input())

students = []

for _ in range(n):
    l = list(input().rstrip().split())
    students.append((-int(l[1]), int(l[2]), -int(l[3]), l[0]))

students.sort()

for s in students:
    print(s[3])
