from itertools import combinations


def solution(line):
    points = []

    INF = float('inf')

    max_x, min_x, max_y, min_y = -INF, INF, -INF, INF

    for i, j in combinations(line, 2):
        A, B, E = i
        C, D, F = j

        div = A * D - B * C
        if div == 0: continue

        x = (B * F - E * D) / div
        y = (E * C - A * F) / div

        if x != int(x) or y != int(y):
            continue
        x, y = int(x), int(y)

        points.append((x, y))

        max_x = max(x, max_x)
        min_x = min(x, min_x)
        max_y = max(y, max_y)
        min_y = min(y, min_y)

    w = max_x - min_x + 1
    h = max_y - min_y + 1

    answer = [['.'] * w for _ in range(h)]
    for x_, y_ in points:
        x_cor = x_ - min_x
        y_cor = max_y - y_
        answer[y_cor][x_cor] = '*'

    answer = list(map(''.join, answer))
    return answer


print('\n'.join(solution([[2, -1, 4], [-2, -1, 4], [0, -1, 1], [5, -8, -12], [5, 8, 12]])),
      solution([[2, -1, 4], [-2, -1, 4], [0, -1, 1], [5, -8, -12], [5, 8, 12]]) == ["....*....",
                                                                                    ".........",
                                                                                    ".........",
                                                                                    "*.......*",
                                                                                    ".........",
                                                                                    ".........",
                                                                                    ".........",
                                                                                    ".........",
                                                                                    "*.......*"])
print('\n'.join(solution([[0, 1, -1], [1, 0, -1], [1, 0, 1]])),
      solution([[0, 1, -1], [1, 0, -1], [1, 0, 1]]) == ["*.*"])
print('\n'.join(solution([[1, -1, 0], [2, -1, 0]])), solution([[1, -1, 0], [2, -1, 0]]) == ["*"])
