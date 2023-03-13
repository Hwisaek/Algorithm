select truncate(price,-4) price_group, count(*) products
from product
group by price_group
order by price_group