package level1;

import java.util.Arrays;
import java.util.HashSet;

public class 두_개_뽑아서_더하기 {

	static class Solution {
		public static int[] solution(int[] numbers) {
			int[] answer;
			HashSet<Integer> s = new HashSet<Integer>();

			for (int i = 0; i < numbers.length - 1; i++) {
				int in = numbers[i];
				for (int j = i + 1; j < numbers.length; j++) {
					s.add(in + numbers[j]);
				}
			}
			answer = new int[s.size()];
			Object[] arr = s.toArray();
			for (int i = 0; i < s.size(); i++) {
				answer[i] = (int) arr[i];
			}
			return answer;
		}

	}

	public static void main(String[] args) {

		System.out.println(Arrays.toString(Solution.solution(new int[] { 2, 1, 3, 4, 1 })));
		System.out.println(Arrays.toString(Solution.solution(new int[] { 5, 0, 2, 7 })));

	}

}
