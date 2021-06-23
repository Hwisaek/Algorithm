import java.util.Scanner;

public class Java_8958 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);

		int n = sc.nextInt();

		String[] ox = new String[n];

		for (int i = 0; i < n; i++) {
			ox[i] = sc.next();
			int score = 0;
			int bonus = 0;

			for (int j = 0; j < ox[i].length(); j++) {
				if (ox[i].charAt(j) == 'O') {
					bonus++;
					score += bonus;
				} else {
					bonus = 0;
				}
			}
			System.out.println(score);
		}
	}

}
