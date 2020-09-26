def solution(n, lost, reserve):
    answer = n
    com = set(lost) & set(reserve)
    lost = list(set(lost) - com)
    reserve = list(set(reserve) - com)
    answer -= len(lost)
    
    for i in lost:
        for j in reserve:
            if i + 1 == j:
                answer += 1
                reserve.remove(j)
                break
            elif i - 1 == j:
                answer += 1
                reserve.remove(j)
                break
                
    return answer


# 테스트 케이서 생성기

def random_case_generator(n):
    import random
    lost = random.sample(range(1, n), random.randint(0,n-1))
    reserve = random.sample(range(1, n), random.randint(0,n-1))
    arr = [1] * n
    for l in lost:
        arr[l - 1] -= 1
    for r in reserve:
        arr[r - 1] += 1

    for i in range(n - 1):
        if arr[i] == 0 and arr[i + 1] == 2:
            arr[i] = 1
            arr[i + 1] = 1
        if arr[i] == 2 and arr[i + 1] == 0:
            arr[i] = 1
            arr[i + 1] = 1
    answer = n - arr.count(0)
    return n, lost, reserve, answer

n = 3
n, lost, reserve, answer = random_case_generator(n)
while True:
    *par, answer = random_case_generator(n)
    if solution(*par) != answer:
        print(f"incorrect for case:{par}")
        print('answer: {}, solution: {}'.format(answer, solution(*par)))
        break