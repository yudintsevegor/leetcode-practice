package main

import (
	"log"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(in []int) ListNode {
	nodes := make([]ListNode, 0, len(in))
	for i, e := range in {
		if i == 0 {
			nodes = append(nodes, ListNode{
				Val: e,
			})
			continue
		}

		nodes = append(nodes, makeNode(e))
	}

	theLast := len(nodes) - 1
	for i := theLast; i >= 0; i-- {
		if i == 0 {
			continue
		}

		pushNode(&nodes[i], &nodes[i-1])
	}

	return nodes[theLast]
}

func main() {
	sliceL1 := []int{5, 3, 4}
	sliceL2 := []int{5, 5, 3}

	l1 := createList(sliceL1)
	l2 := createList(sliceL2)

	s1 := getSliceFromList(&l1)
	s2 := getSliceFromList(&l2)

	log.Println(s1)
	log.Println(s2)

	s1, s2 = appendZeros(s1, s2)

	sum := addition(s1, s2)
	log.Println(sum)

	res := addTwoNumbers(&l1, &l2)
	for {
		if res == nil {
			break
		}
		log.Printf("val: %+v", res.Val)
		res = res.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := getSliceFromList(l1)
	s2 := getSliceFromList(l2)

	s1, s2 = appendZeros(s1, s2)

	sum := addition(s1, s2)

	// log.Println("SUM: ", sum)
	nodes := make([]ListNode, 0, len(sum))
	for i, e := range sum {
		if i == len(sum)-1 {
			nodes = append(nodes, ListNode{
				Val: e,
			})
			continue
		}
		nodes = append(nodes, makeNode(e))
	}

	// log.Println(nodes)

	out := ListNode{}
	for i := range nodes {
		if i == len(nodes)-1 {
			continue
		}

		pushNode(&nodes[i], &nodes[i+1])
	}

	out = nodes[0]

	return &out
}

func getSliceFromList(list *ListNode) []int {
	out := make([]int, 0, 1)
	for {
		out = append(out, list.Val)

		if list.Next == nil {
			break
		}

		list = list.Next
	}

	return out
}

func appendZeros(s1, s2 []int) ([]int, []int) {
	if len(s1) == len(s2) {
		return s1, s2
	}

	diff := math.Abs(float64(len(s1) - len(s2)))
	if len(s1) > len(s2) {
		for i := 0; i < int(diff); i++ {
			s2 = append(s2, 0)
		}
	}

	if len(s1) < len(s2) {
		for i := 0; i < int(diff); i++ {
			s1 = append(s1, 0)
		}
	}

	return s1, s2
}

func addition(in1, in2 []int) []int {
	in1Len := len(in1)
	out := make([]int, in1Len+1)
	digit := 0

	for i := 0; i < in1Len; i++ {
		sum := in1[i] + in2[i] + digit
		if sum >= 10 {
			digit = 1
			sum %= 10
		} else {
			digit = 0
		}

		out[i] = sum
	}

	if digit == 1 {
		out[in1Len] = digit
	}

	if out[in1Len] == 0 {
		return out[:in1Len]
	}

	return out
}

func pushNode(l1, l2 *ListNode) {
	l1.Next = l2
}

func makeNode(val int) ListNode {
	return ListNode{
		Val:  val,
		Next: new(ListNode),
	}
}
