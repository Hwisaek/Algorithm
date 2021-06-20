import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Java_2577 {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		int total = Integer.parseInt(br.readLine()) * Integer.parseInt(br.readLine()) * Integer.parseInt(br.readLine());
		int[] count = new int[10];
		while (true) {
			count[total % 10]++;
			total /= 10;
			if (total == 0) {
				break;
			}
		}
		for (int i : count) {
			System.out.println(i);
		}
	}
}
