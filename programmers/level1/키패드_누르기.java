package level1;

public class 키패드_누르기 {
	public static String solution(int[] numbers, String hand) {
		StringBuilder answer = new StringBuilder();
		hand = hand.substring(0, 1).toUpperCase();

		int[] left = { 3, 0 };
		int[] right = { 3, 2 };

		int length = numbers.length;

		for (int i = 0; i < length; i++) {
			int row = (numbers[i] - 1) / 3;
			int col = (numbers[i] - 1) % 3;
			row = numbers[i] == 0 ? 3 : row;
			col = numbers[i] == 0 ? 1 : col;

			double leftDist = Math.abs(left[0] - row) + Math.abs(left[1] - col);
			double rightDist = Math.abs(right[0] - row) + Math.abs(right[1] - col);

			if (col == 0) {
				answer.append("L");
				left[0] = row;
				left[1] = col;
			} else if (col == 2) {
				answer.append("R");
				right[0] = row;
				right[1] = col;
			} else {
				if (leftDist < rightDist) {
					answer.append("L");
					left[0] = row;
					left[1] = col;
				} else if (rightDist < leftDist) {
					answer.append("R");
					right[0] = row;
					right[1] = col;
				} else {
					if (hand.equals("R")) {
						answer.append("R");
						right[0] = row;
						right[1] = col;
					} else {
						answer.append("L");
						left[0] = row;
						left[1] = col;
					}
				}
			}
		}

		return answer.toString();
	}

	public static void main(String[] args) {
//		int[] arr = { 1, 3, 4, 5, 8, 2, 1, 4, 5, 9, 5 };
//		System.out.println(키패드_누르기.solution(arr, "right"));
		int[] arr2 = { 7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2 };
		System.out.println(키패드_누르기.solution(arr2, "left"));
	}
}
