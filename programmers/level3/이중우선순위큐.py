import heapq


def solution(operations):
    answer = []

    for op in operations:
        oper = op[0]
        value = int(op[2:])
        if oper == "I":  # 삽입
            heapq.heappush(answer, value)
        else:  # 삭제
            if len(answer) > 0:
                if value == 1:  # 최댓값 삭제
                    answer = [-i for i in answer]  # heapq는 최소힙이므로 -1 곱하기 연산
                    heapq.heapify(answer)
                    heapq.heappop(answer)
                    answer = [-i for i in answer]  # -1 곱해진걸 되돌림
                else:  # 최솟값 삭제
                    heapq.heapify(answer)
                    heapq.heappop(answer)

    if len(answer) == 0:
        return [0, 0]
    return [max(answer), min(answer)]


# print(solution(["I 16", "I -5643", "D -1", "D 1", "D 1", "I 123", "D -1"]))
print(solution(["I -45", "I 653", "D 1", "I -642", "I 45", "I 97", "D 1", "D -1", "I 333"]))
