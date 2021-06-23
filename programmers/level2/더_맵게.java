package level2;

import java.util.*;

public class 더_맵게 {
	public static int solution(int[] scoville, int K) {
		int answer = 0;
		PriorityQueue<Integer> queue = new PriorityQueue<Integer>();
		for (int i : scoville) {
			queue.add(i);
		}

		while (queue.peek() < K) {
			if(queue.size()==1) {
				return -1;
			}
			int n = queue.poll();
			n += queue.poll() * 2;
			queue.add(n);
			answer++;
		}
		return answer;
	}

	public static void main(String[] args) {
		System.out.println(solution(new int[] { 1, 2, 3, 9, 10, 12 }, 7));
	}
}
