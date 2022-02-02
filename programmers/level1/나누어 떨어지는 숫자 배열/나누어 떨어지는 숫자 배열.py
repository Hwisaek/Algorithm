def solution(arr, divisor):
    answer = sorted(filter(lambda x: x % divisor == 0, arr))
    return sorted(answer) if answer else [-1]
