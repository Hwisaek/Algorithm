import time


def solution(n, m, arr):
    return max(map(min, arr))  # 각 행의 최소값으로 배열 생성 후 최대값  반환


n, m = map(int, input().split())  # n: 행의 개수, m: 열의 개수
arr = []  # 숫자카드 행렬
for i in range(n):  # 행렬 입력 받기
    arr.append(list(map(int, input().split())))

start = time.time()  # 시작시간

print(solution(n, m, arr))

end = time.time()  # 종료시간
print('수행시간: ', end - start)  # 수행시간
