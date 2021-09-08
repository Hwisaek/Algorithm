import sys

input = sys.stdin.readline

N = int(input())
members = []
for _ in range(N):
    age, name = input().split()
    members.append((int(age), name))

members = sorted(enumerate(members), key=lambda x: (x[1][0], x[0]))

for idx, member in members:
    sys.stdout.write(f'{member[0]} {member[1]}\n')
