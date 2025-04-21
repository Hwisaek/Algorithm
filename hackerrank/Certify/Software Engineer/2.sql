SELECT DISTINCT p.name
              , c.name
FROM schedule s
         LEFT JOIN professor p ON p.id = s.professor_id
         LEFT JOIN course c ON c.id = s.course_id
WHERE p.department_id != c.department_id
;
