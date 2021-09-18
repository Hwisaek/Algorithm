a = int(input())
b = int(input())
c = int(input())

if a + b + c == 180:
    if a == b == c:  # 세 각 모두 60
        print("Equilateral")
    elif a == b or b == c or c == a:  # 두 각이 같은 경우
        print("Isosceles")
    else:  # 같은 각이 없는 경우
        print("Scalene")
else:  # 세 각의 합이 180이 아닌 경우
    print("Error")
