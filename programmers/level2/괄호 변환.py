def solution(w):
    d = {'(': ')', ')': '('}
    while True:
        # 1. 입력이 빈 문자열인 경우, 빈 문자열을 반환합니다.
        if len(w) == 0:
            return w

        # 2. 문자열 w를 두 "균형잡힌 괄호 문자열" u, v로 분리합니다. 단, u는 "균형잡힌 괄호 문자열"로 더 이상 분리할 수 없어야 하며, v는 빈 문자열이 될 수 있습니다.
        u, v = getUV(w)

        # 3. 문자열 u가 "올바른 괄호 문자열" 이라면 문자열 v에 대해 1단계부터 다시 수행합니다.
        if check(u):
            #   3-1. 수행한 결과 문자열을 u에 이어 붙인 후 반환합니다.
            return u + solution(v)

        # 4. 문자열 u가 "올바른 괄호 문자열"이 아니라면 아래 과정을 수행합니다.
        else:
            #   4-1. 빈 문자열에 첫 번째 문자로 '('를 붙입니다.
            #   4-2. 문자열 v에 대해 1단계부터 재귀적으로 수행한 결과 문자열을 이어 붙입니다.
            #   4-3. ')'를 다시 붙입니다.
            s = '(' + solution(v) + ')'
            #   4-4. u의 첫 번째와 마지막 문자를 제거하고, 나머지 문자열의 괄호 방향을 뒤집어서 뒤에 붙입니다.
            u = u[1:-1]
            for i in u:
                s += d[i]
            #   4-5. 생성된 문자열을 반환합니다.
            return s


def check(u):
    stack = []

    for i in u:
        stack.append(i)
        if len(stack) >= 2:
            if stack[-1] == ')' and stack[-2] == '(':
                stack = stack[:-2]

    if stack:
        return False
    else:
        return True


def getUV(w):
    b = {'(': 1, ')': -1}

    count = 0
    for i in range(len(w)):
        count += b[w[i]]
        if count == 0:
            break

    u = w[:i + 1]
    v = w[i + 1:]

    return u, v


print(solution("(()())()"))
print(solution(")("))
print(solution("()))((()"))
