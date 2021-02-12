n = int(input())

i = 0

k = n
while True:
    n = (n % 10) * 10 + (int(n / 10) + (n % 10)) % 10
    i = i + 1
    if n == k:
        break
print(i)
