n = int(input())

prime = [False] * 5001
prime[0] = False
prime[1] = True

for i in [2, 3, 5]:
    for j in range(i, 5001, i):
        prime[j] = True

idx = 0
while n != 0:
    idx += 1
    if prime[idx]:
        n -= 1

print(idx)
