n, k = map(int, input().split())

fact = [1] * 11

for i in range(2, 11):
    fact[i] = fact[i - 1] * i

print(fact[n] // (fact[k] * fact[n - k]))
