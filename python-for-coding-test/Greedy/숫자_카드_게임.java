package greedy;

import java.util.Arrays;
import java.util.Scanner;

public class 숫자_카드_게임 {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);

		System.out.printf("N, M 입력>> ");
		int n = scan.nextInt();
		int m = scan.nextInt();

		int max = 0;
		int[][] arr = new int[n][m];
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < m; j++) {
				arr[i][j] = scan.nextInt();
			}
			Arrays.sort(arr[i]);
			if (arr[i][0] > max)
				max = arr[i][0];
		}

		System.out.println("각 행의 최소값중 최대값: " + max);
	}

}
