def solution(s, n):
    answer = list(s)
    for i in range(len(answer)):
        if answer[i].islower():
            answer[i] = chr((ord(answer[i]) - ord('a') + n) % 26 + ord('a'))
        elif answer[i].isupper():
            answer[i] = chr((ord(answer[i]) - ord('A') + n) % 26 + ord('A'))
    return ''.join(answer)


print(solution("AB", 1))
print(solution("z", 1))
print(solution("a B z", 4))
print(solution("Z", 10))
print(solution(" aBZ", 20))
print(solution("y X Z", 4))
print(solution(" . h z", 20))
