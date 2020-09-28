def solution(jobs):
    answer = 0
    time = 0
    queue = []
    jobs.sort()
    
    doing = 0
    
    while True:
        if doing == 0 and jobs == [] and queue == []:
            break
        
        for i in jobs:
            if i[0] == time:
                queue.append(i[1])
                jobs.remove(i)
                break
        
        # 처리중인 작업이 있으면
        if doing:
            doing -= 1
        else:
            # 없으면 queue 에서 소요시간이 가장 작은 것을 추가
            if queue:
                doing = min(queue)
                queue.remove(doing)
        time += 1
        
    answer = time
    return answer


print(solution([[0, 3], [1, 9], [2, 6]]))