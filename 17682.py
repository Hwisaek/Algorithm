def solution(dartResult):
    answer = 0
    # 점수를 문자로 변환
    bonus_char = [['10', 'k'], ['0', 'a'], ['1', 'b'], ['2', 'c'], ['3', 'd'], ['4', 'e'],
                  ['5', 'f'], ['6', 'g'], ['7', 'h'], ['8', 'i'], ['9', 'j']]

    # 10 같은 경우 2문자 이므로 점수를 한자리로 만들기 위해 점수를 문자로 치환
    for i in bonus_char:
        if i[0] in dartResult:
            dartResult = dartResult.replace(i[0], i[1])

    # 원본을 유지하고 리스트로 변환하여 저장
    dartResult_list = [0 for _ in range(9)]
    for i in range(len(dartResult)):
                   dartResult_list[i] = dartResult[i]

    # 옵션으로 설정되어 있던 보너스를 필수적으로 표시
    for i in range(1, 4):
        if dartResult_list[i * 3 - 1] == '*' or dartResult_list[i * 3 - 1] == '#':
            continue
        else:
            dartResult_list.insert(i * 3 - 1, ' ')
            dartResult_list.pop()

    # 문자를 다시 점수로
    for j in range(3):
        for i in range(len(bonus_char)):
            if dartResult_list[j * 3] == bonus_char[i][1]:
                dartResult_list[j * 3] = int(bonus_char[i][0])
                break

    # 보너스 적용
    for i in range(3):
        if dartResult_list[i * 3 + 1] == 'D':
            dartResult_list[i * 3] = dartResult_list[i * 3] ** 2
        elif dartResult_list[i * 3 + 1] == 'T':
            dartResult_list[i * 3] = dartResult_list[i * 3] ** 3

    # 옵션 적용
    for i in range(3):
        if dartResult_list[i * 3 + 2] == '#':
            dartResult_list[i * 3] = -dartResult_list[i * 3]
        elif dartResult_list[i * 3 + 2] == '*':
            dartResult_list[i * 3] = dartResult_list[i * 3] * 2
            if i > 0:
                dartResult_list[(i - 1) * 3] = dartResult_list[(i - 1) * 3] * 2

    answer = dartResult_list[0] + dartResult_list[3] + dartResult_list[6]
    return answer