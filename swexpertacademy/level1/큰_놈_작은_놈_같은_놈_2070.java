package level1;

import java.util.Scanner;

public class 큰_놈_작은_놈_같은_놈_2070 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int t = sc.nextInt();
		for (int i = 1; i <= 3; i++) {
			int a = sc.nextInt();
			int b = sc.nextInt();
			System.out.printf("#" + i + " ");
			if (a < b) {
				System.out.println("<");
			} else if (a > b) {
				System.out.println(">");
			} else {
				System.out.println("=");
			}
		}
	}
}
