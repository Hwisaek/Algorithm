from collections import Counter


def solution(N, stages):  # N: 스테이지의 개수, stages: 현재 도전 중인 스테이지 번호
    answer = []

    stages_clear = []  # 스테이지를 클리어한 사람 수

    count = Counter(stages)
    total = len(stages)

    for i in range(1, N + 1):
        total -= count[i]
        stages_clear.append(total)
        if count[i] == 0:
            answer.append((i, 0))
        else:
            answer.append((i, count[i] / (total + count[i])))

    answer = [idx for idx, v in sorted(answer, key=lambda x: -x[1])]

    return answer


print(solution(5, [2, 1, 2, 6, 2, 4, 3, 3]))
print(solution(4, [4, 4, 4, 4, 4]))
print(solution(5, [4, 4, 4, 4, 4]))
print(solution(6, [4, 4, 4, 4, 4]))
