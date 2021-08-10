import time


def solution(charge):
    coins = [500, 100, 50, 10]
    count = [0, 0, 0, 0]

    for i, coin in enumerate(coins):
        count[i] += charge // coin
        charge %= coin

    print(f'500원: {count[0]}, 100원: {count[1]}, 50원: {count[2]}, 10원: {count[3]}')
    return sum(count)


start = time.time()
print(solution(1260))
end = time.time()
print('수행시간: ', end - start)
