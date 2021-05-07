package level1;
import java.util.Scanner;

public class 평균값_구하기_2071 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		// 테스트 케이스의 개수 t
		int T;
		T = sc.nextInt();
		for (int test_case = 1; test_case <= T; test_case++) {
			int[] num = new int[10];
			int sum = 0;
			for (int j = 0; j < 10; j++) {
				int n = sc.nextInt();
				num[j] = n;
				sum += n;
			}
			System.out.println("#" + test_case + " " + Math.round(sum / 10.0));
		}
	}

}
