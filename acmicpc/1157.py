s = input().upper()

n = []

for i in set(s):
    n.append(s.count(i))

if n.count(max(n)) != 1:
    print('?')
else:
    print(list(set(s))[n.index(max(n))])