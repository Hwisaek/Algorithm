import sys
from collections import Counter

input = sys.stdin.readline

n = int(input())
numbers = [int(input()) for _ in range(n)]
count = Counter(numbers)
maximum = max(count.values())
result = list(map(lambda x: x[0], filter(lambda x: x[1] == maximum, count.items())))
print(sorted(result)[0])
