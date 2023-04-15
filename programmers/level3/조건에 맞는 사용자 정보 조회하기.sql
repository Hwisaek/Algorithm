SELECT
    user_id
  , nickname
  , concat (city, ' ', street_address1, ' ', street_address2)
  , concat (
        substr (tlno, 1, 3)
      , '-'
      , substr (tlno, 4, 4)
      , '-'
      , substr (tlno, 8, 4)
    )
FROM
    used_goods_user b
WHERE
    user_id IN (
        SELECT
            writer_id
        FROM
            used_goods_board a
        GROUP BY
            writer_id
        HAVING
            COUNT(*) >= 3
    )
ORDER BY
    user_id desc