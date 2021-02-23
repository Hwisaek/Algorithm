package acmicpc;

import java.util.Scanner;

public class Java_1157 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);

		String str = sc.next();
		str = str.toUpperCase();
		int[] cnt = new int[26];
		int max = 0, max_n = -1, index = 0;

		for (int i = 0; i < str.length(); i++) {
			cnt[str.charAt(i) - 'A']++;
		}

		for (int i = 0; i < 26; i++) {
			if (cnt[i] > max) {
				max = cnt[i];
				index = i;
				max_n = 1;
			} else if (cnt[i] == max) {
				max_n++;
			}
		}

		if (max_n > 1) {
			System.out.println("?");
		} else {
			System.out.println((char) (index + 'A'));
		}
	}

}
