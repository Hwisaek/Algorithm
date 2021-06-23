import java.util.Arrays;
import java.util.Scanner;

public class Java_10809 {
	public static void main(String[] args) {

		Scanner sc = new Scanner(System.in);

		String s = sc.next();

		int[] chk = new int[26];
		Arrays.fill(chk, -1);

		for (int i = 0; i < s.length(); i++) {
			int now = s.charAt(i) - 'a';
			if (chk[now] == -1) {
				chk[now] = i;
			}
		}

		for (int i = 0; i < 26; i++) {
			System.out.printf("%d ", chk[i]);
		}
	}
}
