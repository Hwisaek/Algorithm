package acmicpc;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Java_1152 {

	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

		String[] str_splited = br.readLine().trim().split(" ");
		if (str_splited[0].contentEquals("")) {
			System.out.println(0);
		} else {
			System.out.println(str_splited.length);
		}
	}

}
