n = input()
nlst = list(map(int, input().split()))
m = input()
mlst = list(map(int, input().split()))

for i in mlst:
    if nlst.count(i) > 0:
        print(1, end=" ")
    else:
        print(0, end=" ")
