n = int(input())

arr = list(map(int, input().split()))

score = []

for i in arr:
    score.append(i / max(arr) * 100)

print(sum(score) / len(score))
