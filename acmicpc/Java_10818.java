import java.util.Scanner;

public class Java_10818 {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);

		int n = scan.nextInt();
		int max = -1000000;
		int min = 1000000;
		for (int i = 0; i < n; i++) {
			int num = scan.nextInt();
			if (num < min)
				min = num;
			if (num > max)
				max = num;
		}
		System.out.printf("%d %d", min, max);
	}

}
