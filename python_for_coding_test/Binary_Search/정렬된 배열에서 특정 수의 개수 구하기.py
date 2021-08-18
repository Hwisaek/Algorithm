N, x = map(int, input().split())
array = list(map(int, input().split()))

a = 0
b = 0

start = 0
end = N

while start < end:
    mid = (start + end) // 2
    if array[mid] >= x:
        end = mid - 1
    else:
        start = mid + 1
        a = start

start = 0
end = N
while start < end:
    mid = (start + end) // 2
    if array[mid] <= x:
        start = mid + 1
        b = start
    else:
        end = mid - 1

answer = b - a if b - a > 0 else -1
print(answer)
