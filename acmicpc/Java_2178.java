package acmicpc;

import java.util.Scanner;

public class Java_2178 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);

		int n = sc.nextInt(); // За
		int m = sc.nextInt(); // ї­

		int[][] arr = new int[n][m];
		boolean[][] visited = new boolean[n][m];
		
		for (int i = 0; i < n; i++) {
			String str = sc.next();
			char[] chaArr = str.toCharArray();
			for (int j = 0; j < m; j++) {
				arr[i][j] = chaArr[j] - '0';
			}
		}

	}
}
