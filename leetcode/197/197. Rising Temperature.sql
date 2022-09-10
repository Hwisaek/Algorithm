SELECT *
FROM (SELECT *, lag(temperature) over (ORDER BY recordDate) temperature_prev
      FROM weather
      ORDER BY recordDate) t
WHERE temperature > temperature_prev