def solution(answers):
    a_1 = [1, 2, 3, 4, 5]
    a_2 = [2, 1, 2, 3, 2, 4, 2, 5]
    a_3 = [3, 3, 1, 1, 2, 2, 4, 4, 5, 5]

    score = [0, 0, 0]

    for i in range(len(answers)):
        if answers[i] == a_1[i % len(a_1)]:
            score[0] += 1
        if answers[i] == a_2[i % len(a_2)]:
            score[1] += 1
        if answers[i] == a_3[i % len(a_3)]:
            score[2] += 1
    max_s = max(score)
    winner = []
    for i in range(3):
        if max_s == score[i]:
            winner.append(i+1)
    return winner