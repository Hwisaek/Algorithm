def binary_search(array, target, start, end):
    while start <= end:
        mid = (start + end) // 2
        if array[mid] == target:
            return mid
        elif array[mid] > target:
            end = mid - 1
        else:
            start = mid + 1
    return None


def solution(n, nArr, m, mArr):
    nArr.sort()
    for i in mArr:
        if binary_search(nArr, i, 0, len(nArr)) != None:
            print("yes", end=" ")
        else:
            print("no", end=" ")


solution(5, [8, 3, 7, 9, 2], 3, [5, 7, 9])
