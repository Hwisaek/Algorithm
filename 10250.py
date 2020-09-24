t = int(input())

for i in range(t):
    h, w, n = map(int, input().split())
    
    y = n % h
    if y == 0:
        y = h
    x = n / h
    if x != int(x):
        x = int(x) + 1
    else:
        x = int(x)
    print(y * 100 + x)