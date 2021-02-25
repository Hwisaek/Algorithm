package acmicpc;

import java.util.Scanner;

public class Java_2941 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		String input = sc.next();

		String[] croatia = { "c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z=" };
		int result = 0;

		for (String alphabet : croatia) {
			input = input.replaceAll(alphabet, "*");
		}
		System.out.println(input.length());

	}

}
