def solution(jobs):
    answer = 0
    length = len(jobs)
    now = 0  # 현재시간

    while (len(jobs) != 0):
        job = jobs[0]
        now += job[1]
        answer += now - job[0]
        jobs.pop(0)

        jobs.sort(key=lambda x: x[1])

    answer /= length
    return answer


print(solution([[0, 3], [1, 9], [2, 6]]))
