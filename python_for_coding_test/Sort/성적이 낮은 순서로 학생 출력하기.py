def solution(n, arr):
    return ''.join(map(lambda x: str(x[0]) + " ", sorted(arr, key=lambda x: x[1])))


print(solution(2, [["홍길동", 95], ["이순신", 77]]))
