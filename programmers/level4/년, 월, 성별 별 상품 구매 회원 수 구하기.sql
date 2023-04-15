SELECT    y
        , m
        , gender
        , COUNT(1)
FROM      (
          SELECT    YEAR (sales_date)  y
                  , MONTH (sales_date) m
                  , user_id
          FROM      online_sale os
          GROUP BY  YEAR (sales_date)
                  , MONTH (sales_date)
                  , user_id
          ) os
JOIN      user_info ui ON os.user_id = ui.user_id
WHERE     gender IS NOT NULL
GROUP BY  y
        , m
        , gender
ORDER BY  y
        , m
        , gender