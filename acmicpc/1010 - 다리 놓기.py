import sys

input = sys.stdin.readline

T = int(input())

fact = [1] * 31

for i in range(2, 31):
    fact[i] = fact[i - 1] * i

for _ in range(T):
    r, n = map(int, input().split())
    total = fact[n] // (fact[r] * fact[n - r])
    print(total)
