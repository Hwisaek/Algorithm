import java.util.Scanner;

public class Java_11653 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();

		while (n != 1) {
			if (n % 2 == 0) {
				System.out.println(2);
				n /= 2;
				continue;
			}
			for (int i = 3; i <= n; i++) {
				if (i % 2 == 0) {
					continue;
				}
				if (n % i == 0) {
					System.out.println(i);
					n /= i;
					break;
				}
			}
		}

	}

}
