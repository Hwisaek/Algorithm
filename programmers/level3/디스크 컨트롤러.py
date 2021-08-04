import heapq
from collections import deque


def solution(jobs):  # 요청시간, 소요시간, 시작시간
    answer = 0
    length = len(jobs)
    jobs.sort(key=lambda x: x[0])  # 요청시간 순 정렬
    jobs = deque(jobs)

    time = 0  # 시간
    run = 0  # 남은 작업시간
    q = deque()  # 요청된 작업 목록

    task = [0, 0]
    while True:  # 작업이 남아있고, 요청된 작업이 남아있고, 아직 작업중이면 반복
        for job in jobs:  # 작업 요청 처리
            if job[0] <= time:
                q.append(jobs.popleft())
            else:
                break

        # 작업 수행
        if task[1] > 0:
            task[1] -= 1
        else:
            # 대기시간 = 종료시간 - 요청시간
            answer += time - task[0]
            task = q.popleft()

        time += 1
        if not jobs and not q and task[0] == 0:
            break

    return answer / length


print(solution([[0, 3], [2, 6], [1, 9]]))
