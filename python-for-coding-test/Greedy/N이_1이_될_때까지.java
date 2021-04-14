package greedy;

import java.util.Scanner;

public class N이_1이_될_때까지 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.printf("N, K를 입력하세요>> ");

		int n = sc.nextInt();
		int k = sc.nextInt();

		int cnt = 0;

		while (n != 1) {
			if (n % k == 0)
				n /= k;
			else
				n -= 1;
			cnt++;
		}

		System.out.println("N을 1로 만드는 최소 횟수: " + cnt);
	}

}
