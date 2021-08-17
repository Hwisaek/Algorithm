n, target = map(int, input().split())
arr = list(map(int, input().split()))

start = 0
end = max(arr)

result = 0
while start <= end:
    mid = (start + end) // 2
    total = sum(map(lambda x: x - mid if x > mid else 0, arr))
    if total < target:
        end = mid - 1
    else:
        start = mid + 1
        result = mid

print(result)
