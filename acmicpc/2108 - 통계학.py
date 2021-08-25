from collections import Counter

N = int(input())
array = []
for _ in range(N):
    array.append(int(input()))

array.sort()
count = dict(Counter(array))

maxN = -5000
countMax = 0
lst = []
for k, v in count.items():
    if v > maxN:
        maxN = v
        count = 1
        lst = [k]
    elif v == maxN:
        count += 1
        lst.append(k)

lst.sort()

result = 0
if len(lst) > 1:
    result = lst[1]
else:
    result = lst[0]

print(round(sum(array) / N))
print(array[N // 2])
print(result)
print(array[-1] - array[0])
