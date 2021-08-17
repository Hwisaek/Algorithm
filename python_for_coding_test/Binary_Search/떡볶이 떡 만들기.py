def binary_search(array, target, start, end):
    while start <= end:
        mid = (start + end) // 2
        total = sum(list(map(lambda x: x - mid if x > mid else 0, array)))
        if total == target:
            return mid
        elif total > target:
            start = mid + 1
        else:
            end = mid - 1
    return None


def solution(n, m, arr):
    arr.sort()
    return binary_search(arr, m, 0, arr[-1] + 1)


print(solution(4, 6, [19, 15, 10, 17]))
