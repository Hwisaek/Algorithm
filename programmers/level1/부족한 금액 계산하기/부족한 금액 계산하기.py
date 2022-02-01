def solution(price, money, count):
    return max(price * count * (count + 1) / 2, 0)


solution(3, 20, 4)
