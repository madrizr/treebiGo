package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func insertarNodo(root *Node, value int) *Node {
	if root == nil {
		return &Node{value, nil, nil}
	}

	if value < root.value {
		root.left = insertarNodo(root.left, value)
	} else {
		root.right = insertarNodo(root.right, value)
	}

	return root
}

func recorrerArbolIzqDerch(root *Node) {
	if root != nil {
		recorrerArbolIzqDerch(root.left)
		fmt.Printf("%d ", root.value)
		recorrerArbolIzqDerch(root.right)
	}
}

func main() {
	var root *Node

	root = insertarNodo(root, 00)
	insertarNodo(root, 10)
	insertarNodo(root, 4)
	insertarNodo(root, 20)
	insertarNodo(root, 01)
	insertarNodo(root, 98)
	insertarNodo(root, 2)

	recorrerArbolIzqDerch(root)
}