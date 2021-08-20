def solution(n):
    answer = []

    def path(n, start, goal):
        answer.append([start, goal])

    def hanoi(n, start, goal, assistant):
        if n == 1:
            path(n, start, goal)
        else:
            hanoi(n - 1, start, assistant, goal)
            path(n, start, goal)
            hanoi(n - 1, assistant, goal, start)

    hanoi(n, 1, 3, 2)
    return answer


print(solution(2))
