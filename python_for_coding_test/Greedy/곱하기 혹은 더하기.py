def solution(s):
    answer = 0

    for i in s:
        i = int(i)
        if answer <= 1 or i <= 1:  # 0이면 곱할 수 없고, 1은 더하는게 더 좋으므로 더하기
            answer += i
        else:
            answer *= i

    return answer


print(solution("02984"))  # 576
print(solution("567"))  # 210
print(solution("0"))  # 0
print(solution("12"))  # 3
print(solution("11111"))  # 5
print(solution("103"))  # 4
print(solution("203"))  # 6
