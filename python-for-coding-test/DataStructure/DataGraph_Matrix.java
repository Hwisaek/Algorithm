package dataStructure;

/**
 * 행렬 방식은 리스트 방식보다 메모리는 많이 먹으나 속도가 빠름.
 * @author Hwisaek
 *
 */
public class DataGraph_Matrix {
	
	public static final int INF = 999999999; // 무한의 비용 선언

	// 그래프(트리) 생성
	public static int[][] graph = { 
			{ 0,	7,		5 }, 
			{ 7,	0,		INF }, 
			{ 5,	INF,	0 } };

	public static void main(String[] args) {
		for (int i = 0; i < 3; i++) {
			for (int j = 0; j < 3; j++) {
				System.out.print(graph[i][j] + " ");
			}
			System.out.println();
		}
	}

}
