prices = [1, 2, 3, 2, 3]
answer = []
n = len(prices)
print('prices 길이: ', n)

for i in range(n):
    index = 0
    for j in range(i + 1, n):
        if prices[i] <= prices[j]:
            index += 1
    answer.append(index)
