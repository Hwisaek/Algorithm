package level2;

import java.util.Arrays;

public class 스킬트리 {

	public static void main(String[] args) {
//		new Solution().solution("CBD", new String[] { "BACDE", "CBADF", "AECB", "BDA" });
		new Solution().solution("CZ", new String[] { "BACDE" });
	}

}

class Solution {
	int[] order;

	public int solution(String skill, String[] skill_trees) {
		int answer = 0;

		int[] order = new int[skill.length()];

		for (String skill_tree : skill_trees) { // 모든 스킬트리 반복
			Arrays.fill(order, 26);
			for (int i = 0; i < skill.length(); i++) {
				order[i] = skill_tree.indexOf(skill.charAt(i));
				if (order[i] == -1) {
					order[i] = 26;
				}
				if (i > 0) { // 1번째 인덱스부터
					if (order[i] < order[i - 1]) { // 스킬트리 순서가 꼬이면
						break;
					}
				}
				if (i == skill.length() - 1) {
					answer++;
				}
			}
		}
		System.out.println(answer);
		return answer;
	}
}