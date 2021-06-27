import java.io.*;
import java.util.StringTokenizer;

public class Java_10828 {

	public static void main(String[] args) throws NumberFormatException, IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
		StringBuilder stb = new StringBuilder();
		StringTokenizer st;

		int n = Integer.parseInt(br.readLine());
		int[] stack = new int[n];
		int top = -1;

		for (int i = 0; i < n; i++) {
			st = new StringTokenizer(br.readLine());
			switch (st.nextToken()) {
			case "push":
				stack[++top] = Integer.parseInt(st.nextToken());
				break;
			case "pop":
				stb.append((top == -1 ? -1 : stack[top--]) + "\n");
				break;
			case "size":
				stb.append((top + 1) + "\n");
				break;
			case "empty":
				stb.append((top == -1 ? 1 : 0) + "\n");
				break;
			case "top":
				stb.append((top == -1 ? -1 : stack[top]) + "\n");
				break;
			}
		}
		bw.write(stb.toString());
		bw.flush();
		bw.close();
	}

}
