SELECT
       user_id
     , nickname
     , SUM(price) total_sales
FROM
       used_goods_user ugu
       LEFT JOIN used_goods_board ugd ON ugu.user_id = ugd.writer_id
WHERE
       STATUS = 'DONE'
GROUP BY
       writer_id
HAVING
       SUM(price) >= 700000
ORDER BY
       total_Sales