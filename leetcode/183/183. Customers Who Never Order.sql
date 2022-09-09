SELECT name Customers
FROM customers
WHERE id NOT IN (SELECT DISTINCT customerId FROM orders)