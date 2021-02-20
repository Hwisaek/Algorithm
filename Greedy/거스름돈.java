package Greedy;

import java.util.Scanner;

public class 거스름돈 {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);

		System.out.printf("거스름돈을 입력해주세요>> ");
		int n = scan.nextInt();

		int[] coins = { 500, 100, 50, 10 };
		int cnt = 0;

		for (int coin : coins) {
			cnt += n / coin;
			n = n % coin;
		}
		System.out.println("거슬러줘야 할 동전의 최수 개수: " + cnt);
	}

}
