s = input()

n = {2: ['A', 'B', 'C'], 3: ['D', 'E', 'F'], 4: ['G', 'H', 'I'],
     5: ['J', 'K', 'L'], 6: ['M', 'N', 'O'], 7: ['P', 'Q', 'R', 'S'],
     8: ['T', 'U', 'V'], 9: ['W', 'X', 'Y', 'Z']}

t = 0

for i in s:
    for key, value in n.items():
        for j in value:
            if i == j:
                t += key + 1
print(t)
