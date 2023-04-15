SELECT mcdp_cd,count(*)
from appointment
where apnt_ymd between '2022-05-01' and '2022-05-31'
group by mcdp_cd
order by count(*),mcdp_cd