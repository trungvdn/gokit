package main

import "fmt"

//isEmpty() function returns true if stack is empty or false in all other cases.
func (s *StackInt) IsEmpty() bool {
	return len(s.s) == 0
}

//length() function returns the number of elements in the stack.
func (s *StackInt) Length() int {
	return len(s.s)
}

//The print function will print the elements of the array.
func (s *StackInt) Print() {
	for i := 0; i < len(s.s); i++ {
		fmt.Println(s.s[i])
	}
}

//push() function append value to the data.
func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

/* In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the data array and return it. */

func (s *StackInt) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return res
}

/*top() function will first check that the stack is not empty
then returns the value stored in the top element
of stack (does not remove it). */
func (s *StackInt) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.s[len(s.s)-1]
}

func (s *StackLinkedList) Size() int {
	return s.size
}

/* IsEmpty() function will return true is size of the linked list is
equal to zero or false in all other cases. */
func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

/*First, the Peek() function will check if the stack is empty.
If not, it will return the peek value of stack i.e., will return the
head value of the linked list. */
func (s *StackLinkedList) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.head.value, true
}

//Push() function  will add new value at the start of the linked list.
func (s *StackLinkedList) Push(value int) {
	s.head = &Node{
		value: value,
		next:  s.head,
	}
	s.size++
}

/*In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the linked list and return it. */
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

/* Print() function will print the elements of the linked list. */
func (s *StackLinkedList) Print() {
	temp := s.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main() {
	stk := &StackInt{
		s: []int{1, 4, 5, 6, 7, 9},
	}
	sortedInsert(stk, 2)
}

func sortedInsert(stk *StackInt, element int) {
	// [1,3,4,6] and 5
	popEle := []int{}
	for {
		top := stk.Top()
		if top > element {
			ele := stk.Pop()
			popEle = append(popEle, ele)
		} else {
			stk.Push(element)
			break
		}
	}
	for i := len(popEle) - 1; i >= 0; i-- {
		stk.Push(popEle[i])
	}
	fmt.Println(stk)
}
