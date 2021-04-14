package DFS_BFS;

import java.util.ArrayList;
import java.util.Scanner;

public class 음료수_얼려_먹기 {

	public static int[][] graph = new int[1000][1000];
	public static int result = 0;

	public static boolean dfs(int x, int y, int n, int m) {
		if (x <= -1 || x >= n || y <= -1 || y >= m) {
			return false;
		}
		if (graph[x][y] == 0) {
			graph[x][y] = 1;
			dfs(x - 1, y, n, m);
			dfs(x + 1, y, n, m);
			dfs(x, y - 1, n, m);
			dfs(x, y + 1, n, m);
			return true;
		}
		return false;
	}

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.printf("세로 길이 N과 가로 길이 M을 입력하세요>> ");
		int n = sc.nextInt(), m = sc.nextInt();
		sc.nextLine();
		
		System.out.println("얼음 틀의 형태를 입력하세요>> ");
		for (int i = 0; i < n; i++) {
			String str = sc.nextLine();
			for (int j = 0; j < m; j++) {
				graph[i][j] = str.charAt(j) - '0';
			}
		}
		
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < m; j++) {
				if (dfs(i, j, n, m)) {
					result++;
				}
			}
		}

		System.out.println("만들 수 있는 아이스크림의 개수: " + result);
	}

}
