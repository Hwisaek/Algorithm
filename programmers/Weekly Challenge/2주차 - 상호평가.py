def solution(scores):
    scores = [[scores[j][i] for j in range(len(scores))] for i in range(len(scores))]
    sum = [0.] * len(scores)
    for i in range(len(scores)):
        count = 0
        for j in range(len(scores)):
            if i == j and ((scores[i][j] == max(scores[i]) and scores[i].count(max(scores[i])) == 1) or
                           (scores[i][j] == min(scores[i]) and scores[i].count(min(scores[i])) == 1)):
                continue
            sum[i] += scores[i][j]
            count += 1
        sum[i] /= count

    return ''.join(map(grade, sum))


def grade(score):
    if score >= 90:
        return "A"
    elif score >= 80:
        return "B"
    elif score >= 70:
        return "C"
    elif score >= 50:
        return "D"
    else:
        return "F"


print(solution(
    [[100, 90, 98, 88, 65], [50, 45, 99, 85, 77], [47, 88, 95, 80, 67], [61, 57, 100, 80, 65], [24, 90, 94, 75, 65]]))
# print(solution([[50, 90], [50, 87]]))
# print(solution([[70, 49, 90], [68, 50, 38], [73, 31, 100]]))
# print(solution([[80, 90], [60, 70]]))
# print(solution([[80, 90, 100]
#                    , [90, 80, 80]
#                    , [70, 70, 90]]))
