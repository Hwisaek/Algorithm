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
    if count == 0:
        print("YES")
    else:
        print("NO")
