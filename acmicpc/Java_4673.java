package acmicpc;

public class Java_4673 {
	public static void main(String[] args) {

		boolean[] nonSelfNumber = new boolean[10001];

		for (int i = 1; i < 10001; i++) {
			int n = i;
			int sum = 0;

			sum += n;
			while (n != 0) {
				sum += n % 10;
				n = n / 10;
			}
			if (sum < 10001)
				nonSelfNumber[sum] = true;
		}

		for (int i = 1; i < 10001; i++) {
			if (!nonSelfNumber[i])
				System.out.println(i);
		}
	}
}
