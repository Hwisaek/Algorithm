n = int(input())

arr = list(map(int, input().split()))

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

print(len(answer))
