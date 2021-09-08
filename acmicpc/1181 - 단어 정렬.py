import sys

input = sys.stdin.readline

N = int(input())
words = set()
for _ in range(N):
    words.add(input())

words = list(words)
words.sort(key=lambda x: (len(x), x))

print(''.join(words))
