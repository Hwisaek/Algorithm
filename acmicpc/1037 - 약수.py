import sys

input = sys.stdin.readline

n = int(input())

array = sorted(list(map(int, input().split())))
print(array[0] * array[-1])
