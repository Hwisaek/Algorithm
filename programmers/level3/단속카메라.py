def solution(routes):
    answer = 0
    cam = -30000
    routes.sort(key=lambda x: x[1])  # 진출지점 기준 정렬
    chk = [False] * len(routes)  # 카메라 단속 체크

    for idx, route in enumerate(routes):
        if not chk[idx]:  # 카메라를 만나지 않았으면 실행
            cam = route[1]  # 카메라를 해당 차량의 진출지점에 설치
            for i, r in enumerate(routes):
                if not chk[i] and r[0] <= cam <= r[1]:
                    chk[i] = True
            answer += 1
    return answer


print(solution([[-20, 15], [-14, -5], [-18, -13], [-5, -3]]))
