# https://softeer.ai/practice/info.do?eventIdx=1&psProblemId=408
gears = list(map(int, input().split()))


def solution(gears):
    last = gears[0]
    if last == 1:
        answer = "ascending"
    elif last == 8:
        answer = "descending"
    else:
        return "mixed"

    for gear in gears[1:]:
        if answer == "ascending" and gear == last + 1:
            last = gear
        elif answer == "descending" and gear == last - 1:
            last = gear
        else:
            return "mixed"
    return answer


print(solution(gears))
