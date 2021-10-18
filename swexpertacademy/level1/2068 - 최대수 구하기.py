t = int(input())

for i in range(t):
    array = list(map(int, input().split()))
    answer = max(array)
    print(f'#{i + 1} {answer}')
