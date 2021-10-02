# https://programmers.co.kr/learn/courses/30/lessons/17680
from collections import deque


# set을 추가로 사용하지만 시간이 살짝 짧아짐
def solution(cacheSize, cities):
    answer = 0

    s = set()
    q = deque()
    for city in cities:
        city = city.lower()
        if city in s:
            q.remove(city)
            q.append(city)
            answer += 1
        else:
            q.append(city)
            s.add(city)
            answer += 5

        if len(q) > cacheSize:
            s.remove(q.popleft())

    return answer


# 코드는 간단하나 살짝 시간이 더 걸림
def solution2(cacheSize, cities):
    answer = 0

    q = deque(maxlen=cacheSize)
    for city in cities:
        city = city.lower()
        if city in q:
            q.remove(city)
            q.append(city)
            answer += 1
        else:
            q.append(city)
            answer += 5

    return answer


print(solution(3, ["Jeju", "Pangyo", "Seoul", "NewYork", "LA", "Jeju", "Pangyo", "Seoul", "NewYork", "LA"]))
print(solution(3, ["Jeju", "Pangyo", "Seoul", "Jeju", "Pangyo", "Seoul", "Jeju", "Pangyo", "Seoul"]))
print(solution(
    2,
    ["Jeju", "Pangyo", "Seoul", "NewYork", "LA", "SanFrancisco", "Seoul", "Rome", "Paris", "Jeju", "NewYork", "Rome"]))
print(solution(
    5,
    ["Jeju", "Pangyo", "Seoul", "NewYork", "LA", "SanFrancisco", "Seoul", "Rome", "Paris", "Jeju", "NewYork", "Rome"]))
print(solution(2, ["Jeju", "Pangyo", "NewYork", "newyork"]))
print(solution(0, ["Jeju", "Pangyo", "Seoul", "NewYork", "LA"]))
