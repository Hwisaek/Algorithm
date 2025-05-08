with A as (
    select host_team,
        guest_team,
        case
            when host_goals > guest_goals then 1
            when host_goals = guest_goals then 0
            when host_goals < guest_goals then -1
        end result
    from matches
)
select *
from (
        select b.team_id,
            b.team_name,
            coalesce(a.score, 0) score
        from teams b
            left join (
                select team,
                    sum(score) score
                from (
                        select host_team team,
                            sum(
                                case
                                    when result = 1 then 3
                                    when result = 0 then 1
                                    else 0
                                end
                            ) score
                        from A
                        group by host_team
                        union all
                        select guest_team team,
                            sum(
                                case
                                    when result = -1 then 3
                                    when result = 0 then 1
                                    else 0
                                end
                            ) score
                        from A
                        group by guest_team
                    ) a
                group by team
            ) a on a.team = b.team_id
    ) a
order by score desc,
    team_id