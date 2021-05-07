package level1;

import java.util.Arrays;

public class 소수_찾기 {
	public static void main(String[] args) {
		System.out.println(", 개수: " + Solution.solution(10));
		System.out.println(", 개수: " + Solution.solution(5));
	}

	static class Solution {
		public static int solution(int n) {
			int answer = 0;

			boolean[] chk = new boolean[n + 1];
			Arrays.fill(chk, true);
			chk[0] = false;
			chk[1] = false;
			System.out.print(n + "이하의 소수: ");
			for (int i = 2; i <= n; i++) {
				if (chk[i]) {
					System.out.print(i + " ");
					for (int j = 2; i * j <= n; j++) {
						chk[i * j] = false;
					}
					answer++;
				}
			}
			return answer;
		}
	}
}
