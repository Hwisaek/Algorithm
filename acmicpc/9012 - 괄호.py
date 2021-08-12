t = int(input())

for _ in range(t):
    vps = input()
    count = 0
    for ps in vps:
        if ps == '(':
            count += 1
        else:
            count -= 1

        if count < 0:
            break
    if count < 0 or vps[-1] == "(" or count != 0:
        print("NO")
    else:
        print("YES")
