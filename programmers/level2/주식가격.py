def solution(prices):
    # prices 의 길이만큼 answer 배열 생성
    answer = [0] * len(prices)

    # 자신보다 뒤에 있는 값들과 비교하면서, 나보다 작은 값을 만날 때까지 1씩 증가시킴
    for i, price in enumerate(prices):
        for j in range(i + 1, len(prices)):
            if prices[i] > prices[j]:  # 자신보다 작으면 반복 종료
                break
        answer[i] = j - i

    print(answer)
    return answer


solution([1, 2, 3, 2, 3])
solution([9, 8, 7, 6, 5])
solution([1, 1, 1, 1, 1])
