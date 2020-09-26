def solution(bridge_length, weight, truck_weights):
    answer = 0
    
    # 다리 위에 올라간 트럭을 표현
    bridge = [0] * bridge_length
    
    time = 0
    
    for _ in range(bridge_length * len(truck_weights) + 1):
        # 1초가 지남
        # print(bridge)
        del bridge[0]
        if len(bridge) < bridge_length:
            bridge.append(0)
        time += 1
        
        # 다리가 무게를 버티고 남은 트럭이 있으면 트럭 추가
        if truck_weights and sum(bridge) + truck_weights[0] <= weight:
            bridge[bridge_length - 1] = truck_weights[0]
            truck_weights = truck_weights[1:]
        
        # 모든 트럭이 다리를 건너면 반복문 탈출
        if ((truck_weights == []) and (sum(bridge) == 0)):
            break
    answer = time
    print(answer)
    return answer


solution(1, 100, [10])