# 시간:분 을 분으로 변환
def t2m(t):
    h, m = map(int, t.split(':'))
    return h * 60 + m

# 문자열의 A#, B#등을 a, b로 변환
def u2l(s):
    s_tmp = s
    for i in range(len(s)):
        if s[i] == '#':
            s_tmp = s_tmp.replace('{}#'.format(s[i - 1]), s[i - 1].lower())
            # print('s[i]: {}, s[i - 1]: {}'.format(s[i], s[i-1]))
    return s_tmp


def solution(m, musicinfos):
    answer = ''
    m = u2l(m)
    for i in musicinfos:
        start, end, title, melody = i.split(',')
        melody = u2l(melody)
        # print(start, end, title, melody)

        # 재생시간 = end - start, 분 단위
        playtime = t2m(end) - t2m(start)
        # print(playtime)

        # 재생시간 동안 재생된 멜로디
        song = []
        for j in range(playtime):
            song.append(melody[j % len(melody)])
        song = ''.join(song)
        # print(song)

        # 기억한 멜로디를 노래에서 찾으면 딕셔너리에 추가
        if m in song:
            if answer == '':
                answer = title
                answer_playtime = playtime
            else:
                if answer_playtime < playtime:
                    answer = title
                    answer_playtime = playtime
            # print(title)
            # print('m.index: {}'.format(song.index(m)))

        # 조건이 일치하는 음악이 없을 때
    if answer == '':
        answer = "(None)"

    return answer


# print(solution("ABCDEFG", ["12:00,12:14,HELLO,CDEFGAB",
#                 "13:00,13:05,WORLD,ABCDEF"]))
# print(solution("CC#BCC#BCC#BCC#B",
#                ["03:00,03:30,FOO,CC#B", "04:00,04:08,BAR,CC#BCC#BCC#B"]))
# print(solution("ABC",
#                ["12:00,12:14,HELLO,C#DEFGAB", "13:00,13:05,WORLD,ABCDEF"]))
print(solution("ABC",
               ["12:00,12:04,HELLO,A#BCDEF", "13:00,13:06,WORLD,AB#CDEF"]))