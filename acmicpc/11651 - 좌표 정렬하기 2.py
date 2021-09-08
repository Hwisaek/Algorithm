import sys

input = sys.stdin.readline

N = int(input())
points = []
for _ in range(N):
    x, y = map(int, input().split())
    points.append((x, y))

points.sort(key=lambda x: (x[1], x[0]))

for x, y in points:
    sys.stdout.write(f'{x} {y}\n')
