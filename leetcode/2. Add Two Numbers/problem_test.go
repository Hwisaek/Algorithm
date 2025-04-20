package problem

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

// createList는 배열에서 연결 리스트를 생성합니다
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// listToSlice는 연결 리스트를 슬라이스로 변환합니다
func listToSlice(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "두 숫자 더하기: 342 + 465 = 807",
			l1:       []int{2, 4, 3}, // 342 (역순으로 표현)
			l2:       []int{5, 6, 4}, // 465 (역순으로 표현)
			expected: []int{7, 0, 8}, // 807 (역순으로 표현)
		},
		{
			name:     "받아올림이 있는 경우: 99 + 99 = 198",
			l1:       []int{9, 9},
			l2:       []int{9, 9},
			expected: []int{8, 9, 1},
		},
		{
			name:     "서로 다른 길이의 리스트: 999 + 88 = 1087",
			l1:       []int{9, 9, 9},
			l2:       []int{8, 8},
			expected: []int{7, 8, 0, 1},
		},
		{
			name:     "빈 리스트와 숫자 더하기",
			l1:       []int{},
			l2:       []int{5, 6, 7},
			expected: []int{5, 6, 7},
		},
		{
			name:     "두 개의 빈 리스트",
			l1:       []int{},
			l2:       []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.l1)
			l2 := createList(tt.l2)

			result := addTwoNumbers(l1, l2)
			resultSlice := listToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("addTwoNumbers() = %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}
