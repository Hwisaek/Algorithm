package level2;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedList;

public class 순위_검색 {

	public static void main(String[] args) {
		String[] info = { "java backend junior pizza 150", "python frontend senior chicken 210",
				"python frontend senior chicken 150", "cpp backend senior pizza 260", "java backend junior chicken 80",
				"python backend senior chicken 50" };

		String[] query = { "java and backend and junior and pizza 100",
				"python and frontend and senior and chicken 200", "cpp and - and senior and pizza 250",
				"- and backend and senior and - 150", "- and - and - and chicken 100", "- and - and - and - 150" };

		Solution.solution(info, query);
	}

	static class Solution {
		public static int[] solution(String[] info, String[] query) {
			HashMap<String, Integer> language = new HashMap<String, Integer>();
			language.put("-", 0);
			language.put("cpp", 1);
			language.put("java", 1);
			language.put("python", 1);
			HashMap<String, Integer> position = new HashMap<String, Integer>();
			position.put("-", 0);
			position.put("backend", 1);
			position.put("frontend", 2);
			HashMap<String, Integer> career = new HashMap<String, Integer>();
			career.put("-", 0);
			career.put("junior", 1);
			career.put("senior", 2);
			HashMap<String, Integer> food = new HashMap<String, Integer>();
			food.put("-", 0);
			food.put("chicken", 1);
			food.put("pizza", 2);
			ArrayList<HashMap<String, Integer>> conditionList = new ArrayList<HashMap<String, Integer>>();
			conditionList.add(language);
			conditionList.add(position);
			conditionList.add(career);
			conditionList.add(food);

			int info_size = info.length;
			int query_size = query.length;

			int[] answer = new int[query_size];
			LinkedList<Integer> list = new LinkedList<Integer>();

			int[][] info_arr = new int[info_size][5];
			for (int i = 0; i < info_size; i++) {
				String[] arr = info[i].split(" ");
				for (int j = 0; j < 4; j++) {
					HashMap<String, Integer> a = conditionList.get(j);
					info_arr[i][j] = a.get(arr[j]);
				}
				info_arr[i][4] = Integer.parseInt(arr[4]);
			}

			int[][] query_arr = new int[query_size][5];
			for (int i = 0; i < query_size; i++) {
				String[] arr = query[i].split(" ");
				for (int j = 0; j < 4; j++) {
					query_arr[i][j] = conditionList.get(j).get(arr[j * 2]);
				}
				query_arr[i][4] = Integer.parseInt(arr[7]);
			}
			// 언어 직군 경력 소울푸드 점수
			// -, cpp, java, python
			// -, backend, frontend
			// -, junior, senior
			// -, chicken, pizza

			// 1. 각 쿼리를 가져옴
			for (int i = 0; i < query_size; i++) {
				// 2. 코테 점수 비교해서 해당하는 사람만 list에 넣음
				for (int j = 0; j < info_size; j++) {
					if (info_arr[j][4] >= query_arr[i][4]) {
						list.add(j);
					}
				}

				// 3. list에 있는 사람들한테 나머지 검색조건 설정
				Iterator<Integer> iter = list.iterator();
				while (iter.hasNext()) {
					int j = iter.next();
					for (int k = 0; k < 4; k++) {
						if (!(query_arr[i][k] == 0) && !(info_arr[j][k] == query_arr[i][k])) {
							iter.remove();
							break;
						}
					}
				}
				answer[i] = list.size();
				System.out.println(answer[i]);
				list.removeAll(list);
			}

			return answer;
		}
	}
}
