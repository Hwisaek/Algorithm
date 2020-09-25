def move(n, x, y):
    print('{} {}'.format(x, y))

def hanoi(n, x, y, z):
    if n == 1:
        move(n, x, y)
    else:
        hanoi(n - 1, x, z, y)
        move(n, x, y)
        hanoi(n -1, z, y, x)

n = int(input())

print(2 ** n - 1)
hanoi(n, 1, 3, 2)