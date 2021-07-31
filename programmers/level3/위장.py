from functools import reduce


def solution(clothes):
    answer = 1  # 곱셈 연산을 위해 1로 초기화
    dic = {}  # 옷들을 딕셔너리로 담을 변수

    for name, type in clothes:  # 옷들을 dic에 담음
        dic[type] = dic.get(type, ["none"])  # 특정 부위에 아무것도 안 입을 경우도 있으므로 "none"을 기본값으로 넣어줌
        dic[type].append(name)

    for key in dic:  # 옷을 입는 조합의 수 = 옷 종류별 수들의 곱
        answer *= len(dic[key])

    return answer - 1  # 스파이는 하루에 최소 한 개의 의상은 입으므로 하나를 뺌


solution([["yellowhat", "headgear"], ["bluesunglasses", "eyewear"], ["green_turban", "headgear"]])
solution([["crowmask", "face"], ["bluesunglasses", "face"], ["smoky_makeup", "face"]])

r = reduce(lambda a, b: a * b, [1, 2, 3, 4, 5])
print(r)
