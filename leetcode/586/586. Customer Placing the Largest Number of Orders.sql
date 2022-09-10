SELECT customer_number
FROM (SELECT customer_number,
             dense_rank() over(ORDER BY COUNT(*) DESC) dr
      FROM orders
      GROUP BY customer_number) t
WHERE dr = 1