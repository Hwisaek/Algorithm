def solution(d, budget):
    answer = 0
    for n in sorted(d):
        budget -= n
        if budget < 0:
            break
        answer += 1

    return answer
