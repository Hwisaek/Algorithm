n = int(input())

lst = []

for _ in range(n):
    x, y = map(int, input().split())
    lst.append([x, y])

for i in range(len(lst)):
    count = 1
    for j in range(len(lst)):
        if lst[i][0] < lst[j][0] and lst[i][1] < lst[j][1]:
            count += 1
    print(count, end=' ')