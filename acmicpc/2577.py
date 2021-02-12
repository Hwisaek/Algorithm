arr = []
for i in range(3):
    arr.append(int(input()))

num = list(str(arr[0] * arr[1] * arr[2]))

for i in range(10):
    print(num.count('%d' % i))
