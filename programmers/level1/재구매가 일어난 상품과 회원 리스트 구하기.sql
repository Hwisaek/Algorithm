select user_id, product_id
from online_sale os
group by user_id, product_id
having count(1) > 1
order by user_id, product_id desc