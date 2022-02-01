def solution(phone_book):
    # List to set
    dic = {i for i in phone_book}

    # 전화번호 한개씩 꺼냄
    for num in phone_book:
        n = ""
        # 한글자씩 추가해가며 dic에 있는지 확인
        for i in num[:-1]:
            n += i
            # dic에 존재하면 False 반환
            if n in dic:
                return False

    return True


print(solution(["119", "97674223", "1195524421"]))
print(solution(["123", "456", "789"]))
print(solution(["12", "123", "1235", "567", "88"]))
