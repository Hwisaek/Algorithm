import re


def solution(n, k):
    answer = 0

    transited = trans(n, k)
    transited = re.sub('0+', '0', transited)
    if transited[0] == '0':
        transited = transited[1:]
    if transited[-1] == '0':
        transited = transited[:-1]
    nums = list(map(int, transited.rsplit('0')))

    for i in nums:
        if isPrime(i):
            answer += 1

    return answer


def isPrime(n):
    for i in range(2, int(n ** 0.5 + 1)):
        if n % i == 0:
            return False
    return n > 1


# return str
def trans(n, q):
    rev_base = ''

    while n > 0:
        n, mod = divmod(n, q)
        rev_base += str(mod)

    return rev_base[::-1]


print(solution(437674, 3))
print(solution(110011, 10))
print(solution(1000000, 10))
print(solution(1000000, 3))
print(solution(1, 3))
print(solution(1, 10))
print(solution(524287, 2))
