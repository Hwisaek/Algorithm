package acmicpc;

import java.util.Scanner;

public class Java_2577 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int a = sc.nextInt(), b = sc.nextInt(), c = sc.nextInt();
		String num = Integer.toString(a * b * c);
		int[] cnt = new int[10];

		for (int i = 0; i < num.length(); i++) {
			cnt[Integer.parseInt(String.valueOf(num.charAt(i)))]++;
		}

		for (int i = 0; i < 10; i++) {
			System.out.println(cnt[i]);
		}
	}

}
