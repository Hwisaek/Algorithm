n = int(input())

array = [1e7] * (n + 1)
array[1] = 0
i = 1
while i < n:
    if i + 1 < n + 1:
        array[i + 1] = min(array[i + 1], array[i] + 1)
    if i * 2 < n + 1:
        array[i * 2] = min(array[i * 2], array[i] + 1)
    if i * 3 < n + 1:
        array[i * 3] = min(array[i * 3], array[i] + 1)
    i += 1

print(array[n])
