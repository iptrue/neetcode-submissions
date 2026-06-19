type ListNode struct {
	value int
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}

func (ll *LinkedList) Get(index int) int {
	cur := ll.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}
	if cur == nil {
		return -1
	}
	return cur.value
}

func (ll *LinkedList) InsertHead(val int) {
	node := &ListNode{value: val, next: ll.head}
	ll.head = node
	if ll.tail == nil {
		ll.tail = node
	}
}

func (ll *LinkedList) InsertTail(val int) {
	node := &ListNode{value: val, next: nil}
	if ll.tail == nil {
		ll.head = node
		ll.tail = node
		return
	}
	ll.tail.next = node
	ll.tail = node
}

func (ll *LinkedList) Remove(index int) bool {
	if ll.head == nil {
		return false
	}
	if index == 0 {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}
		return true
	}
	prev := ll.head
	for i := 0; i < index-1; i++ {
		if prev.next == nil {
			return false
		}
		prev = prev.next
	}
	if prev.next == nil {
		return false
	}
	removed := prev.next
	prev.next = removed.next
	if removed == ll.tail {
		ll.tail = prev
	}
	return true
}

func (ll *LinkedList) GetValues() []int {
	values := []int{}
	for cur := ll.head; cur != nil; cur = cur.next {
		values = append(values, cur.value)
	}
	return values
}
