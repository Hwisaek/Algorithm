def solution(bridge_length, weight, truck_weights):
    answer = 0
    bridge = [0 for _ in range(bridge_length)]
    total = 0  # sum(bridge) 시간초과 대비용 변수

    while total != 0 or len(truck_weights) > 0:  # 다리에 트럭이 없고, 모든 트럭이 지날 때 까지 반복
        total -= bridge.pop(0)  # 맨 앞 트럭 통과, 다리 적재량 감소
        if truck_weights and total + truck_weights[0] <= weight:  # 다음 트럭의 무게를 버틸 수 있으면 통과
            truck = truck_weights.pop(0)  # 다음에 통과할 트럭
            total += truck  # 다리 적재량 중가
            bridge.append(truck)  # 다리에 다음 트럭 진입
        else:
            bridge.append(0)
        answer += 1  # 1초 증가

    return answer


solution(2, 10, [7, 4, 5, 6])
solution(100, 100, [10])
solution(100, 100, [10, 10, 10, 10, 10, 10, 10, 10, 10, 10])
