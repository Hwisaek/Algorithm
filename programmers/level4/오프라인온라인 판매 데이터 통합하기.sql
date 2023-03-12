select date_format(SALES_DATE, '%Y-%m-%d') sales_date, PRODUCT_ID, USER_ID, SALES_AMOUNT
from ONLINE_SALE
where sales_date between '2022-03-01' and '2022-03-31'
union all
select date_format(SALES_DATE, '%Y-%m-%d') sales_date, PRODUCT_ID, null USER_ID, SALES_AMOUNT
from OFFLINE_SALE
where sales_date between '2022-03-01' and '2022-03-31'
order by sales_date, product_id, user_id