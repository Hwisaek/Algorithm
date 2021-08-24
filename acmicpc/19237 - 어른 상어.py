n, m, k = map(int, input().split())

# 상어의 위치, 방향을 포함하는 2차원 리스트
array = []
for i in range(n):
    array.append(list(map(int, input().split())))

# 상어의 방향 정보
directions = list(map(int, input().split()))

# 각 격자마다 [상어 번호, 냄새 잔여 시간]을 저장하는 2차원 리스트
smell = [[[0, 0]] * n for _ in range(n)]

# 각 상어의 회전 방향 우선순위 정보
priorities = [[] for _ in range(m)]
for i in range(m):
    for j in range(4):
        priorities[i].append(list(map(int, input().split())))

dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]


# 냄새 정보 업데이트
def update_smell():
    for i in range(n):
        for j in range(n):
            if smell[i][j][1] > 0:  # 냄새가 존재하면, 잔여시간 1 감소
                smell[i][j][1] -= 1
            # 상어가 존재하는 해당 위치의 냄새를 k로 설정
            if array[i][j] != 0:
                smell[i][j] = [array[i][j], k]


# 모든 상어 이동
def move():
    new_array = [[0] * n for _ in range(n)]

    for x in range(n):
        for y in range(n):
            if array[x][y] != 0:  # 상어가 존재하는 경우
                direction = directions[array[x][y] - 1]  # 현재 상어의 방향
                found = False

                for index in range(4):  # 냄새가 존재하지 않는 곳이 있는지 확인
                    nx = x + dx[priorities[array[x][y] - 1][direction - 1][index] - 1]
                    ny = y + dy[priorities[array[x][y] - 1][direction - 1][index] - 1]
                    if 0 <= nx < n and 0 <= ny < n:
                        if smell[nx][ny][1] == 0:  # 냄새가 존재하지 않으면 상어 이동
                            directions[array[x][y] - 1] = priorities[array[x][y] - 1][direction - 1][index]

                            if new_array[nx][ny] == 0:  # 상어가 없으면 이동
                                new_array[nx][ny] = array[x][y]
                            else:  # 상어가 있으면 작은 번호의 상어가 차지하도록 설정
                                new_array[nx][ny] = min(new_array[nx][ny], array[x][y])
                            found = True
                            break
                if found:
                    continue

                # 냄새가 존재하지 않는 칸이 없으면 실행
                for index in range(4):  # 자신의 냄새가 존재하는 칸 탐색
                    nx = x + dx[priorities[array[x][y] - 1][direction - 1][index] - 1]
                    ny = y + dy[priorities[array[x][y] - 1][direction - 1][index] - 1]
                    if 0 <= nx < n and 0 <= ny < n:
                        if smell[nx][ny][0] == array[x][y]:  # 자신의 냄새가 있는 곳이면 이동
                            directions[array[x][y] - 1] = priorities[array[x][y] - 1][direction - 1][index]
                            new_array[nx][ny] = array[x][y]
                            break
    return new_array


time = 0
while True:
    update_smell()
    new_array = move()
    array = new_array
    time += 1

    check = True
    for i in range(n):
        for j in range(n):
            if array[i][j] > 1:
                check = False

    if check:
        print(time)
        break

    if time >= 1000:
        print(-1)
        break
