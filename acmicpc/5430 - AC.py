import sys

input = sys.stdin.readline
t = int(input().rstrip())

for _ in range(t):
    p = input().rstrip().replace('RR', '')
    start = 0
    end = int(input().rstrip())
    array = input().rstrip()[1: -1].split(',')

    order = 1
    for opr in p:
        if opr == 'R':
            order *= -1
        elif opr == 'D':
            if order == 1:
                start += 1
            else:
                end -= 1

    array = array[start:end]
    if start <= end:
        print('[' + ','.join(array[::order]) + ']')
    else:
        print('error')
