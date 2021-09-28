def solution(left, right):
    answer = 0
    divisors = [1] * (right + 1)
    for n in range(2, right + 1):
        for i in range(n, right + 1, n):
            divisors[i] += 1

    for n in range(left, right + 1):
        if divisors[n] % 2 == 0:
            answer += n
        else:
            answer -= n

    return answer
