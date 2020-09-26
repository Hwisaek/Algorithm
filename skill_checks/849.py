def conv(number, base):
    T = "0123456789ABCDEF"
    i, j = divmod(number, base)
    if i == 0:
        return T[j]
    else:
        return conv(i, base) + T[j]


def solution(n):
    answer = ''
    n -= 1
    n = conv(n, 3)
    n = str(n)
    
    answer = n
    # print(answer)
    return answer

for i in range(1, 10):
    print('{}: {}'.format(i, solution(i)))