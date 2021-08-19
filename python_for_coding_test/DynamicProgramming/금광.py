def solution(T, arr):  # T: 테스트케이스의 수
    for size, array in arr:

        # 금광 생성
        gold = []
        for i in range(size[0]):
            gold.append([])
            for j in range(size[1]):
                gold[i].append(array[i * size[1] + j])

        answer = gold.copy()
        for j in range(1, size[1]):
            for i in range(size[0]):
                total = []
                if i > 0:
                    total.append(gold[i][j] + answer[i - 1][j - 1])
                total.append(gold[i][j] + answer[i][j - 1])
                if i + 1 <= size[0] - 1:
                    total.append(gold[i][j] + answer[i + 1][j - 1])
                answer[i][j] = max(total)

        print(max([answer[i][size[1] - 1] for i in range(size[0])]))


solution(2,
         [[[3, 4], [1, 3, 3, 2, 2, 1, 4, 1, 0, 6, 4, 7]],
          [[4, 4], [1, 3, 1, 5, 2, 2, 4, 1, 5, 0, 2, 3, 0, 6, 1, 2]]])
