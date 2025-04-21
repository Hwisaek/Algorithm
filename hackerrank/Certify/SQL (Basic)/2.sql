/*
Enter your query below.
Please append a semicolon ";" at the end of the query
*/
SELECT c.customer_id
     , c.name
     , '+' || cc.country_code || c.phone_number
FROM customers c
         LEFT JOIN country_codes cc ON cc.country = c.country
ORDER BY c.customer_id
;