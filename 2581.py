m = int(input())
n = int(input())

arr = []

for i in range(m, n + 1):
    arr.append(i)

answer = []

for i in arr:
    if i == 2:
        answer.append(i)
    else:
        for j in range(2, i):
            if i % j == 0:
                break
            else:
                if j >= int(i / 2):
                    answer.append(i)
                    break

if len(answer) > 0:
    print(sum(answer))
    print(min(answer))
else:
    print(-1)