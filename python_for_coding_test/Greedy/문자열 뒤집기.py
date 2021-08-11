def solution(s):
    map = {"0": 0, "1": 0}  # 연속된 0과 1의 개수
    last = -2
    for i in s:
        if last != i:
            map[i] += 1
            last = i
    return min(map.values())


print(solution("0001100"))
print(solution("1000110"))
print(solution("0110110110111111110111111"))
