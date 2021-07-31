def solution(priorities, location):
    answer = 0

    while len(priorities) > 0:
        if priorities[0] == max(priorities):  # 우선순위 확인
            # 맨 앞 문서 인쇄
            priorities.pop(0)
            answer += 1
            if location == 0:  # 요청한 문서이면 반복문 종료
                break
        else:
            # 맨 앞 문서를 대기 목록의 가장 마지막에 추가
            priorities.append(priorities.pop(0))
        location = location - 1 if location > 0 else len(priorities) - 1  # 요청 문서 위치 변경
    return answer


# solution([2, 1, 3, 2], 2)
solution([1, 1, 9, 1, 1, 1], 0)
# solution([9, 1, 9, 1, 2, 1], 3)
