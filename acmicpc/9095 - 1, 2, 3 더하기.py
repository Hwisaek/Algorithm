import sys

input = sys.stdin.readline

t = int(input())

for _ in range(t):
    n = int(input())

    count = 0


    def dp(n):
        global count
        if n == 0:
            count += 1
            return
        elif n < 0:
            return

        dp(n - 1), dp(n - 2), dp(n - 3)


    dp(n)
    print(count)
