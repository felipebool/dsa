// This package implements a binary search tree

package binary

import (
	"fmt"
	"github.com/felipebool/dsa/ds/element"
)

const (
	InOrder TraverseAlgorithm = iota
	PreOrder
	PostOrder
)

type TraverseAlgorithm int

// Node represents a node in a Tree. Each node has an
// element which implements the Element interface
// (GetKey() int) and two other Nodes. The Node's key
// must be greater than all the keys on the left subtree
// and smaller than all the keys on the right subtree.
type Node struct {
	element element.GetterSetter
	parent  *Node
	left    *Node
	right   *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(element element.GetterSetter) {
	if t.root == nil {
		t.root = &Node{element: element}
		return
	}

	var previous *Node
	current := t.root
	for current != nil {
		previous = current
		if element.GetKey() < current.element.GetKey() {
			current = current.left
			continue
		}
		current = current.right
	}

	newNode := &Node{element: element}
	if previous == nil {
		t.root = newNode
		return
	}

	newNode.parent = previous
	if newNode.element.GetKey() < previous.element.GetKey() {
		previous.left = newNode
		return
	}
	previous.right = newNode
}

func (t *Tree) Search(key int) *Node {
	current := t.root
	for current != nil {
		if key > current.element.GetKey() {
			current = current.right
			continue
		}
		if key < current.element.GetKey() {
			current = current.left
			continue
		}
		return current
	}
	return nil
}

func (t *Tree) Remove(key int) {
	node := t.Search(key)
	if node == nil {
		return
	}

	// we found the key in the root node
	if node.parent == nil {
		t.root = t.removeNode(node)
		return
	}

	if node.element.GetKey() < node.parent.element.GetKey() {
		node.parent.left = t.removeNode(node)
		return
	}
	node.parent.right = t.removeNode(node)
}

func (t *Tree) removeNode(node *Node) *Node {
	if node.right == nil {
		root := node.left
		node.left = nil
		return root
	}

	successor := t.leftMost(node.right)

	// if the successor is not the right node
	if successor.parent.element.GetKey() != node.element.GetKey() {
		successor.parent.left = successor.right
		successor.right = node.right
	}

	successor.left = node.left
	node.left = nil
	node.right = nil
	return successor
}

func (t *Tree) Traverse(algorithm TraverseAlgorithm) string {
	switch algorithm {
	case InOrder:
		return t.inOrder(t.root)
	case PreOrder:
		return t.preOrder(t.root)
	case PostOrder:
		return t.postOrder(t.root)
	default:
		return "unknown traversal algorithm"
	}
}

func (t *Tree) inOrder(root *Node) string {
	if root == nil {
		return ""
	}
	result := t.inOrder(root.left)
	result += fmt.Sprintf("[%d] ", root.element.GetKey())
	result += t.inOrder(root.right)
	return result
}

func (t *Tree) preOrder(root *Node) string {
	if root == nil {
		return ""
	}
	result := fmt.Sprintf("[%d] ", root.element.GetKey())
	result += t.preOrder(root.left)
	result += t.preOrder(root.right)
	return result
}

func (t *Tree) postOrder(root *Node) string {
	if root == nil {
		return ""
	}
	result := t.postOrder(root.left)
	result += t.postOrder(root.right)
	result += fmt.Sprintf("[%d] ", root.element.GetKey())
	return result
}

func (t *Tree) leftMost(node *Node) *Node {
	if node == nil {
		return nil
	}
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func NewTree() *Tree {
	return &Tree{}
}
