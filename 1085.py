x, y, w, h = map(int, input().split())

answer = []

if x < w / 2:
    answer.append(x)
else:
    answer.append(w - x)

if y < h / 2:
    answer.append(y)
else:
    answer.append(h - y)

print(min(answer))