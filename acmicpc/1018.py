n, m = map(int, input().split())

lst = []

count = []

for i in range(n):
    lst.append(list(input()))

for i in range(n - 7):
    for j in range(m - 7):
        cnt = 0
        for a in range(8):
            if a % 2 == 0:
                if lst[i+a][j] != lst[i][j]:
                    cnt += 1
            else:
                if lst[i+a][j] == lst[i][j]:
                    cnt += 1
            for b in range(8):
                if b % 2 == 0:
                    if lst[i+a][j+b] != lst[i+a][j]:
                        cnt += 1
                else:
                    if lst[i+a][j+b] == lst[i+a][j]:
                        cnt += 1
        count.append(cnt)
print(min(count))