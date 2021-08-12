n = list(map(int, input()))
l = len(n) // 2
print("LUCKY" if sum(n[:l]) == sum(n[l:]) else "READY")