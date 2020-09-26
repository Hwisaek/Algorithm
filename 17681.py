def solution(n, arr1, arr2):
    answer = []
    
    # 0, 1은 arr1, arr2, 2는 둘을 합친 것
    map = [[], [], []]
    for i in range(n):
        map[2].append([])
    
    for i in arr1:
        # 배열의 값을 이진수로 변환하여 저장
        i_bin = '{0:b}'.format(i).zfill(n)
        i_bin = [int(i) for i in i_bin]
        map[0].append(i_bin)
    # print(map[0])
    
    for i in arr2:
        # 배열의 값을 이진수로 변환하여 저장
        i_bin = '{0:b}'.format(i).zfill(n)
        i_bin = [int(i) for i in i_bin]
        map[1].append(i_bin)
    # print(map[1])
    
    # 지도 두개를 합침
    for i in range(n):
        for j in range(n):
            map[2][i].append(map[0][i][j] + map[1][i][j])
    # print(map[2])
    
    # 0은 공백, 숫자는 #으로 표시
    for i in range(n):
        for j in range(n):
            if map[2][i][j] == 0:
                map[2][i][j] = ' '
            else:
                map[2][i][j]  = '#'
    # print(map[2])
    
    for i in range(n):
        map[2][i] = ''.join(map[2][i])
    
    # print(map[2])
    answer = map[2]
    return answer


print(solution(5, [9, 20, 28, 18, 11], [30, 1, 21, 17, 28]))