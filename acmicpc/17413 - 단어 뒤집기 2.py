s = input() + ' '

answer = ""

while s:
    word = ''

    start, end = 0, 0

    in_bracket = False
    for idx, c in enumerate(s):
        if in_bracket:
            if c == '>':
                end = idx + 1
                in_bracket = False
                word += s[start:end]
        elif c == '<':
            start = idx
            in_bracket = True
            word += s[end:start][::-1]
        elif c == ' ':
            sep = idx
            word += s[end:sep][::-1]
            break

    s = s[sep + 1:]
    answer += word + ' '

print(answer)
