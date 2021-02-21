package DFS_BFS;

import java.util.*;

public class BFS {

	public static boolean[] visited = new boolean[9];
	public static ArrayList<ArrayList<Integer>> graph = new ArrayList<ArrayList<Integer>>();

	public static void bfs(int start) {
		Queue<Integer> q = new LinkedList<>();
		q.offer(start);
		visited[start] = true; // 시작점 방문

		while (!q.isEmpty()) { // Queue가 빌 때까지 반복
			int x = q.poll(); // 큐의 맨 앞 제거 후 저장
			System.out.print(x + " ");
			for (int i = 0; i < graph.get(x).size(); i++) { // x와 연결된 노드들 탐색
				int y = graph.get(x).get(i);
				if (!visited[y]) { // 방문하지 않은 노드이면
					q.offer(y); // 큐에 해당 노드를 삽입
					visited[y] = true; // 방문 표시 
				} // x와 연결된 노드들 방문 후 다음 노드로 넘어감
			}
		}
	}

	public static void main(String[] args) {
		for (int i = 0; i < 9; i++) {
			graph.add(new ArrayList<Integer>());
		}

		graph.get(1).add(2);
		graph.get(1).add(3);
		graph.get(1).add(8);

		graph.get(2).add(1);
		graph.get(2).add(7);

		graph.get(3).add(1);
		graph.get(3).add(4);
		graph.get(3).add(5);

		graph.get(4).add(3);
		graph.get(4).add(5);

		graph.get(5).add(3);
		graph.get(5).add(4);

		graph.get(6).add(7);

		graph.get(7).add(2);
		graph.get(7).add(6);
		graph.get(7).add(8);

		graph.get(8).add(1);
		graph.get(8).add(7);

		bfs(1);
	}

}
