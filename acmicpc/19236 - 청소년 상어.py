import copy

array = [[None] * 4 for _ in range(4)]

for i in range(4):
    data = list(map(int, input().split()))
    for j in range(4):
        # 각 위치마다 [물고기의 번호, 방향]을 저장
        array[i][j] = [data[j * 2], data[j * 2 + 1] - 1]

dx = [-1, -1, 0, 1, 1, 1, 0, -1]
dy = [0, -1, -1, -1, 0, 1, 1, 1]


def turn_left(direction):
    return (direction + 1) % 8


result = 0  # 최종 결과


# 특정한 번호의 물고기 찾기
def find_fish(array, index):
    for i in range(4):
        for j in range(4):
            if array[i][j][0] == index:
                return (i, j)
    return None


# 모든 물고기 이동
def move_all_fishes(array, now_x, now_y):
    for i in range(1, 17):
        # 물고기 찾기
        position = find_fish(array, i)
        if position != None:  # 해당 물고기가 존재하면
            x, y = position[0], position[1]
            direction = array[x][y][1]

            # 해당 물고기의 방향을 왼쪽으로 회전시키며 이동이 가능한지 확인
            for j in range(8):
                nx = x + dx[direction]
                ny = y + dy[direction]

                if 0 <= nx < 4 and 0 <= ny < 4:
                    if not (nx == now_x and ny == now_y):
                        array[x][y][1] = direction
                        array[x][y], array[nx][ny] = array[nx][ny], array[x][y]
                        break
                direction = turn_left(direction)


# 먹을 수 있는 모든 물고기의 위치 반환
def get_possible_positions(array, now_x, now_y):
    positions = []
    direction = array[now_x][now_y][1]

    for i in range(4):
        now_x += dx[direction]
        now_y += dy[direction]
        if 0 <= now_x < 4 and 0 <= now_y < 4:
            if array[now_x][now_y][0] != -1:
                positions.append((now_x, now_y))
    return positions


def dfs(array, now_x, now_y, total):
    global result
    array = copy.deepcopy(array)

    total += array[now_x][now_y][0]
    array[now_x][now_y][0] = -1

    move_all_fishes(array, now_x, now_y)

    positions = get_possible_positions(array, now_x, now_y)

    if len(positions) == 0:
        result = max(result, total)
        return

    for next_x, next_y in positions:
        dfs(array, next_x, next_y, total)


dfs(array, 0, 0, 0)
print(result)
