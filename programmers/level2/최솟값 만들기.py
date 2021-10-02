# https://programmers.co.kr/learn/courses/30/lessons/12941
def solution(A, B):
    return sum(a * b for a, b in zip(sorted(A), sorted(B, reverse=True)))


def solution3(A, B):
    return sum(map(lambda x: x[0] * x[1], zip(sorted(A), sorted(B, reverse=True))))


def solution2(A, B):
    answer = 0
    A.sort()
    B.sort()

    while A:
        if A[-1] > B[-1]:
            answer += A[-1] * B[0]
            A.pop()
            B.pop(0)
        elif A[-1] < B[-1]:
            answer += A[0] * B[-1]
            A.pop(0)
            B.pop()
        else:
            if A[0] < B[0]:
                answer += A[0] * B[-1]
                A.pop(0)
                B.pop()
            else:
                answer += A[-1] * B[0]
                A.pop()
                B.pop(0)
    return answer


print(solution([1, 4, 2], [5, 4, 4]))
print(solution([1, 2], [3, 4]))
