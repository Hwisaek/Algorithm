package level1;

public class _3진법_뒤집기 {
	public static int solution(int n) {
		StringBuilder sb = new StringBuilder();
		while (n > 0) {
			sb.append(n % 3);
			n /= 3;
		}
		return Integer.parseInt(sb.toString(), 3);
	}

	public static void main(String[] args) {
		System.out.println(solution(45));
		System.out.println(solution(125));
	}
}
