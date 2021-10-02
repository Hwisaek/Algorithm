# https://programmers.co.kr/learn/courses/30/lessons/12981
def solution(n, words):
    last = words[0][-1]
    lose = False
    used = set([words[0]])
    for i in range(1, len(words)):
        if last != words[i][0] or words[i] in used:
            lose = True
            break
        last = words[i][-1]
        used.add(words[i])

    if lose:
        return [i % n + 1, i // n + 1]
    else:
        return [0, 0]


print(solution(3, ["tank", "kick", "know", "wheel", "land", "dream", "mother", "robot", "tank"]))
print(solution(5,
               ["hello", "observe", "effect", "take", "either", "recognize", "encourage", "ensure", "establish", "hang",
                "gather", "refer", "reference", "estimate", "executive"]))
print(solution(2, ["hello", "one", "even", "never", "now", "world", "draw"]))
