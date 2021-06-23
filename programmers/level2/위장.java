package level2;

import java.util.HashMap;

public class 위장 {
	public static int solution(String[][] clothes) {
		int answer = 1;
		HashMap<String, Integer> map = new HashMap<String, Integer>();
		for (String[] str : clothes) {
			if (map.containsKey(str[1])) {
				map.put(str[1], map.get(str[1]) + 1);
			} else {
				map.put(str[1], 1);
			}
		}

		for (String key : map.keySet()) {
			answer *= map.get(key) + 1;
		}
		return answer - 1;
	}

	public static void main(String[] args) {
		System.out.println(solution(new String[][] { { "yellowhat", "headgear" }, { "bluesunglasses", "eyewear" }, { "green_turban", "headgear" } }));
	}
}
