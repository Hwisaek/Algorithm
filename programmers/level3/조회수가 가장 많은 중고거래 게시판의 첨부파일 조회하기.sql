SELECT
    concat (
        '/home/grep/src/'
      , a.board_id
      , '/'
      , file_id
      , file_name
      , file_ext
    )
FROM
    used_goods_board a
    LEFT JOIN used_goods_file b ON a.board_id = b.board_id
WHERE
    views = (
        SELECT
            MAX(views)
        FROM
            used_goods_board
    )