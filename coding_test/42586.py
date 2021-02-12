progresses = [93, 30, 55]
speeds = [1, 30, 5]

answer = []

real_answer = []

time = []

n = len(progresses)



for i in range(n):
    t = 0
    while True:
        prog = progresses[i] + speeds[i]*t
        t += 1
        if prog >= 100:
            break
    answer.append(t)

while True:
    if len(answer) == 0:
        break
    val = answer[0]
    t = 0
    for i in range(len(answer)):
        if answer[0] > val:
            break
        answer.pop(0)
        t += 1

    real_answer.append(t)
