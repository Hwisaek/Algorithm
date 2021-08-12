# key: m*m 크기의 열쇠, lock: n*n 크기의 자물쇠
def solution(key, lock):
    m = len(key)
    n = len(lock)

    new_lock = [[0] * (n * 3) for _ in range(n * 3)]

    for i in range(n):
        for j in range(n):
            new_lock[i + n][j + n] = lock[i][j]

    for rotation in range(4):
        key = rotate(key)
        for x in range(n * 2):
            for y in range(n * 2):

                for i in range(m):
                    for j in range(m):
                        new_lock[x + i][y + j] += key[i][j]

                if check(new_lock):
                    return True

                for i in range(m):
                    for j in range(m):
                        new_lock[x + i][y + j] -= key[i][j]

    return False


def check(new_lock):
    lock_length = len(new_lock) // 3
    for i in range(lock_length, lock_length * 2):
        for j in range(lock_length, lock_length * 2):
            if new_lock[i][j] != 1:
                return False
    return True


# 시계방향 90도 회전
def rotate(m):
    n = len(m)
    ret = [[0] * n for _ in range(n)]
    for r in range(n):
        for c in range(n):
            ret[c][n - 1 - r] = m[r][c]
    return ret


print(solution([[0, 0, 0], [1, 0, 0], [0, 1, 1]], [[1, 1, 1], [1, 1, 0], [1, 0, 1]]))
