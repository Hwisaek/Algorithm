from collections import Counter

t = int(input())

for i in range(t):
    test_num = input()

    array = input().split()

    c = Counter(array)

    answer = -1
    max_num = -1

    for k, v in c.items():
        if v > max_num:
            answer = k
            max_num = v
        elif v == max_num:
            answer = max(answer, k)

    print(f'#{test_num} {answer}')
