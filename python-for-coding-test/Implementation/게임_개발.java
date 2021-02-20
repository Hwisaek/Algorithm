package implementation;

import java.util.Scanner;

public class 게임_개발 {

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);

		System.out.printf("맵의 크기(N*M)을 입력하세요>> ");
		int n = sc.nextInt(), m = sc.nextInt();

		System.out.printf("캐릭터의 좌표와 방향(0: 북, 1: 동, 2: 남, 3: 서)을 입력하세요>> ");
		int x = sc.nextInt(), y = sc.nextInt(), d = sc.nextInt();

		System.out.println("맵의 정보를 입력하세요>> ");
		int[][] map = new int[n][m];
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < m; j++) {
				map[i][j] = sc.nextInt();
			}
		}

		int turn = 0; // 회전 횟수
		int result = 0; // 이동 횟수
		int stop = 0; // 움직임 버튼
		while (stop == 0) { // 캐릭터 이동 시작
			if (d > 0)
				d--;
			else
				d = 3;
			turn++;

			if (turn == 4) {
				switch (d) {
				case 0:
					if (map[x+1][y] == 0) {
						map[x][y]=1;
						x++;}
					else
						stop = 1;
					break;
				case 1:
					if (map[x][y-1] == 0) {
						map[x][y]=1;
						y--;}
					else
						stop = 1;
					break;
				case 2:
					if (map[x-1][y] == 0) {
						map[x][y]=1;
						x--;}
					else
						stop = 1;
					break;
				case 3:
					if (map[x][y+1] == 0) {
						map[x][y]=1;
						y++;}
					else
						stop = 1;
					break;
				}
			} else {
				switch (d) {
				case 0:
					if (map[x-1][y] == 0) {
						map[x][y]=1;
								x--;
						turn = 0;
						result++;
					}
					break;
				case 1:
					if (map[x][y+1] == 0) {
						map[x][y]=1;
								y++;
						turn = 0;
						result++;
					}
					break;
				case 2:
					if (map[x+1][y] == 0) {
						map[x][y]=1;
								x++;
						turn = 0;
						result++;
					}
					break;
				case 3:
					if (map[x][y-1] == 0) {
						map[x][y]=1;
								y--;
						turn = 0;
						result++;
					}
					break;
				}
			}
		}
		System.out.println("이동 횟수: " + result);
	}

}
