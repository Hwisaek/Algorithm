select substr(product_code,1,2),count(*) from product
group by substr(product_code,1,2)
order by substr(product_code,1,2)