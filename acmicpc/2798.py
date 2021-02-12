n, m = map(int, input().split())

card = list(map(int, input().split()))


lst = []

answer = 0

for i in range(len(card) - 2):
    for j in range(i + 1, len(card) - 1):
        for k in range(j + 1, len(card)):
            lst = [card[i], card[j], card[k]]
            if m >= sum(lst) and m - sum(lst) < m - answer:
                answer = sum(lst)
print(answer)