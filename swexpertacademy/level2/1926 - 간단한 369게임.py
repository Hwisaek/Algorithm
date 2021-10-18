n = int(input())

for i in range(1, n + 1):
    clap = str(i).count('3') + str(i).count('6') + str(i).count('9')
    print('-' * clap if clap > 0 else i, end=' ')
