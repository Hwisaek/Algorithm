import sys

input = sys.stdin.readline

n = int(input())
a = sorted(list(map(int, input().split())))
m = int(input())
b = list(map(int, input().split()))


def binary_search(array, target, start, end):
    while start <= end:
        mid = (start + end) // 2
        if array[mid] == target:
            return True
        elif array[mid] > target:
            end = mid - 1
        else:
            start = mid + 1
    return False


answer = ''
for i in b:
    if binary_search(a, i, 0, n - 1):
        answer += '1\n'
    else:
        answer += '0\n'

print(answer)
