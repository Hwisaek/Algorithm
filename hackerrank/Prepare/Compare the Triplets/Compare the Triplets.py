import math
import os
import random
import re
import sys


def compareTriplets(a, b):
    count_a = 0
    count_b = 0

    for x, y in zip(a, b):
        if x > y:
            count_a += 1
        elif x < y:
            count_b += 1

    return count_a, count_b


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    a = list(map(int, input().rstrip().split()))

    b = list(map(int, input().rstrip().split()))

    result = compareTriplets(a, b)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
