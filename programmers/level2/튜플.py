def solution(s):
    answer = []
    s = eval(s.replace('{', '[').replace('}', ']'))
    s.sort(key=lambda x: len(x))

    for array in s:
        for n in array:
            if n not in answer:
                answer.append(n)
    return answer


print(solution("{{2},{2,1},{2,1,3},{2,1,3,4}}"))
print(solution("{{1,2,3},{2,1},{1,2,4,3},{2}}"))
print(solution("{{20,111},{111}}"))
print(solution("{{123}}"))
print(solution("{{4,2,3},{3},{2,3,4,1},{2,3}}"))
