t = int(input())

for i in range(t):
    x_1, y_1, r_1, x_2, y_2, r_2 = map(int, input().split())

    d = ((x_1 - x_2) ** 2 + (y_1 - y_2) ** 2) ** 0.5

    if (d > r_1 + r_2) or (abs(r_1 - r_2) > d):
        print(0)
    elif (d == 0) and (r_1 == r_2):
        print(-1)
    elif (d == r_1 + r_2) or (d == abs(r_1 - r_2)):
        print(1)
    else:
        print(2)