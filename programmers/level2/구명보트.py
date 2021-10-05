# https://programmers.co.kr/learn/courses/30/lessons/42885
def solution(people, limit):
    answer = 0
    people.sort()
    l, r = 0, len(people) - 1
    while l <= r:
        if people[l] + people[r] <= limit:
            l += 1
            r -= 1
        else:
            r -= 1
        answer += 1
    return answer


print(solution([70, 50, 80, 50], 100))  # 3
print(solution([70, 80, 50], 100))  # 3
print(solution([70, 50, 80, 50, 50, 40], 100))  # 4
print(solution([100], 100))  # 1
print(solution([50], 100))  # 1
print(solution([40, 60, 60, 70, 100], 100))  # 4
print(solution([40, 40, 40, 50, 50, 50, 50, 60], 100))  # 4
