def solution(n, k, A, B):
    A.sort()
    B.sort(reverse=True)
    for i in range(k):
        if A[i] >= B[i]:
            break
        A[i], B[i] = B[i], A[i]
    return sum(A)


print(solution(5, 3, [1, 2, 5, 4, 3], [5, 5, 6, 6, 5]))
