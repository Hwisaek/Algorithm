SELECT    t.author_id
        , a.author_name
        , t.category
        , t.total_sales
FROM      (
          SELECT    author_id
                  , category
                  , SUM(price * sales) total_sales
          FROM      book_sales bs
          LEFT JOIN book b ON bs.book_id = b.book_id
          WHERE     sales_date BETWEEN '2022-01-01' AND '2022-01-31'
          GROUP BY  author_id
                  , category
          ) t
LEFT JOIN author a ON t.author_id = a.author_id
ORDER BY  author_id
        , category desc