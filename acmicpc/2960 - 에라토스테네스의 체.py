n, k = map(int, input().split())

prime = [True] * (n + 1)  # 소수 판별

for i in range(2, n + 1):
    if prime[i]:  # 소수면 실행
        for j in range(i, n + 1, i):  # 소수의 배수 지우기
            if prime[j]:
                prime[j] = False
                k -= 1
                if k == 0:
                    print(j)
                    break
    if k == 0:
        break
