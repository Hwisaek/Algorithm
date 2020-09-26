def solution(d, budget):
    answer = 0
    
    for i in range(len(d)):
        if budget >= min(d):
            budget -= min(d)
            d.remove(min(d))
            answer += 1
        else:
            break
    return answer