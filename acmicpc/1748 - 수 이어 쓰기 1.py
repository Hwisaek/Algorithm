# 1   ~   9    | 9 * 1
# 10  ~   99   | 90 * 2
# 100 ~   999  | 900 * 3
# 1000~   9999 | 9000 * 4

n = input()
length = len(n)

count = 0

if length > 1:
    i = 0
    for i in range(1, length):
        count += 9 * (10 ** (i - 1)) * i

    n = int(n) % (10 ** i)
    n += 1
    count += n * (i + 1)
else:
    count = n

print(count)

# 9 * 1 + 90 * 2 + 900 * 3 + 9000 * 4
