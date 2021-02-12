self_num = set(range(1, 10001))
def_num = set()

for i in range(1, 10001):
    for j in str(i):
        i += int(j)
    def_num.add(i)

for i in sorted(self_num - def_num):
    print(i)