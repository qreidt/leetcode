package main

import "fmt"

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := handleNodes(l1, l2, 0)
	return resultNode
}

func handleNodes(node1 *ListNode, node2 *ListNode, carry int) *ListNode {
	value := carry
	if node1 != nil {
		value += node1.Val
		node1 = node1.Next
	}

	if node2 != nil {
		value += node2.Val
		node2 = node2.Next
	}

	if value >= 10 {
		carry = 1
		value %= 10
	} else {
		carry = 0
	}

	if node1 == nil && node2 == nil && carry == 0 {
		return &ListNode{Val: value, Next: nil}
	}

	if node1 == nil && node2 == nil && carry == 1 {
		return &ListNode{Val: value, Next: &ListNode{Val: 1, Next: nil}}
	}

	return &ListNode{
		Val:  value,
		Next: handleNodes(node1, node2, carry),
	}
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to print a linked list
func printLinkedList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
		l = l.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	l1 := createLinkedList([]int{2, 4, 3})
	l2 := createLinkedList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	fmt.Print("Output: ")
	printLinkedList(result) // Expected: 7 -> 0 -> 8

	// Example 2
	l1 = createLinkedList([]int{0})
	l2 = createLinkedList([]int{0})
	result = addTwoNumbers(l1, l2)
	fmt.Print("Output: ")
	printLinkedList(result) // Expected: 0

	// Example 3
	l1 = createLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = createLinkedList([]int{9, 9, 9, 9})
	result = addTwoNumbers(l1, l2)
	fmt.Print("Output: ")
	printLinkedList(result) // Expected: 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
}
