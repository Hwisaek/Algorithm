def solution(arr):
    answer = ""
    for i in sorted(arr, reverse=True):
        answer += str(i) + " "
    return answer


print(solution([3, 15, 27, 12]))
