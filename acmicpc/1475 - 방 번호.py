from collections import Counter

N = dict(Counter(input()))
N['6'] = (N.get('6', 0) + N.get('9', 0) + 1) // 2
N['9'] = 0

print(max(N.values()))
