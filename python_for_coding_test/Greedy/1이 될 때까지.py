import time


def solution(n, k):
    count = 0

    while n > 1:  # n 이 1이 될 때까지 반복
        if n % k == 0:  # n이 k로 나누어 떨어지면 k로 나눔
            count += 1  # 수행 횟수 추가
            n /= k
        else:  # n이 k로 나누어 떨어지지 않으면 -1
            count += n % k  # 수행 횟수 추가
            n -= n % k

    if n == 0:
        count -= 1

    return count


n, k = map(int, input().split())
start = time.time()  # 시작시간
print(solution(n, k))
end = time.time()  # 종료시간
print('수행시간: ', end - start)  # 수행시간
