from bisect import bisect_left, bisect_right

N, x = map(int, input().split())
array = list(map(int, input().split()))

left = bisect_left(array, x)
right = bisect_right(array, x)
answer = right - left if right - left > 0 else -1
print(answer)
