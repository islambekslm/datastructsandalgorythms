package linkedlist

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (ll *LinkedList) Add(value int) {
	node := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = node
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (ll *LinkedList) Remove(value int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
