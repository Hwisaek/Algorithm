import sys

input = sys.stdin.readline

N, K, Q = map(int, input().split())

msg = []

read = [False] * N
read[0] = True

for _ in range(K):
    R, P = input().split()
    R = int(R)
    msg.append((R, P))

R, P = msg[(Q - 1)]

if R == 0:
    print("-1")
    sys.exit(0)

for m in msg:
    if m[0] >= R:
        idx = ord(m[1]) - 65
        read[idx] = True

answer = ''

for idx in range(len(read)):
    if not read[idx]:
        answer += chr(idx + 65) + " "

print(answer)
