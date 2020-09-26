def solution(skill, skill_trees):
    answer = 0
    
    # 스킬순서를 리스트로 설정
    skill = list(skill)
    
    # 스킬 트리를 하나씩 불러옴
    for i in skill_trees:
        # 불러온 스킬트리를 임시저장
        tmp = list(i)
        
        # 스킬트리에서 하나씩 불러옴
        for j in i:
            # 불러온 스킬이 스킬 순서에 해당하지 않는 경우 해당 스킬을 삭제
            if j not in skill:
                tmp.remove(j)
        
        cnt = 0
        for j in range(len(tmp)):
            if skill[j] == tmp[j]:
                cnt += 1
            else:
                cnt = 0
                break
        if cnt:
            answer += 1
        # 스킬트리에서 아무것도 안 배운 경우
        elif len(tmp) == 0:
            answer += 1
        print(tmp, cnt)
    return answer

print(solution("CBD", ["OPQ"]))