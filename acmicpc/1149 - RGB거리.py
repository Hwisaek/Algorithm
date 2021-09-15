import sys

input = sys.stdin.readline

n = int(input())

rgb = [(0, 0, 0)]
for i in range(1, n + 1):
    r, g, b = map(int, input().split())
    nr = r + min(rgb[i - 1][1], rgb[i - 1][2])
    ng = g + min(rgb[i - 1][2], rgb[i - 1][0])
    nb = b + min(rgb[i - 1][1], rgb[i - 1][0])
    rgb.append((nr, ng, nb))

print(min(rgb[n]))
