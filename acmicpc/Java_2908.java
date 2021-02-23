package acmicpc;

import java.util.Scanner;

public class Java_2908 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);

		String a = sc.next(), b = sc.next();

		a = new StringBuffer(a).reverse().toString();
		b = new StringBuffer(b).reverse().toString();
		int num_a = Integer.parseInt(a), num_b = Integer.parseInt(b);

		if (num_a > num_b) {
			System.out.println(a);

		} else {
			System.out.println(b);
		}
	}

}
