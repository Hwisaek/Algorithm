def solution(n):
    answer = ''
    
    while True:
        if n == 0:
            break
        m = n % 3
        n = n // 3
        
        if m == 1:
            answer = 'a' + answer
        elif m == 2:
            answer = 'b' + answer
        else:
            answer = 'c' + answer
    answer.replace('a', '1')
    answer.replace('b', '2')
    answer.replace('c', '4')
    return answer

for i in range(1, 10):
    print(solution(i))