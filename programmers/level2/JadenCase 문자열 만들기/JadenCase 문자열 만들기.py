# https://programmers.co.kr/learn/courses/30/lessons/12951

def solution(s):
    arr = s.split(" ")
    for i in range(len(arr)):
        arr[i] = arr[i].capitalize()
    return ' '.join(arr)


print(solution("3people unFollowed me"))  # 3people Unfollowed Me
print(solution("for the last week"))  # For The Last Week
print(solution("forthelastweek"))  # Forthelastweek
print(solution("for  the last week"))  # For  The Last Week
