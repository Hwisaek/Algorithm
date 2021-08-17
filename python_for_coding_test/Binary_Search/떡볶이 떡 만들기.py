def solution(n, m, arr):
    arr.sort()
    height = arr[-1] + 1
    total = 0
    while total < m:
        height -= 1
        l = list(map(lambda x: x - height if x > height else 0, arr))
        total = sum(l)
    return height


print(solution(4, 6, [19, 15, 10, 17]))
