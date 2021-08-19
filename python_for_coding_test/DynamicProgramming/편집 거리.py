a = input()
b = input()

array = [[0] * (len(a) + 1) for _ in range(len(b) + 1)]

array[0] = [_ for _ in range(len(a) + 1)]
for _ in range(len(b) + 1):
    array[_][0] = _

for i in range(1, len(b) + 1):
    for j in range(1, len(a) + 1):
        array[i][j] = min(array[i - 1][j], array[i][j - 1], array[i - 1][j - 1])
        if a[j - 1] != b[i - 1]:
            array[i][j] += + 1

print(array[i][j])
