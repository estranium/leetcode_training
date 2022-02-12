// Add_Two_Numbers

// You are given two non-empty linked lists
// representing two non-negative integers.
// The digits are stored in reverse order,
// and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero,
// except the number 0 itself.

package main

import (
	"fmt"
)

// TODO: write comments
// TODO: Add tests instead of a single check in main via Print

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var res *ListNode = nil
	var tail *ListNode = nil

	borrow := 0
	for (l1 != nil) || (l2 != nil) || (borrow != 0) {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}
		sum := (val1 + val2 + borrow)
		borrow = 0
		if (sum / 10) > 0 {
			borrow = 1
		}
		node := &ListNode{}
		node.Val = sum % 10

		if res == nil {
			res = node
		} else {
			tail.Next = node
		}
		tail = node

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	return res
}

func createListFromArray(array []int) *ListNode {
	var newList *ListNode = nil
	var tail *ListNode = nil

	for _, elem := range array {

		node := &ListNode{}
		node.Val = elem
		if newList == nil {
			newList = node
		} else {
			tail.Next = node
		}
		tail = node
	}

	return newList
}
func createArrayFromList(list *ListNode) []int {
	newArray := []int{}

	for list != nil {

		newArray = append(newArray, list.Val)
		list = list.Next
	}

	return newArray
}

func printList(list *ListNode) {
	fmt.Print("[")
	for list != nil {
		fmt.Print(list.Val)
		list = list.Next
		if list != nil {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	number1 := []int{2, 4, 3}
	number2 := []int{5, 6, 4}
	expectedResult := []int{7, 0, 8}
	list1 := createListFromArray(number1)
	list2 := createListFromArray(number2)
	resultList := addTwoNumbers(list1, list2)
	calculatedResult := createArrayFromList(resultList)

	fmt.Print("number1: ")
	fmt.Println(number1)
	fmt.Print("number2: ")
	fmt.Println(number2)

	fmt.Print("expectedResult: ")
	fmt.Println(expectedResult)
	fmt.Print("calculatedResult: ")
	fmt.Println(calculatedResult)

	fmt.Println()
	if intSliceEqual(calculatedResult, expectedResult) {
		fmt.Println("OK, expected result array is equal to calculated result")
	} else {
		fmt.Println("FAIL, expected result array is not equal to calculated result")
	}

}
