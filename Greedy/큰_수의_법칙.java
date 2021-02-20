package Greedy;

import java.util.Arrays;
import java.util.Scanner;

public class 큰_수의_법칙 {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);

		System.out.printf("N, M, K를 입력해주세요>> ");
		int n = scan.nextInt();
		int m = scan.nextInt();
		int k = scan.nextInt();

		int[] arr = new int[n];
		System.out.printf("N개의 자연수를 입력해주세요>> ");
		for (int i = 0; i < n; i++) {
			arr[i] = scan.nextInt();
		}

		// 정렬
		Arrays.sort(arr);
		int first = arr[n - 1];
		int second = arr[n - 2];

		int rpt = first * 3 + second;

		int result = 0;
		result += rpt * m / (k + 1);
		result += first * m % (k + 1);
		System.out.println("큰 수의 법칙에 따른 결과: " + result);
	}

}
