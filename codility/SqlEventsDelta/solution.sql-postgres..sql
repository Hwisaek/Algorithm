with a as (
    SELECT *
    FROM (
            SELECT ROW_NUMBER() OVER (
                    PARTITION BY event_type
                    ORDER BY time desc
                ) rn,
                *
            FROM events
            order by event_type,
                time DESC
        ) rn
    where rn <= 2
)
select a1.event_type,
    a1.value - a2.value
from a a1
    join a a2 on a1.rn = 1
    and a2.rn = 2
    and a1.event_type = a2.event_type
where a1.rn = 1