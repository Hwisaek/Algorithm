x = int(input())

line = 1

i = 1
j = 1

while x > line:
    x -= line
    line += 1

if line % 2 == 0:
    i = x
    j = line - x + 1
else:
    i = line - x + 1
    j = x

print('%d/%d' % (i, j))
