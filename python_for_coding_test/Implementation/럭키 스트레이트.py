def solution(n):
    n = list(map(int, n))
    length = len(n)
    if sum(n[:length // 2]) == sum(n[length // 2:]):
        return "LUCKY"
    return "READY"

n = input()
print(solution(n))
