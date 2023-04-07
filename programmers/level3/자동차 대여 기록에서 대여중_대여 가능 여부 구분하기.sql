SELECT
    car_id
  , CASE
        WHEN EXISTS (
            SELECT
                CAR_ID
            FROM
                CAR_RENTAL_COMPANY_RENTAL_HISTORY b
            WHERE
                START_DATE <= '2022-10-16'
                AND END_DATE >= '2022-10-16'
                AND a.car_id = b.car_id
        ) THEN '대여중'
        ELSE '대여 가능'
    END AS availability
FROM
    (
        SELECT DISTINCT
            car_id
        FROM
            CAR_RENTAL_COMPANY_RENTAL_HISTORY
    ) a
ORDER BY
    car_id desc;