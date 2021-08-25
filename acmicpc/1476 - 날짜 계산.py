# E: 지구, S: 태양, M: 달
# (1 ≤ E ≤ 15, 1 ≤ S ≤ 28, 1 ≤ M ≤ 19)
E, S, M = map(int, input().split())

answer = 1
x, y, z = 1, 1, 1
while True:
    if x == E and y == S and z == M:
        break
    x = x + 1 if x < 15 else 1
    y = y + 1 if y < 28 else 1
    z = z + 1 if z < 19 else 1
    answer += 1

print(answer)
