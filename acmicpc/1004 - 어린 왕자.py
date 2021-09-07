import sys

input = sys.stdin.readline

T = int(input())

for _ in range(T):
    answer = 0
    x1, y1, x2, y2 = map(int, input().split())
    n = int(input())
    planet = []
    for _ in range(n):
        cx, cy, r = list(map(int, input().split()))
        d1 = ((x1 - cx) ** 2 + (y1 - cy) ** 2) ** 0.5
        d2 = ((x2 - cx) ** 2 + (y2 - cy) ** 2) ** 0.5
        if (d1 < r < d2) or (d1 > r > d2):
            answer += 1
    print(answer)
