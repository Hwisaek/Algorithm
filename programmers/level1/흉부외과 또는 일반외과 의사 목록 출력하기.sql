select DR_NAME, DR_ID, MCDP_CD, date_format(hire_ymd, '%Y-%m-%d') as hire_ymd
from doctor
where mcdp_cd in ('CS', 'GS')
order by hire_ymd desc, dr_name