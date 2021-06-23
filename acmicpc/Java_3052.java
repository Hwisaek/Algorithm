import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.HashSet;

public class Java_3052 {
	public static void main(String[] args) {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
		HashSet<Integer> set = new HashSet<Integer>();
		try {
			for (int i = 0; i < 10; i++) {
				set.add(Integer.parseInt(br.readLine()) % 42);
			}
			bw.write(String.valueOf(set.size()));
			bw.flush();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
