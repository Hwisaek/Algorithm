# https://softeer.ai/practice/info.do?eventIdx=1&psProblemId=407
k, p, n = map(int, input().split())

print((k * pow(p, n, 1000000007)) % 1000000007)
