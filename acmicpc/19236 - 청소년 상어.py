fish = [[-1] * 4 for _ in range(4)]
fish_d = [[-1] * 4 for _ in range(4)]

# 0번 인덱스 사용 x
dx = [0, 0, -1, -1, -1, 0, 1, 1, 1]
dy = [0, -1, -1, 0, 1, 1, 1, 0, -1]

for i in range(4):
    data = list(map(int, input().split()))
    for j in range(4):
        fish[i][j] = data[j * 2]
        d[i][j] = data[j * 2 + 1]

shark = [0, 0, 0]  # 상어의 좌표 x, y, 방향

total = 0


# 1. 상어가 현재 위치의 물고기를 먹고 같은 방향을 가짐
# 2. 물고기 이동
# 3. 상어가 보는 방향 탐색
# 3-1. 물고기가 존재하면 해당 위치로 이동 후 1로 이동
# 3-2. 물고기가 존재하지 않으면 종료

def dfs(fish, fish_d, shark):
    # 1. 상어가 현재 위치의 물고기를 먹고 같은 방향을 가짐
    x, y = shark  # 상어의 좌표
    d = fish_d[x][y]  # 상어의 방향
    size = fish[x][y]  # 물고기 크기만큼 사이즈 증가
    fish[x][y] = -1  # 물고기 삭제

    fish, fish_d = move(fish, fish_d)

    lst = []
    # 상어가 보는 방향 탐색
    for i in range(1, 4):
        nx = x + dx[d] * i
        ny = y + dy[d] * i
        if not 0 <= x < 4 or not 0 <= y < 4:  # 범위를 벗어나면 귀환
            return 0
        return total + max(dfs(fish, fish_d, [nx, ny, nd]))

def move(fish, fish_d):

    return fish, fish_d


dfs(fish, fish_d, shark)

print(total)
