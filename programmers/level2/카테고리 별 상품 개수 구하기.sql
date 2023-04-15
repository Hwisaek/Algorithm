SELECT
    substr(product_code, 1, 2),
    count(*)
FROM
    product
GROUP BY
    substr(product_code, 1, 2)
ORDER BY
    substr(product_code, 1, 2)