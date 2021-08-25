a, b = input().split()
answer = 50

for i in range(len(b) - len(a) + 1):
    count = 0
    for j in range(len(a)):
        if b[i + j] != a[j]:
            count += 1
    if count < answer:
        answer = count

print(answer)
