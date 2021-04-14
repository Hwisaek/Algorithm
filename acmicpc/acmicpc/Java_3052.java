package acmicpc;

import java.util.Scanner;

public class Java_3052 {
	public static void main(String[] args) {

		Scanner sc = new Scanner(System.in);

		int[] arr = new int[10];
		int[] mod = new int[42];

		for (int i = 0; i < 10; i++) {
			arr[i] = sc.nextInt();
			mod[arr[i] % 42]++;
		}

		int result = 0;

		for (int i = 0; i < 42; i++) {
			if (mod[i] != 0) {
				result++;
			}
		}
		System.out.println(result);
	}
}
