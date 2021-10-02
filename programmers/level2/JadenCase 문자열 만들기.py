# https://programmers.co.kr/learn/courses/30/lessons/12951

def solution(s):
    arr = s.split(" ")
    for i in range(len(arr)):
        if arr[i]:
            arr[i] = arr[i][0].upper() + arr[i][1:].lower()
    return ' '.join(arr)


print(solution("3people unFollowed me"))
print(solution("for the last week"))
print(solution("forthelastweek"))
print(solution("for  the last week"))
