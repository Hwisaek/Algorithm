n = int(input())

array = [list(input()) for _ in range(n)]
visited = [[False] * n for _ in range(n)]

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]

group = 0
counts = []
for a in range(n):
    for b in range(n):

        count = 0


        def dfs(x, y):
            global array
            global count
            global visited

            if visited[x][y]:
                return

            visited[x][y] = True
            if array[x][y] == '0':
                return

            count += 1

            for i in range(4):
                nx = x + dx[i]
                ny = y + dy[i]
                if -1 < nx < n and -1 < ny < n and not visited[nx][ny] and array[nx][ny] == '1':
                    dfs(nx, ny)


        dfs(a, b)
        if count != 0:
            group += 1
            counts.append(count)

print(group)

for count in sorted(counts):
    print(count)
