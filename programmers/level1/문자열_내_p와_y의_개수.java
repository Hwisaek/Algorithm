package level1;

public class 문자열_내_p와_y의_개수 {

	public static void main(String[] args) {
		System.out.println(new Solution().solution("Pyy"));
		System.out.println(new Solution().solution("pPoooyY"));
	}

}

class Solution {
	boolean solution(String s) {
		boolean answer = true;

		int n_p = 0;
		int n_y = 0;

		for (char ch : s.toUpperCase().toCharArray()) {
			if (ch == 'Y') {
				n_y++;
			} else if (ch == 'P') {
				n_p++;
			}

		}
		if (n_y != n_p) {
			answer = false;
		}
		System.out.println(n_y + ", " + n_p);
		return answer;
	}
}
