package level2;

import java.util.LinkedList;
import java.util.Queue;

public class 프린터 {

	public static int solution(int[] priorities, int location) {
		int answer = 0;
		int[] count = new int[10]; // 우선순위 별 개수를 넣음
		int max = 0;
		Queue<Integer> queue = new LinkedList<Integer>();

		// 큐에 데이터를 순서대로 넣음
		for (int i : priorities) {
			max = max < i ? i : max;
			count[i]++;
			queue.add(i);
		}

		while (true) {
			int n = queue.poll();
			location--;
			if (n == max) {
				answer++;
				if (location == -1) {
					break;
				}
				count[max]--;
				for (int i = max; i > 0; i--) {
					if (count[i] > 0) {
						max = i;
						break;
					}
				}
			} else {
				if (location == -1) {
					location = queue.size();
				}
				queue.add(n);
			}
		}

		return answer;
	}

	public static void main(String[] args) {
//		System.out.println(solution(new int[] { 2, 1, 3, 2 }, 2));
		System.out.println(solution(new int[] { 1, 1, 9, 1, 1, 1 }, 0));
	}
}
