package level1;

public class 소수_찾기 {
	public static void main(String[] args) {
		System.out.println(solution(10));
	}

	public static int solution(int n) {
		int count = n - 1;
		boolean[] isPrime = new boolean[n + 1];
		int max = (int) Math.ceil(Math.sqrt(n));
		for (int i = 2; i <= max; i++) {
			if (!isPrime[i]) {
				for (int j = 2; j * i <= n; j++) {
					if (!isPrime[j * i]) {
						isPrime[j * i] = true;
						count--;
					}
				}
			}
		}
		return count;
	}
}
