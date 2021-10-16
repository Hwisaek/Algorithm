k, p, n = map(int, input().split())

print((k * pow(p, n, 1000000007)) % 1000000007)
