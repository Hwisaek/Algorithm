# https://programmers.co.kr/learn/courses/30/lessons/12924
def solution1(n):
    answer = 1
    for i in range(1, n // 2 + 1):
        sum = 0
        a = i
        while True:
            sum += a
            a += 1
            if sum > n:
                chk = False
                break
            elif sum == n:
                chk = True
                break
        if chk:
            answer += 1
    return answer


def solution(n):
    return len([i for i in range(1, n + 1, 2) if n % i is 0])


print(solution(15))  # 4
print(solution(1))  # 1
print(solution(2))  # 1
