package main

import "fmt"
import "errors"

type Node struct {
	next *Node
	prev *Node
	data interface{}
	list *DoubleLinkedList
}

type DoubleLinkedList struct {
	first *Node
	last  *Node
}

func (s *DoubleLinkedList) Len() int {
	first, _ := s.First()
	length := 0
	if first == nil{
		return 0
	}

	for first.Next() != nil {
		length++
		first = first.Next()
	}
	return length
}

func (s *DoubleLinkedList) First() (*Node, error) {
	if s.first == nil{
		return nil, errors.New("Empty list")
	}else {
		return s.first, nil
	}

}

func (s *DoubleLinkedList) Last() (*Node, error){
	if s.last == nil{
		return nil, errors.New("Empty list")
	}else {
		return s.last, nil
	}
}

func (s *DoubleLinkedList) PushFront(v interface{}) {
	if val, _ := s.First(); val == nil {
		node := &Node{data: v, next: nil, prev: nil, list:s}
		s.first = node
		s.last = node
	} else {
		node := &Node{data: v, next: s.first, prev: nil, list:s}
		s.first.prev = node
		s.first = node
	}
}

func (s *DoubleLinkedList) PushBack(v interface{}) {
	if val, _ := s.Last(); val == nil {
		node := &Node{data: v, next: nil, prev: nil, list:s}
		s.first = node
		s.last = node
	} else {
		node := &Node{data: v, next: nil, prev: s.last, list:s}
		s.last.next = node
		s.last = node
	}
}

func (s *Node) Value() interface{} {
	return s.data
}

func (s *Node) Next() *Node {
	return s.next
}

func (s *Node) Prev() *Node {
	return s.prev
}

func (s *Node) Remove() {

	switch {
	case s.next != nil:
		s.next.prev = s.prev
		if s.prev == nil{
			s.list.first = s.next
		}else {
			s.prev.next = s.next
		}

	case s.prev != nil:
		s.prev.next = s.next
		if s.next == nil{
			s.list.last = s.prev
		}else {
			s.next.prev = s.prev
		}

	case s.next == nil && s.prev == nil:
		s.list.first = nil
		s.list.last = nil
	}
}

func (d *DoubleLinkedList) Print() {

	elem := d.first
	if elem == nil {
		fmt.Printf("List is empty")
	} else {
		index := 0
		for elem != nil {

			fmt.Printf("%v: %v\n", index, elem.data)
			index++
			elem = elem.next
		}
	}
}

func main() {
	dll := DoubleLinkedList{}

	s := []int{1,2,3,4,5}
	for _, v := range s {
		dll.PushBack(v)
	}
	fmt.Printf("Len %v\n", dll.Len())
	dll.Print()

	dll.first.next.Remove()
	fmt.Printf("Len %v\n", dll.Len())
	dll.Print()

	fmt.Printf("\n\n")

	dll.PushBack(100)
	dll.PushFront(200)
	dll.Print()
}
