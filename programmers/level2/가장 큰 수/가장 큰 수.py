def solution(numbers):
    answer = ''

    numbers = list(map(str(x) * 3, numbers))

    numbers.sort(reverse=True, key=lambda x: x * 3)

    for n in numbers:
        answer += n
    return str(int(answer))


# solution([6, 10, 2])
# solution([3, 30, 34, 5, 9])
solution([0, 0, 0, 0])
