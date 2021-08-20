def dfs(x, y):
    if x < 0 or x >= n or y < 0 or y >= m:
        return False
    if frame[x][y] == 0:
        frame[x][y] = 1
        dfs(x + 1, y)
        dfs(x - 1, y)
        dfs(x, y + 1)
        dfs(x, y - 1)
        return True
    return False
