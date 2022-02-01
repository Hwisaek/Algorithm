package level2;

import java.util.HashMap;

public class 전화번호_목록 {
	public static boolean solution(String[] phone_book) {
		boolean answer = true;
		HashMap<String, Object> map = new HashMap<String, Object>();
		for (String str : phone_book) {
			map.put(str, null);
		}
		for (String str : phone_book) {
			for (int i = 0; i < str.length(); i++) {
				if (map.containsKey(str.substring(0, i))) {
					return false;
				}
			}
		}
		return answer;
	}

	public static void main(String[] args) {
		System.out.println(solution(new String[] { "119", "97674223", "1195524421" }));
		System.out.println(solution(new String[] { "123", "456", "789" }));
		System.out.println(solution(new String[] { "12", "123", "1235", "567", "88" }));
		System.out.println(solution(new String[] { "1195524421", "97674223", "119" }));
		System.out.println(solution(new String[] { "113", "44", "4544" }));
		System.out.println(solution(new String[] { "010111", "010" }));
	}
}
