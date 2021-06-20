package level1;

import java.util.HashSet;

public class 폰켓몬 {

	public static void main(String[] args) {
		int[] nums = { 3, 1, 2, 3 };

		System.out.println(solution(nums));
	}

	public static int solution(int[] nums) {
		int answer = 0;
		int max = nums.length / 2;
		HashSet<Integer> set = new HashSet<Integer>();
		for (int n : nums) {
			set.add(n);
		}
		answer = max > set.size() ? set.size() : max;
		return answer;
	}
}
