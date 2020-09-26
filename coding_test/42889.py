def solution(N, stages):
    answer = []

    # 스테이지를 클리어 한 사람의 수
    stages_clear = []

    # 스테이지에 도달한 사람의 수
    stages_reach = []

    user = len(stages)

    for i in range(1, N + 1):
        # 해당 스테이지의 유저 수 카운트
        count = stages.count(i)
        # 해당 스테이지에 도달한 유저의 수
        stages_reach.append(user)
        stages_clear.append(user - count)
        # 다음 스테이지로 넘어갈 때 현재 층의 유저 수 만큼 감소
        user -= count

    # 실패율 계산
    fail_rate = []
    for i in range(N):
        if stages_reach[i] == 0:
            fail_rate.append(0)
        else:
            fail_rate.append((stages_reach[i] - stages_clear[i]) / stages_reach[i])
    fail_rate_tmp = [i for i in fail_rate]

    for i in range(N):
        answer.append(fail_rate.index(max(fail_rate)) + 1)
        fail_rate[fail_rate.index(max(fail_rate))] = -1

    print(answer)
    return answer