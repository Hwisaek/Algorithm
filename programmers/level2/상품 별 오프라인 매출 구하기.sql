select product_code, sum(price*sales_amount)
from product p join offline_sale os on p.product_id = os.product_id
group by product_code
order by sum(price*sales_amount) desc, product_code