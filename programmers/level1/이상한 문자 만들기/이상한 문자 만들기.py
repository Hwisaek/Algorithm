def solution(s):
    s = s.split(' ')
    for j in range(len(s)):
        word = list(s[j])
        for i in range(len(word)):
            if i % 2 != 0:
                word[i] = word[i].lower()
            else:
                word[i] = word[i].upper()
        s[j] = ''.join(word)
    return ' '.join(s)


print(solution("try   hello world"))
print(solution("ab ab ab ab"))
