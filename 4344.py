c = int(input())

arr = []

for i in range(c):
    n, *arr = map(int, input().split())
    avg = sum(arr) / len(arr)
    num = 0
    for i in arr:
        if i > avg:
            num = num + 1
    print('%.3f%%' % ((num / n) * 100))
