# N: 배열의 크기, M: 숫자가 총 더해지는 횟수, K: 연속해서 더할 수 있는 횟수
N, M, K = map(int, input().split())

# 더해질 숫자
number = list(map(int, input().split()))

number.sort()

first = number[-1]
second = number[-2]

print('number = ', number)
print('first', first, 'second', second)
a = M // (K + 1)
b = M % (K + 1)

print('a, b = ', a, b)

c = first * K + second
result = 0
for i in range(a):
    result += c

result += first * b

print(result)
