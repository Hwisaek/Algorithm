package level1;

public class 소수_만들기 {
	public static int solution(int[] nums) {
		boolean[] isNotPrime = checkNotPrime();
		int answer = 0;
		for (int i = 0; i < nums.length - 2; i++) {
			for (int j = i + 1; j < nums.length - 1; j++) {
				for (int k = j + 1; k < nums.length; k++) {
					int num = nums[i] + nums[j] + nums[k];
					if (!isNotPrime[num]) {
						answer++;
					}
				}
			}
		}
		return answer;
	}

	public static boolean[] checkNotPrime() {
		int n = 2997;
		boolean[] chk = new boolean[n + 1];
		chk[0] = true;
		chk[1] = true;
		for (int i = 2; i <= n; i++) {
			// 해당 값이 false. 즉, 소수이면 실행
			if (!chk[i]) {
				for (int j = 2; i * j <= n; j++) {
					chk[i * j] = true;
				}
			}
		}
		return chk;
	}

	public static void main(String[] args) {
		System.out.println(소수_만들기.solution(new int[] { 1, 2, 3, 4 }));
	}

}
