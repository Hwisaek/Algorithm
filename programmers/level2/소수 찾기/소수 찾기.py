from itertools import permutations


def solution(numbers):
    answer = 0
    s = set()
    maxN = int("".join(sorted(list(numbers), reverse=True)))

    prime = [True] * (maxN + 1)  # 에라토스테네스의 체
    prime[0:2] = [False, False]

    # 에라토스 테네스의 체 만들기
    for i in range(2, int(maxN ** 0.5 + 1)):
        if prime[i]:
            for j in range(i * 2, maxN + 1, i):
                prime[j] = False

    # set에 넣어서 중복 숫자 제거
    for i in range(1, len(numbers) + 1):
        s = s.union(set(map(int, map("".join, permutations(numbers, i)))))

    # 소수 개수 판별
    for number in s:
        if prime[number]:
            answer += 1

    return answer


# print(solution("17"))
print(solution("011"))
