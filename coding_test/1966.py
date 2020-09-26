t = int(input())

# t회 만큼 반복

for _ in range(t):
    N, M = map(int, input().split())
    Q = list(map(int, input().split()))

    # 프린트 한 횟수
    print_n = 0
    
    # 원하는 문서가 출력 될 때까지 반복
    while True:
        # 최상단 문서가 우선순위가 가장 높은가?
        if Q[0] == max(Q):
            # 높으면 출력한 문서 수를 1 증가
            print_n += 1
            
            # 출력할 문서가 원하는 문서인가?
            if M == 0:
                print(print_n)
                break
            # 아니면 출력할 문서를 내보내고 원하는 문서의 위치 1 감소
            else:
                Q = Q[1:]
                M -= 1
        else:
            # 우선 순위가 가장 높지 않으면 최상단의 문서를 맨 아래로 이동
            Q.append(Q[0])
            Q = Q[1:]
            # 문서의 순서가 변경 되었으므로 M의 값도 하나 감소
            if M == 0:
                M = len(Q) - 1
            else:
                M -= 1