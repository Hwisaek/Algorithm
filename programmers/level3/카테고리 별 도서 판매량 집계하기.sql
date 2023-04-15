select category, sum(sales) total_sales
from book b
         join book_sales bs on b.book_id = bs.book_id
where sales_date between '2022-01-01' and '2022-01-31'
group by category
order by category