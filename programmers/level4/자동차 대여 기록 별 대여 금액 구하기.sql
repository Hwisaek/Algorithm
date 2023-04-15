SELECT
    history_id
  , CEIL(
        (datediff (end_date, start_date) + 1) * daily_fee * (
            CASE
                WHEN datediff (end_date, start_date) + 1 >= 7 THEN 100 - (
                    SELECT
                        DISCOUNT_RATE
                    FROM
                        CAR_RENTAL_COMPANY_DISCOUNT_PLAN
                    WHERE
                        CAR_tYPE = '트럭'
                        AND DURATION_TYPE = '7일 이상'
                )
                WHEN datediff (end_date, start_date) + 1 >= 30 THEN 100 - (
                    SELECT
                        DISCOUNT_RATE
                    FROM
                        CAR_RENTAL_COMPANY_DISCOUNT_PLAN
                    WHERE
                        CAR_tYPE = '트럭'
                        AND DURATION_TYPE = '30일 이상'
                )
                WHEN datediff (end_date, start_date) + 1 >= 90 THEN 100 - (
                    SELECT
                        DISCOUNT_RATE
                    FROM
                        CAR_RENTAL_COMPANY_DISCOUNT_PLAN
                    WHERE
                        CAR_tYPE = '트럭'
                        AND DURATION_TYPE = '90일 이상'
                )
                ELSE 100
            END / 100
        )
    ) FEE
FROM
    CAR_RENTAL_COMPANY_RENTAL_HISTORY h
    JOIN CAR_RENTAL_COMPANY_CAR c ON h.car_id = c.car_id
WHERE
    car_Type = '트럭'
ORDER BY
    fee desc
  , history_id desc;


SELECT
    history_id
  , CEIL(
        CASE
            WHEN duration < 7 THEN duration * daily_fee
            WHEN duration < 30 THEN duration * daily_fee * (
                100 - (
                    SELECT
                        discount_rate
                    FROM
                        car_rental_company_discount_plan
                    WHERE
                        car_type = '트럭'
                        AND duration_type = '7일 이상'
                ) / 100
            )
            WHEN duration < 90 THEN duration * daily_fee * (
                100 - (
                    SELECT
                        discount_rate
                    FROM
                        car_rental_company_discount_plan
                    WHERE
                        car_type = '트럭'
                        AND duration_type = '30일 이상'
                ) / 100
            )
            ELSE duration * daily_fee * (
                100 - (
                    SELECT
                        discount_rate
                    FROM
                        car_rental_company_discount_plan
                    WHERE
                        car_type = '트럭'
                        AND duration_type = '90일 이상'
                ) / 100
            )
        END
    ) FEE
FROM
    (
        SELECT
            history_id
          , datediff (end_date, start_date) + 1 duration
          , DAILY_FEE
        FROM
            CAR_RENTAL_COMPANY_RENTAL_HISTORY h
            JOIN CAR_RENTAL_COMPANY_CAR c ON h.car_id = c.car_id
        WHERE
            car_Type = '트럭'
    ) t
ORDER BY
    fee desc
  , history_id desc;