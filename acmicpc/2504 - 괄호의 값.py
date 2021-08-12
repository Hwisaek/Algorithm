import sys

s = input()

dic = {"(": ")", ")": "(", "[": "]", "]": "["}  # 괄호 쌍
dicN = {"(": "2", ")": "2", "[": "3", "]": "3"}  # 괄호-숫자 쌍
wrong = {"(": "]", ")": "[", "[": ")", "]": "("}  # 괄호-숫자 쌍

stack = []


def add(c):
    global stack
    if len(stack) > 0:
        if c.isdigit():  # 숫자인 경우
            if stack[-1].isdigit():  # 앞 자리도 숫자인 경우
                return add(str(int(stack.pop()) + int(c)))
            else:
                return c
        else:  # 괄호인 경우
            if c == "(" or c == "[":  # 여는 괄호면 바로 반환
                return c
            else:  # 닫는 괄호인 경우
                if stack[-1].isdigit():  # 앞 자리가 숫자인 경우 곱셈
                    a = str(int(stack.pop()) * int(dicN[c]))
                    if stack and stack[-1] == dic[c]:
                        stack.pop()
                    else:
                        print(0)
                        sys.exit(0)
                    return add(a)
                else:  # 앞 자리가 괄호인 경우
                    if stack[-1] == wrong[c]:  # 앞 자리가 잘못된 괄호인 경우
                        print(0)
                        sys.exit(0)
                    else:  # 앞 자리가 올바른 괄호인 경우
                        stack.pop()
                        return add(dicN[c])
    elif c == "]" or c == ")":
        print(0)
        sys.exit(0)
    else:
        return c


if len(s) == 1 or len(s) % 2 != 0:
    print(0)
    sys.exit(0)

for c in s:
    stack.append(add(c))

if stack[0].isdigit():
    print(stack[0])
else:
    print(0)
