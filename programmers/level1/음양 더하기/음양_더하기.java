package level1;

public class 음양_더하기 {
	public static int solution(int[] absolutes, boolean[] signs) {
		int answer = 0, size = signs.length;
		for (int i = 0; i < size; i++) {
			answer += absolutes[i] * (signs[i] ? 1 : -1);
		}
		return answer;
	}

	public static void main(String[] args) {
		System.out.println(solution(new int[] { 4, 7, 12 }, new boolean[] { true, false, true }));
	}
}
