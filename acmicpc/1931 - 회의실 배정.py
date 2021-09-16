import sys

input = sys.stdin.readline

n = int(input())

array = []
for i in range(n):
    start, end = map(int, input().split())
    array.append((start, end))

array.sort(key=lambda x: (x[1], x[0]))

count = 1
end = array[0][1]
for i in range(1, n):
    if array[i][0] >= end:
        end = array[i][1]
        count += 1
        
print(count)
