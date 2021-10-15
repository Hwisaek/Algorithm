n, k = map(int, input().split())
s = list(map(int, input().split()))

for _ in range(k):
    a, b = map(int, input().split())
    avg = sum(s[a - 1:b]) / (b - a + 1)
    print(f'{avg:.2f}')
