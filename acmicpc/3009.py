arr = [[], []]

for _ in range(3):
    a = list(map(int, input().split()))
    for i in range(2):
        arr[i].append(a[i])

for i in range(2):
    arr_dict = {}
    for j in range(3):
        arr_dict[arr[i][j]] = arr[i].count(arr[i][j])
    for key, value in arr_dict.items():
        if value == 1:
            print(key, end=' ')