select car_id,
       case (SELECT count(*)
             from CAR_RENTAL_COMPANY_RENTAL_HISTORY
             where '2022-10-16' between start_date and end_date
               and car_id = crcrh.car_id)
           when 0 then '대여중'
           else '대여 가능' end availability
from CAR_RENTAL_COMPANY_RENTAL_HISTORY crcrh
group by car_id
order by car_id desc