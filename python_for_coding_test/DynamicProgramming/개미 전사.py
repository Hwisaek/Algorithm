N = int(input())
array = list(map(int, input().split()))

for i in range(2, N):
    array[i] += array[i - 2]

print(max(array[N - 1], array[N - 2]))
