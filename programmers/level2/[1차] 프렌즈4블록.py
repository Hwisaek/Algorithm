# https://programmers.co.kr/learn/courses/30/lessons/17679
dx = [1, 0, 1, 0]
dy = [0, 1, 1, 0]


def solution(m, n, board):
    answer = 0
    board = list(map(list, board))
    while True:
        visited = [[False] * n for _ in range(m)]
        chk = True  # 4블럭 제거 유무
        # 4블록 위치 찾기
        for x in range(m - 1):
            for y in range(n - 1):
                type = board[x][y]
                if type != '@':
                    found = True  # 2x2가 일치한지 확인
                    for i in range(3):
                        nx, ny = x + dx[i], y + dy[i]
                        if board[nx][ny] != type:
                            found = False
                            break

                    # 2x2 발견하면 해당 칸 기록
                    if found:
                        chk = False
                        for i in range(4):
                            nx, ny = x + dx[i], y + dy[i]
                            visited[nx][ny] = True

        for x in range(m):
            for y in range(n):
                if visited[x][y]:
                    for i in range(1, x + 1)[::-1]:
                        board[i][y] = board[i - 1][y]
                    board[0][y] = "@"
                    answer += 1
        if chk:
            break
    return answer


print(solution(4, 5, ["CCBDE", "AAADE", "AAABF", "CCBBF"]), solution(4, 5, ["CCBDE", "AAADE", "AAABF", "CCBBF"]) == 14)
print(solution(6, 6, ["TTTANT", "RRFACC", "RRRFCC", "TRRRAA", "TTMMMF", "TMMTTJ"]),
      solution(6, 6, ["TTTANT", "RRFACC", "RRRFCC", "TRRRAA", "TTMMMF", "TMMTTJ"]) == 15)
