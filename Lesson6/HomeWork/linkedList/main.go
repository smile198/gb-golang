package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n *Node) reverse(prevNode *Node) *Node {
	result := prevNode

	if n != nil {
		result = n.next.reverse(n)
		n.next = prevNode
	}

	return result
}

func main() {
	fmt.Println("Введите последовательность целых чисел по одному.")
	fmt.Println("По окончанию последовательности введите любой символ, кроме цифр.")
	fmt.Println("")
	fmt.Print("Введите первое целое число: ")

	var prevNode *Node
	var firstNode *Node

	var number int
	for {
		if prevNode != nil {
			fmt.Print("Введите следующее число: ")
		}

		_, err := fmt.Scanln(&number)
		if err != nil {
			break
		}

		currentNode := Node{data: number}
		if prevNode != nil {
			(*prevNode).next = &currentNode
		} else {
			firstNode = &currentNode
		}
		prevNode = &currentNode
	}

	fmt.Println("")
	fmt.Println("Заданная последовательность:")
	printNode(*firstNode)
	fmt.Println("")

	newFirstNode := firstNode.reverse(nil)
	fmt.Println("Последовательность в обратном порядке:")
	printNode(*newFirstNode)
	fmt.Println("")
}

func printNode(node Node) {
	fmt.Print(node.data)
	if node.next != nil {
		fmt.Print(" -> ")
		printNode(*node.next)
	}
}
