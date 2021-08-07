def solution(brown, yellow):
    total = brown + yellow

    for h in range(3, int(total ** 0.5) + 1):  # w: 가로, h: 세로 (최소: 3)
        w = total // h
        if h * w == total and (h - 2) * (w - 2) == yellow:
            break

    if h > w:  # 세로가 가로보다 길면 서로 변경
        h, w = w, h

    return [w, h]


print(solution(10, 2))  # [4, 3]
print(solution(8, 1))  # [3,3]
print(solution(24, 24))  # [8, 6]
print(solution(14, 4))  # [6, 3]
print(solution(50, 22))  # [24, 3]
