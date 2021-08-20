def solution(s):
    answer = 1

    l = len(s)

    for start in range(l):
        for end in range(start + 2, l + 1):
            a = s[start:end]
            if len(a) < answer:
                continue
            if a == a[::-1]:
                answer = max(answer, end - start)
    return answer


print(solution("abcdcba"))
print(solution("abacde"))
