package level3;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Iterator;

public class 가장_먼_노드 {
	public static int solution(int n, int[][] edge) {
		int answer = 0;
		ArrayList<HashSet<Integer>> list = new ArrayList<HashSet<Integer>>();
		for (int i = 0; i < n + 1; i++) {
			list.add(new HashSet<Integer>());
		}
		for (int[] arr : edge) {
			list.get(arr[0]).add(arr[1]);
		}

		for (int i = 0; i < list.size(); i++) {
			System.out.print(i + ":");
			Iterator<Integer> iter = list.get(i).iterator();
			while (iter.hasNext()) {
				System.out.print(iter.next() + " ");
			}
			System.out.println();
		}

		return answer;
	}

	public static void main(String[] args) {
		System.out.println(solution(6, new int[][] { { 3, 6 }, { 4, 3 }, { 3, 2 }, { 1, 3 }, { 1, 2 }, { 2, 4 }, { 5, 2 } }));
	}
}
