package acmicpc;

import java.util.Scanner;

public class Java_1546 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();
		double score[] = new double[n];
		double max = 0;
		for (int i = 0; i < n; i++) {
			score[i] = sc.nextDouble();
			if (score[i] > max) {
				max = score[i];
			}
		}
		double total = 0;
		for (int i = 0; i < n; i++) {
			score[i] = score[i] / max * 100;
			total += score[i];
		}

		double result = total / n;
		System.out.println(result);
	}

}
