DELETE
FROM person
WHERE id IN (SELECT id
             FROM (SELECT *, rank() over(partition BY email ORDER BY id) r
                   FROM person) t
             WHERE r != 1)