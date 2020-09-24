n = int(input())

c = 0

b = 0

for i in range(n):
    s = list(input())
    for j in range(1, len(s)):
        if s[j] == s[j-1]:
            continue
        else:
            if s[j] in s[:j]:
                b = 1
                break
    if b == 1:
        b = 0
        continue
    else:
        c += 1
print(c)