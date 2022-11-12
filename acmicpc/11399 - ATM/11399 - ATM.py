import sys

input = sys.stdin.readline

n = int(input())

array = list(map(int, input().split()))
array.sort()

for i in range(1, n):
    array[i] += array[i - 1]

print(sum(array))
