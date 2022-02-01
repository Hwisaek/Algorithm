def solution(p, c):
    count = {}
    for i in p:
        count[i] = count.get(i, 0) + 1

    for i in c:
        count[i] -= 1
        if (count[i] == 0):
            count.pop(i)

    return list(count.keys())[0]

solution(["leo", "kiki", "eden"], ["eden", "kiki"])
solution(["marina", "josipa", "nikola", "vinko", "filipa"], ["josipa", "filipa", "marina", "nikola"])
solution(["mislav", "stanko", "mislav", "ana"], ["stanko", "ana", "mislav"])
