package datastructures

type User struct {
	name string
	age  int
}

type Node struct {
	user User
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(user User) {
	node := &Node{user, nil}

	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *LinkedList) Display() {
	node := l.head

	for node != nil {
		println(node.user.name, node.user.age)
		node = node.next
	}
}

func LinkedListExemple()  {
	list := LinkedList{}

	list.Insert(User{"Eduardo", 25})
	list.Insert(User{"Maria", 20})
	list.Insert(User{"Jose", 30})

	list.Display()
}