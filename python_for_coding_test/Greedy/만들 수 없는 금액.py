from itertools import combinations


def solution(n, arr):  # n: 동전의 개수, arr: 각 동전의 화폐 단위
    chk = [False] * (sum(arr) + 1)
    comb = []

    for i in range(1, n + 1):  # 동전을 1개부터 n개까지 고르는 경우의 수
        comb.extend(list(combinations(arr, i)))

    for c in comb:  # 동전을 고른 경우의 합계를 기록
        chk[sum(c)] = True

    for i in range(1, sum(arr) + 1):  # chk를 1부터 확인하며 가장 작은 False의 index를 리턴
        if not chk[i]:
            return i

    return sum(arr) + 1


print(solution(5, [3, 2, 1, 1, 9]))  # 8
print(solution(3, [3, 5, 7]))  # 1
print(solution(10, [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]))  # 64
print(solution(6, [1, 2, 4, 8, 16, 32]))  # 64


# 예시 답안
def n(n, arr):
    arr.sort()

    target = 1
    for x in arr:
        if target < x:
            break
        target += x
    return target


print(n(5, [3, 2, 1, 1, 9]))  # 8
print(n(3, [3, 5, 7]))  # 1
print(n(10, [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]))  # 64
print(n(6, [1, 2, 4, 8, 16, 32]))  # 64
