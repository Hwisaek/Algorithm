import java.util.Scanner;

public class Java_2562 {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);

		int[] arr = new int[9];

		int max = 0;
		int index = -1;
		for (int i = 0; i < 9; i++) {
			arr[i] = scan.nextInt();
			if (arr[i] > max) {
				max = arr[i];
				index = i+1;
			}
		}

		System.out.println(max);
		System.out.println(index);
	}

}
