package DataStructure;

import java.util.Stack;

public class DataStack {

	public static void main(String[] args) {
		Stack<Integer> s = new Stack<>(); // Stack 생성

		s.push(5);	// 삽입(5)
		s.push(2);	// 삽입(2)
		s.push(2);	// 삽입(2)
		s.push(7);	// 삽입(7)
		s.pop();	// 삭제()
		s.push(1);	// 삽입(1)
		s.push(4);	// 삽입(4)
		s.pop();	// 삭제()

		// 결과 스택 구조 :
		// 1
		// 2
		// 2
		// 5
		while (!s.empty()) { // Stack.empty(): 스택이 비어있는지를 반환
			System.out.println(s.peek()); // Stack.peek(): 스택의 맨 위 값 반환
			s.pop();
		}
	}

}
