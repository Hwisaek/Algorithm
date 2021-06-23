import java.util.Scanner;

public class Java_2292 {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();

		int i = 0;
		while (true) {
			if (n == 1) {
				i = 1;
				break;
			}
			int a = 3 * (int) Math.pow(i - 1, 2) - 3 * (i - 1) + 1;
			int b = 3 * (int) Math.pow(i, 2) - 3 * i + 1;
			if (a < n && n <= b) {
				break;
			}
			i++;
		}
		System.out.println(i);
	}
}
