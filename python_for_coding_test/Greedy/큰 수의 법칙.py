import time


def solution():
    n, m, k = map(int, input().split())  # N: 배열의 크기, M: 숫자가 총 더해지는 횟수, K: 연속해서 더할 수 있는 횟수
    nums = list(map(int, input().split()))  # 더해질 숫자

    start = time.time()  # 시작시간
    nums.sort()

    roop = nums[-1] * k + nums[-2]  # 한번의 루프로 가장 큰 값을 만들 수 있는 수
    i = m // (k + 1)  # roop가 더해질 횟수
    j = m % (k + 1)  # 가장 큰 수만 더해질 횟수

    end = time.time()  # 종료시간
    print('수행시간: ', end - start)  # 수행시간

    return roop * i + nums[-1] * j


print(solution())
