import time


def solution(n, m, k, nums):
    nums.sort()

    roop = nums[-1] * k + nums[-2]  # 한번의 루프로 가장 큰 값을 만들 수 있는 수
    i = m // (k + 1)  # roop가 더해질 횟수
    j = m % (k + 1)  # 가장 큰 수만 더해질 횟수

    return roop * i + nums[-1] * j


n, m, k = 5, 8, 3  # N: 배열의 크기, M: 숫자가 총 더해지는 횟수, K: 연속해서 더할 수 있는 횟수
nums = [2, 4, 5, 4, 6]  # 더해질 숫자

start = time.time()  # 시작시간

print(solution(n, m, k, nums))

end = time.time()  # 종료시간
print('수행시간: ', end - start)  # 수행시간
