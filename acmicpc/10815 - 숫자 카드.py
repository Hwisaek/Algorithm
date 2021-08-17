n = input()
nlst = set(map(int, input().split()))
m = input()
mlst = list(map(int, input().split()))

for i in mlst:
    if i in nlst:
        print(1, end=" ")
    else:
        print(0, end=" ")
