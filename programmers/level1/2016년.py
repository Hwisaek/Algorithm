def solution(a, b):
    day = ['SUN', 'MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT']
    d = [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    return day[(5 + sum(d[:a - 1]) + b - 1) % 7]


solution(5, 24)
