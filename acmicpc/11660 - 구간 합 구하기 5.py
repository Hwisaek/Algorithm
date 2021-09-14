import sys

input = sys.stdin.readline


def get2dSum(array):
    n = len(array)
    m = len(array[0])

    dp = [[0] * (m + 1) for _ in range(n + 1)]
    for i in range(1, n + 1):
        for j in range(1, m + 1):
            dp[i][j] = array[i - 1][j - 1] + dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1]

    return dp


def cal_sum(dp, i, j, x, y):
    return dp[x][y] - dp[i - 1][y] - dp[x][j - 1] + dp[i - 1][j - 1]


n, m = map(int, input().split())

array = [list(map(int, input().split())) for _ in range(n)]

dp = get2dSum(array)

for _ in range(m):
    x1, y1, x2, y2 = map(int, input().split())
    print(cal_sum(dp, x1, y1, x2, y2))
