package main

import "fmt"

// Node adalah elemen dari linked list
type Node[T any] struct {
	Value T        // Nilai dari node
	Next  *Node[T] // Pointer ke node berikutnya
}

// LinkedList adalah struktur data generik
type LinkedList[T any] struct {
	Head *Node[T] // Pointer ke node pertama
}

// Add menambahkan elemen baru ke dalam linked list
func (l *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// Display mencetak semua elemen dalam linked list
func (l *LinkedList[T]) Display() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func main() {
	// LinkedList untuk int
	intList := LinkedList[int]{}
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	fmt.Println("LinkedList of int:")
	intList.Display()

	// LinkedList untuk string
	stringList := LinkedList[string]{}
	stringList.Add("Hello")
	stringList.Add("World")
	fmt.Println("LinkedList of string:")
	stringList.Display()

	// LinkedList untuk float
	floatList := LinkedList[float64]{}
	floatList.Add(1.1)
	floatList.Add(1.3)
	floatList.Add(3.5)
	fmt.Println("LinkedList of float:")
	floatList.Display()

}
