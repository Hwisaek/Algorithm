def solution(table, languages, preference):
    lst = []
    max = 0
    point = {}
    for i in range(len(languages)):
        point[languages[i]] = preference[i]

    for row in table:
        row = row.split()
        total = 0
        for i in range(-1, -len(row), -1):
            if row[i] in point:
                total += point[row[i]] * (-i)
        if total > max:
            lst = [row[0]]
            max = total
        elif total == max:
            lst.append(row[0])
    answer = min(lst)
    return answer


print(solution(["SI JAVA JAVASCRIPT SQL PYTHON C#", "CONTENTS JAVASCRIPT JAVA PYTHON SQL C++",
                "HARDWARE C C++ PYTHON JAVA JAVASCRIPT", "PORTAL JAVA JAVASCRIPT PYTHON KOTLIN PHP",
                "GAME C++ C# JAVASCRIPT C JAVA"], ["PYTHON", "C++", "SQL"], [7, 5, 5]))
print(solution(["SI JAVA JAVASCRIPT SQL PYTHON C#", "CONTENTS JAVASCRIPT JAVA PYTHON SQL C++",
                "HARDWARE C C++ PYTHON JAVA JAVASCRIPT", "PORTAL JAVA JAVASCRIPT PYTHON KOTLIN PHP",
                "GAME C++ C# JAVASCRIPT C JAVA"], ["JAVA", "JAVASCRIPT"], [7, 5]))
