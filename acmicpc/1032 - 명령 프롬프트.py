n = int(input())

files = []
for _ in range(n):
    files.append(input())

answer = list(files[0])
idx = -1
for i in range(1, len(files)):
    for j in range(len(answer)):
        if answer[j] == "?":
            continue
        if answer[j] != files[i][j]:
            answer[j] = '?'

print(''.join(answer))
