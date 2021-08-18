x = int(input())

l = [0] * 30001

for i in range(1, x + 1):
    if l[i * 5] != 0:
        l[i * 5] = min(l[i * 5], l[i] + 1)
    else:
        l[i * 5] = l[i] + 1
    if l[i * 3] != 0:
        l[i * 3] = min(l[i * 3], l[i] + 1)
    else:
        l[i * 3] = l[i] + 1
    if l[i * 2] != 0:
        l[i * 2] = min(l[i * 2], l[i] + 1)
    else:
        l[i * 2] = l[i] + 1
    if l[i + 1] != 0:
        l[i + 1] = min(l[i + 1], l[i] + 1)
    else:
        l[i + 1] = l[i] + 1

print(l[x])
