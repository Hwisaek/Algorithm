def solution(phone_book):
    dic = {}
    for i in phone_book:
        dic[i] = False

    for num in phone_book:
        n = ""
        for i in num[:-1]:
            n += i
            if n in dic:
                return False

    return True


print(solution(["119", "97674223", "1195524421"]))
print(solution(["123", "456", "789"]))
print(solution(["12", "123", "1235", "567", "88"]))
