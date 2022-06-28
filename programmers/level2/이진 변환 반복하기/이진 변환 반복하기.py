# https://programmers.co.kr/learn/courses/30/lessons/70129
def solution(s):
    answer = [0, 0]
    s = list(s)
    while len(s) > 1:
        i = 0
        while i < len(s):
            if s[i] == '0':
                s.pop(i)
                answer[1] += 1
            else:
                i += 1
        s = list(bin(len(s))[2:])
        answer[0] += 1
    return answer


print(solution("0111010"), solution("0111010") == [2, 5])
print(solution("110010101001"), solution("110010101001") == [3, 8])
print(solution("01110"), solution("01110") == [3, 3])
print(solution("1111111"), solution("1111111") == [4, 1])
