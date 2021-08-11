def solution(s):
    answer = 0
    x = ord(s[0]) - ord('a')
    y = int(s[1]) - 1
    d = [[1, 2], [1, -2], [-1, 2], [-1, -2], [2, 1], [2, -1], [-2, 1], [-2, -1]]

    for dx, dy in d:
        if 0 <= x + dx < 8 and 0 <= y + dy < 8:
            answer += 1

    return answer


print(solution("a1"))
