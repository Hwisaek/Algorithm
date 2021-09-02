T = int(input())
fibonacci = [[1, 0], [0, 1], [1, 1], [1, 2], [2, 3]]
for i in range(5, 41):
    fibonacci.append([fibonacci[-1][0] + fibonacci[-2][0], fibonacci[-1][1] + fibonacci[-2][1]])

for _ in range(T):
    N = int(input())
    print(fibonacci[N][0], fibonacci[N][1])

# 2 = 1 + 0
# 3 = 2 + 1 = 2 * 1 + 0
# 4 = 3 + 2 = 3 * 1 + 2 * 0
# 5 = 4 + 3 = 5 * 1 + 3 * 0
# 6 = 5 + 4 = 8 * 1 + 5 * 0
# 7 = 6 + 5 = 13 * 1 + 8 * 0
