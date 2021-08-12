def solution(s):
    l = len(s)
    answer = l
    for d in range(1, l // 2 + 1):  # d: 문자열 반복 길이
        code = []  # 압축된 문자열
        last = s[0:d]  # 마지막 문자열
        count = 1  # 문자열 반복 횟수
        for i in range(d, l, d):  # d개씩 건너뜀
            if last != s[i:i + d]:  # 마지막 문자와 다르면 실행
                if count != 1:  # 1번만 반복되면 숫자 생략
                    code.append(str(count))
                code.append(last)
                count = 1
                last = s[i:i + d]  # 마지막 문자열 값 갱신
            else:  # 문자열이 반복되면 count 추가
                count += 1

        # 마지막 반복된 문자열 추가
        if count != 1:
            code.append(str(count))
        code.append(last)

        # 최소길이인 경우 값 갱신
        answer = min(answer, len(''.join(code)))

    return answer


print(solution("aabbaccc"))
print(solution("ababcdcdababcdcd"))
print(solution("abcabcdede"))
print(solution("abcabcabcabcdededededede"))
print(solution("xababcdcdababcdcd"))
