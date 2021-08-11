def solution(n, plan):  # n: 공간의 크기, plan: 이동 계획서
    x, y = 1, 1
    d = {'L': [0, -1], 'R': [0, 1], 'U': [-1, 0], 'D': [1, 0]}  # 방향 맵핑
    for p in plan:
        x += d[p][0]
        y += d[p][1]
        if not 1 <= x <= n:
            x -= d[p][0]
        if not 1 <= y <= n:
            y -= d[p][1]

    return f'{x} {y}'


print(solution(5, ['R', 'R', 'R', 'U', 'D', 'D']))
