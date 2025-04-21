/*
Enter your query below.
Please append a semicolon ";" at the end of the query
*/
SELECT si.roll_number
     , si.name
FROM student_information si
         LEFT JOIN faculty_information fi ON fi.employee_id = si.advisor
WHERE (fi.gender = 'M' AND fi.salary > 15000)
   OR (fi.gender = 'F' AND fi.salary > 20000)
;