def solution(numbers):
    answer = set()
    for i in range(len(numbers)):
        for j in range(i):
            answer.add(numbers[i] + numbers[j])
    answer = list(answer)
    answer = sorted(answer)
    return answer