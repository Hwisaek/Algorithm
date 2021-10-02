# https://programmers.co.kr/learn/courses/30/lessons/12911
def solution(n):
    c = bin(n).count('1')
    while True:
        n += 1
        a = bin(n).count('1')
        if a == c:
            return n


print(solution(78))  # 83
print(solution(15))  # 23
