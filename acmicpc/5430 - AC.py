import sys

input = sys.stdin.readline
t = int(input())

for _ in range(t):
    p = input()
    n = input()
    array = input()[1: -1]

    error = False
    order = True
    for opr in p:
        if opr == 'R':
            order = not order
        elif opr == 'D':
            if not array:
                error = True
                print('error')
            elif order:
                array = array[2:]
            else:
                array = array[:-2]
    if not error:
        print('[' + (array if order else array[::-1]) + ']')
