import math


def solution(n, m):
    a = math.gcd(n, m)
    return [a, n * m / a]


print(solution(3, 12))
print(solution(2, 5))
