def solution(s):
    char = []
    num = 0
    for i in s:
        if i.isalpha():
            char.append(i)
        else:
            num += int(i)
    char.sort()
    return ''.join(char) + str(num)


print(solution('K1KA5CB7'))
print(solution('AJKDLSI412K4JSJ9D'))
