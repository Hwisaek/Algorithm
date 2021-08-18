import sys

input = sys.stdin.readline

N = int(input().rstrip())
array = list(map(int, input().rstrip().split()))

start = 0
end = N

idx = -1

while start <= end:
    mid = (start + end) // 2
    if array[mid] == mid:
        idx = mid
        break
    elif array[mid] < mid:
        start = mid + 1
    else:
        end = mid - 1

print(idx if idx > 0 else -1)
