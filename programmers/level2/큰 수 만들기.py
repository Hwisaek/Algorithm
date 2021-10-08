def solution(number, k):
    answer = []

    for n in number:
        answer.append(int(n))
        while True:
            if k > 0:
                if len(answer) > 1:
                    if answer[-2] < answer[-1]:
                        answer.pop(-2)
                        k -= 1
                    else:
                        break
                else:
                    break
            else:
                break

    return ''.join(map(str, answer))[:len(number) - k]


print(solution("1924", 2) == "94")
print(solution("1231234", 3) == "3234")
print(solution("4177252841", 4) == "775841")
print(solution("77777", 1), solution("77777", 1) == '7777')
