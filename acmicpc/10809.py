s = input()

a = list(range(ord('a'), ord('z') + 1))

for i in a:
    print(s.find(chr(i)), end=' ')