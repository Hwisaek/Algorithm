import java.util.Scanner;

public class Java_5622 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] dial = { 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10 };
		Scanner sc = new Scanner(System.in);

		String input = sc.next();
		int result = 0;
		for (int i = 0; i < input.length(); i++) {
			result += dial[input.charAt(i) - 'A'];
		}
		System.out.println(result);
	}

}
