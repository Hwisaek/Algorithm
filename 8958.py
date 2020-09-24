t = int(input())

for i in range(t):
    arr = list(input())
    k = 0
    sum = 0
    for j in arr:
        if j == 'O':
            k = k + 1
            sum = sum + k
        else:
            k = 0
    print(sum)
