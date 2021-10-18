n = int(input())

array = list(map(int, input().split()))
print(sorted(array)[n // 2])
