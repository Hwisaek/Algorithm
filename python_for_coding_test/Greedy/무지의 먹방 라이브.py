from collections import deque


# food_times: 각 음식을 모두 먹는데 필요한 시간, k: 방송이 중단된 시간
def solution(food_times, k):
    answer = 0
    food_times = [[idx, val] for idx, val in enumerate(food_times)]
    food_times.sort(reverse=True, key=lambda x: x[1])
    last = len(food_times) - 1
    while k // len(food_times) != 0:
        k -= len(food_times)
        food_times = [[i[0], i[1] - 1] for i in food_times]
        while food_times[last][1] == 0:
            last -= 1
            if last <= 0:
                return -1
        food_times = food_times[:last + 1]

    food_times.sort(key=lambda x: x[0])
    answer = food_times[k][0]
    return answer + 1


print(solution([3, 1, 2], 5))  # 1
print(solution([3, 1, 2], 8))  # -1
print(solution([3, 1, 2], 1))  # 2
print(solution([1], 5))  # -1
print(solution([5], 5))  # -1
# print(solution([200000] * int(1e8), 2e13))
