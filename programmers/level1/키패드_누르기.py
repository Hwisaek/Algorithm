import numpy as np


def solution(numbers, hand):
    hand = hand[:1].upper()
    answer = ''
    left = np.array([3, 0])
    right = np.array([3, 2])
    for n in numbers:
        if n % 3 == 2 or n == 0:
            btn = np.array([3 if n == 0 else n // 3, 1])
            distL = np.abs(btn - left).sum()
            distR = np.abs(btn - right).sum()
            if distR < distL:
                right = btn
                answer += 'R'
            elif distL < distR:
                left = btn
                answer += 'L'
            else:
                if hand == 'R':
                    right = btn
                else:
                    left = btn
                answer += hand
        elif n % 3 == 1:
            left = [n // 3, 0]
            answer += 'L'
        elif n % 3 == 0:
            right = [n // 3 - 1, 2]
            answer += 'R'
    return answer


print(solution([1, 3, 4, 5, 8, 2, 1, 4, 5, 9, 5], "right"))
print(solution([7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2], "left"))
