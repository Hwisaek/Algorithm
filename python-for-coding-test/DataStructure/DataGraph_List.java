package DataStructure;

import java.util.ArrayList;

/**
 * 리스트 방식은 행렬 방식보다 메모리는 적게 먹으나 속도가 느림.
 * 
 * 자신과 연결된 노드의 정보만 가지고 있으므로 모든 인접 노드를 순회해야 하는 경우 행렬 방식보다 메모리를 덜 사용함.
 * 
 * @author Hwisaek
 *
 */

class Node { // 노드 클래스

	private int index;
	private int distance;

	public Node(int index, int distance) { // 생성자
		this.index = index;
		this.distance = distance;
	}

	public void show() {
		System.out.printf("(" + this.index + "," + this.distance + ") ");
	}
}

public class DataGraph_List {

	// ArrayList<Node>를 데이터로 가지는 ArrayList 생성
	// ArrayList 내부의 ArrayList 는 Node로 이루어져있음.
	// 각 노드에는 연결된 노드와의 비용이 저장되어 있음.
	public static ArrayList<ArrayList<Node>> graph = new ArrayList<ArrayList<Node>>();

	public static void main(String[] args) {
		for (int i = 0; i < 3; i++) {
			graph.add(new ArrayList<Node>()); // 그래프에 노드 1, 2, 3 추가
		}

		graph.get(0).add(new Node(1, 7)); // 노드 0 에 비용 7의 노드 1 추가
		graph.get(0).add(new Node(2, 5)); // 노드 0 에 비용 5의 노드 2 추가

		graph.get(1).add(new Node(0, 7)); // 노드 1 에 비용 7의 노드 0 추가

		graph.get(2).add(new Node(0, 5)); // 노드 2 에 비용 5의 노드 0 추가

		for (int i = 0; i < 3; i++) {
			for (int j = 0; j < graph.get(i).size(); j++) {
				graph.get(i).get(j).show();
			}
			System.out.println();
		}
	}

}
