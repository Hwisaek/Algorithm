package DataStructure;

import java.util.LinkedList;
import java.util.Queue;

public class DataQueue {

	public static void main(String[] args) {
		// 큐 생성
		Queue<Integer> q = new LinkedList();

		// 삽입(5) - 삽입(2) - 삽입(3) - 삽입(7) - 삭제() - 삽입(1) - 삽입(4) - 삭제()
		q.offer(5); // 삽입(5)
		q.offer(2); // 삽입(2)
		q.offer(3); // 삽입(3)
		q.offer(7); // 삽입(7)
		q.poll(); // 삭제()
		q.offer(1); // 삽입(1)
		q.offer(4); // 삽입(4)
		q.poll(); // 삭제()

		// 큐 구조: 4 - 1 - 7 - 3
		while (!q.isEmpty()) {
			System.out.println(q.poll()); // Queue.poll()은 큐의 맨 앞 자료를 반환하고 삭제함
		}
	}

}
