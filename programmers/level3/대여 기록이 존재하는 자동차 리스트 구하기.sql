SELECT distinct
    a.car_id
FROM
    CAR_RENTAL_COMPANY_CAR a
    LEFT JOIN CAR_RENTAL_COMPANY_RENTAL_HISTORY b ON a.car_id = b.car_id
where
    a.car_Type = '세단'
    and b.start_date >= '2022-10-01'
order by
    car_id desc