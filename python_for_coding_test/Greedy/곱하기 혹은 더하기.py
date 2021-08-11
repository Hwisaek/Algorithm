def solution(s):
    answer = 0
    for i in s:
        i = int(i)
        if answer <= 1 or i <= 1:  # 0이면 곱할 수 없고, 1은 더하는게 더 좋으므로 더하기
            answer += i
        else:
            answer *= i

    return answer


print(solution("02984"))
print(solution("567"))
print(solution("0"))
print(solution("12"))
print(solution("11111"))
