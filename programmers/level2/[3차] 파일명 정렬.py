# https://programmers.co.kr/learn/courses/30/lessons/17686
def solution(files):
    for i in range(len(files)):
        start = numberIdx(files[i])
        end = tailIdx(files[i], start)
        files[i] = (files[i][:start], files[i][start:end], files[i][end:])
    return list(map(''.join, sorted(files, key=lambda x: (x[0].lower(), int(x[1])))))


def numberIdx(s):
    for i in range(len(s)):
        if s[i].isdigit():
            return i


def tailIdx(s, idx):
    for i in range(idx, len(s)):
        if not s[i].isdigit():
            return i
    return len(s)


print(solution(["MUZI01.zip", "muzi1.png"]))
print(solution(["muzi1.png", "MUZI01.zip"]))
print(solution(["foo9.txt", "foo010bar020.zip", "F-15"]))
print(solution(["img12.png", "img10.png", "img02.png", "img1.png", "IMG01.GIF", "img2.JPG"]))
print(solution(["F-5 Freedom Fighter", "B-50 Superfortress", "A-10 Thunderbolt II", "F-14 Tomcat"]))
