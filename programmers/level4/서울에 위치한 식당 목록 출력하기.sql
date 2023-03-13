select ri.rest_id, rest_name, food_type, favorites, address, score
from rest_info ri
         join (select rest_id, round(avg(review_score), 2) score
               from rest_review
               group by rest_id) rr on ri.rest_id = rr.rest_id
where address like '서울%'
order by score desc, favorites desc