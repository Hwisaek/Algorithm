//https://programmers.co.kr/learn/courses/30/lessons/1829
package level2;

import java.util.*;

class 카카오프렌즈_컬러링북 {
	int numberOfArea = 0;
	int maxSizeOfOneArea = 0;
	int[] answer = new int[] { 0, 0 };

	boolean[][] visited;
	int[][] graph;

	int[] dx = new int[] { 1, -1, 0, 0 };
	int[] dy = new int[] { 0, 0, 1, -1 };

	int m, n;

	public int[] solution(int m, int n, int[][] picture) {
		visited = new boolean[m][n];
		graph = picture;
		this.m = m;
		this.n = n;

		for (int i = 0; i < m; i++) {
			for (int j = 0; j < n; j++) {
				if (!visited[i][j] && picture[i][j] != 0) {
					bfs(i, j);
				}
			}
		}

		return answer;
	}

	public void bfs(int x, int y) {
		int count = 0;
		Queue<int[]> queue = new LinkedList<int[]>();
		queue.add(new int[] { x, y });
		int color = graph[x][y];

		while (!queue.isEmpty()) {
			int[] now = queue.poll();

			for (int i = 0; i < 4; i++) {
				int nx = now[0] + dx[i];
				int ny = now[1] + dy[i];

				if (-1 < nx && nx < m && -1 < ny && ny < n) {
					if (!visited[nx][ny] && graph[nx][ny] == color) {
						queue.add(new int[] { nx, ny });
						visited[nx][ny] = true;
						count++;
					}
				}
			}

		}
		answer[0]++;
		answer[1] = answer[1] > count ? answer[1] : count;
		return;
	}

	public static void main(String[] args) {
		카카오프렌즈_컬러링북 s = new 카카오프렌즈_컬러링북();
		System.out.println(Arrays.toString(s.solution(6, 4, new int[][] { { 1, 1, 1, 0 }, { 1, 2, 2, 0 },
				{ 1, 0, 0, 1 }, { 0, 0, 0, 1 }, { 0, 0, 0, 3 }, { 0, 0, 0, 3 } })));
	}
}