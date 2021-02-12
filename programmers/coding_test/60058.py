def chk_s(p):
    if p[0] == ')':
        return False
    else:
        return True


def get_uv(p):
    # 문자열을 리스트로
    p = list(p)
    # n = p의 길이, 코드 간략화
    n = len(p)

    u, v = "", ""

    cnt_l = 0
    cnt_r = 0

    for i in range(n):
        if p[i] == '(':
            cnt_l += 1
        elif p[i] == ')':
            cnt_r += 1

        if cnt_l or cnt_r:
            if cnt_l == cnt_r:
                u = p[:i + 1]
                v = p[i + 1:]
                break
    return ''.join(u), ''.join(v)

def reverse(c):
    if c == '(':
        return ')'
    else:
        return'('

def solution(s):
    str_tmp = ''
    if s:
        u, v = get_uv(s)
        if chk_s(u):
            str_tmp += u
            return str_tmp + solution(v)
        else:
            str_tmp += '('
            str_tmp += (solution(v) + ')')
            u = list(u[1:-1])
            for i in range(len(u)):
                u[i] = reverse(u[i])
            u = ''.join(u)
            str_tmp += u
            return str_tmp
    else:
        return ""


print(solution("(()())()"))
print(solution(")("))
print(solution("()))((()"))