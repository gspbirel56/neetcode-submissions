type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Head     *LinkedListNode
	Capacity int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	if ll.Capacity == 0 || index > ll.Capacity - 1 {
		return -1
	}
	cursor := ll.Head
	for i := 0; i <= index - 1; i++ {
		cursor = cursor.Next
	}
	return cursor.Value
}

func (ll *LinkedList) InsertHead(val int) {
	newHead := &LinkedListNode{Value: val, Next: ll.Head}
	prevHead := ll.Head
	newHead.Next = prevHead
	ll.Head = newHead
	ll.Capacity++
}

func (ll *LinkedList) InsertTail(val int) {
	if ll.Capacity == 0 {
		ll.Head = &LinkedListNode{Value: val, Next: nil}
		ll.Capacity = 1
		return
	}
	cursor := ll.Head
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = &LinkedListNode{
		Value: val,
		Next:  nil,
	}
	ll.Capacity++

}

func (ll *LinkedList) Remove(index int) bool {
		if index > ll.Capacity-1 {
		return false
	}

	// handle 1-element case
	if index == 0 && ll.Capacity == 1 {
		ll.Head = nil
		ll.Capacity = 0
		return true
	}
	if index == 0 {
		ll.Head = ll.Head.Next
		ll.Capacity--
		return true
	}

	// handle populated list case
	cursor := ll.Head
	for i := 0; i <= index-2; i++ {
		//if cursor.Next != nil && cursor.Next.Next == nil {
		//	return false
		//}
		cursor = cursor.Next
	}
	cursor.Next = cursor.Next.Next
	ll.Capacity--
	return true

}

func (ll *LinkedList) GetValues() []int {
	values := make([]int, ll.Capacity)
	cursor := ll.Head
	for i := 0; i < ll.Capacity; i++ {
		values[i] = cursor.Value
		cursor = cursor.Next
	}
	return values
}
