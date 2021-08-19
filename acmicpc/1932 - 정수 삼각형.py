n = int(input())
array = []
for _ in range(n):
    array.append(list(map(int, input().split())))

for i in range(1, n):
    for j in range(len(array[i])):
        total = []

        if j > 0:
            total.append(array[i][j] + array[i - 1][j - 1])

        if j < len(array[i - 1]):
            total.append(array[i][j] + array[i - 1][j])

        array[i][j] = max(total)

print(max(array[n - 1]))
