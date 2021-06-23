package level2;

import java.util.Stack;

public class 짝지어_제거하기 {

	public static int solution(String s) {
		Stack<Character> stack = new Stack<Character>();
		for (int i = 0; i < s.length(); i++) {
			if (stack.isEmpty()) {
				stack.push(s.charAt(i));
			} else {
				char last = stack.peek();
				char now = s.charAt(i);
				if (last == now) {
					stack.pop();
				} else {
					stack.push(now);
				}
			}
		}
		return stack.size() == 0 ? 1 : 0;
	}

	public static void main(String[] args) {
		System.out.println(solution(""));
		System.out.println(solution("cdcd"));
	}

//	public static int solution(String s) {
//		char last = 0;
//		LinkedList<Character> list = new LinkedList<Character>();
//		for (char c : s.toCharArray()) {
//			list.add(c);
//		}
//		for (int i = 0; i < list.size(); i++) {
//			if (last == list.get(i)) {
//				list.remove(i - 1);
//				list.remove(i - 1);
//				last = 0;
//				i = -1;
//			} else {
//				last = list.get(i);
//			}
//		}
//		return list.size() == 0 ? 1 : 0;
//	}

//	public int solution(String s)
//    {
//        StringBuilder sb = new StringBuilder(s);
//		char last = 0;
//		for (int i = 0; i < sb.length(); i++) {
//			if (last == sb.charAt(i)) {
//				sb = sb.delete(i - 1, i + 1);
//				last = 0;
//				i = -1;
//			} else {
//				last = sb.charAt(i);
//			}
//		}
//
//		return sb.length() == 0 ? 1 : 0;
//    }
}
