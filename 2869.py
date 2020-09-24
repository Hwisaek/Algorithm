a, b, v = map(int, input().split())

k = (v - b) / (a - b)

if int(k) != k:
    k += 1
print(int(k))