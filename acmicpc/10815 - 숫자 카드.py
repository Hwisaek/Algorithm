import sys

input = sys.stdin.readline
n = input()
nlst = set(map(int, input().split()))
m = input()
mlst = list(map(int, input().split()))

answer = ""

for i in mlst:
    if i in nlst:
        answer += "1 "
    else:
        answer += "0 "

print(answer)
