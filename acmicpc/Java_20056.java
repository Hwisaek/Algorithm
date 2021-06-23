import java.util.ArrayList;
import java.util.Iterator;
import java.util.Scanner;

class Fireball {
	int r; // 행
	int c;// 렬
	int m;// 질량
	int d;// 방향
	int s;// 속도
}

public class Java_20056 {

	public static void main(String[] args) {

		Scanner sc = new Scanner(System.in);

		int[] dr = { -1, -1, 0, 1, 1, 1, 0, -1 };
		int[] dc = { 0, 1, 1, 1, 0, -1, -1, -1 };

		int n = sc.nextInt(); // 격자의 크기
		int m = sc.nextInt(); // 줄의 수, 파이어볼의 갯수
		int k = sc.nextInt(); // 이동 횟수

		// 각 칸에 파이어볼들이 들어있음.
		ArrayList<ArrayList<ArrayList<Fireball>>> fbarr = new ArrayList<ArrayList<ArrayList<Fireball>>>(n);
		for (int i = 0; i < n; i++) {
			fbarr.add(new ArrayList<ArrayList<Fireball>>(n));
			for (int j = 0; j < n; j++) {
				fbarr.get(i).add(new ArrayList<Fireball>());
			}
		}

		ArrayList<ArrayList<ArrayList<Fireball>>> fbarr_sub = new ArrayList<ArrayList<ArrayList<Fireball>>>(n);
		for (int i = 0; i < n; i++) {
			fbarr_sub.add(new ArrayList<ArrayList<Fireball>>(n));
			for (int j = 0; j < n; j++) {
				fbarr_sub.get(i).add(new ArrayList<Fireball>());
			}
		}

		// 파이어볼 입력 받아서 fbarr 의 해당 행, 렬에 넣음
		for (int j = 0; j < m; j++) {
			Fireball fb = new Fireball();
			fb.r = sc.nextInt() - 1;
			fb.c = sc.nextInt() - 1;
			fb.m = sc.nextInt();
			fb.s = sc.nextInt();
			fb.d = sc.nextInt();
			fbarr.get(fb.r).get(fb.c).add(fb);
		}

		// k번 이동
		while (k-- > 0) {

			// 1. 모든 파이어볼 이동
			for (int i = 0; i < fbarr.size(); i++) { // 행 반복
				for (int j = 0; j < fbarr.get(i).size(); j++) { // 열 반복
					Iterator<Fireball> iter = fbarr.get(i).get(j).iterator();
					while (iter.hasNext()) { // 각 칸에 들어있는 파이어볼 이동
						Fireball fb = iter.next();
						iter.remove();
						fb.r += dr[fb.d] * fb.s;
						fb.c += dc[fb.d] * fb.s;
						fb.r %= n;
						fb.c %= n;
						if (fb.r < 0) {
							fb.r += n;
						} else if (fb.r >= n) {
							fb.r -= n;
						}

						if (fb.c < 0) {
							fb.c += n;
						} else if (fb.c >= n) {
							fb.c -= n;
						}

						fbarr_sub.get(fb.r).get(fb.c).add(fb); // 보조배열에 추가
					}
				}
			}

			fbarr = fbarr_sub;
			fbarr_sub = new ArrayList<ArrayList<ArrayList<Fireball>>>(n);
			for (int i = 0; i < n; i++) { // 보조 배열 초기화
				fbarr_sub.add(new ArrayList<ArrayList<Fireball>>(n));
				for (int j = 0; j < n; j++) {
					fbarr_sub.get(i).add(new ArrayList<Fireball>());
				}
			}

			// 2.이동이 모두 끝난 뒤, 2개 이상의 파이어볼이 있는 칸에서는 다음과 같은 일이 일어난다.
			// 2.1같은 칸에 있는 파이어볼은 모두 하나로 합쳐진다.
			// 2.2파이어볼은 4개의 파이어볼로 나누어진다.
			// 2.3나누어진 파이어볼의 질량, 속력, 방향은 다음과 같다.
			// 2.3.1질량은 ?(합쳐진 파이어볼 질량의 합)/5?이다.
			// 2.3.2속력은 ?(합쳐진 파이어볼 속력의 합)/(합쳐진 파이어볼의 개수)?이다.
			// 2.3.3합쳐지는 파이어볼의 방향이 모두 홀수이거나 모두 짝수이면, 방향은 0, 2, 4, 6이 되고, 그렇지 않으면 1, 3, 5, 7이
			// 된다.
			// 2.4질량이 0인 파이어볼은 소멸되어 없어진다.

			for (int i = 0; i < fbarr.size(); i++) { // 행 반복
				for (int j = 0; j < fbarr.get(i).size(); j++) { // 열 반복
					Fireball bigfireball = new Fireball();
					bigfireball.r = i;
					bigfireball.c = j;
					boolean dflag = true; // 합쳐지는 파이어볼의 방향. true : 모두 홀수 또는 모두 짝수, false: 섞여있음.
					if (fbarr.get(i).get(j).size() >= 2) {
						int d_base = fbarr.get(i).get(j).get(0).d % 2;  // 첫번쨰 파이어볼의 방향이 짝수인지 홀수인지 판단

						for (Fireball fb : fbarr.get(i).get(j)) { // 각 칸에 들어있는 파이어볼 합치기
							bigfireball.m += fb.m;
							bigfireball.s += fb.s;
							if (fb.d % 2 != d_base) { // 값이 다르면
								dflag = false;
							}
						}
						int cnt = fbarr.get(i).get(j).size();
						fbarr.get(i).get(j).removeAll(fbarr.get(i).get(j)); // 해당 구역 파이어볼 초기화
						for (int a = 0; a < 4; a++) {
							int b;
							if (dflag) {
								b = a * 2;
							} else {
								b = a * 2 + 1;
							}
							Fireball smallfireball = new Fireball();
							smallfireball.r = i;
							smallfireball.c = j;
							smallfireball.m = bigfireball.m / 5;
							smallfireball.s = bigfireball.s / cnt;
							smallfireball.d = b;
							fbarr.get(i).get(j).add(smallfireball);
						}
					}
				}
			}
		}

		int result = 0;
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < n; j++) {
				for (Fireball fb : fbarr.get(i).get(j)) {
					result += fb.m;
				}
			}
		}

		System.out.println(result);

	}

}
