package level1;

import java.util.Arrays;

public class 로또의_최고_순위와_최저_순위 {

	public static int[] solution(int[] lottos, int[] win_nums) {
		int[] rank = { 6, 6, 5, 4, 3, 2, 1 };
		int count = 0, min = 0;
		for (int i : lottos) {
			if (i == 0) {
				count++;
			} else {
				for (int j : win_nums) {
					if (i == j) {
						min++;
					}
				}
			}
		}
		return new int[] { rank[min + count], rank[min] };
	}

	public static void main(String[] args) {
		System.out.println(Arrays.toString(solution(new int[] { 44, 1, 0, 0, 31, 25 }, new int[] { 31, 10, 45, 1, 6, 19 })));
	}
}
