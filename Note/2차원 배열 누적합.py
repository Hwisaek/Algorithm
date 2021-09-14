# 2차원 배열의 누적합 배열 구하기
## 누적합, 부분합, 구간합

def get2dSum(array):
    n = len(array)
    m = len(array[0])

    dp = [[0] * (m + 1) for _ in range(n + 1)]
    for i in range(1, n + 1):
        for j in range(1, m + 1):
            dp[i][j] = array[i - 1][j - 1] + dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1]

    return dp


import sys

input = sys.stdin.readline

n, m = map(int, input().split())

array = [list(map(int, input().split())) for _ in range(n)]

# 누적 합 배열
dp = get2dSum(array)

k = int(input())
for _ in range(k):
    i, j, x, y = map(int, input().split())
    print(dp[x][y] - dp[i - 1][y] - dp[x][j - 1] + dp[i - 1][j - 1])
