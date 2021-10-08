def solution(number, k):
    answer = []

    for n in number:
        while k > 0 and len(answer) > 0 and answer[-1] < n:
            answer.pop()
            k -= 1
        answer.append(n)

    return ''.join(answer[:len(number) - k])


print(solution("1924", 2), solution("1924", 2) == "94")
print(solution("1231234", 3), solution("1231234", 3) == "3234")
print(solution("4177252841", 4), solution("4177252841", 4) == "775841")
print(solution("77777", 1), solution("77777", 1) == '7777')
