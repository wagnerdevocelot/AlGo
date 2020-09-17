package main

import "fmt"

//Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// main method
func main() {

}

//AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {

	var node = Node{}
	node.property = property

	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = &node
}

//IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {

	var node *Node

	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

//LastNode method returns the last Node
func (linkedList *LinkedList) LastNode() *Node {

	var node *Node
	var lastNode *Node

	for node = linkedList.headNode; node != nil; node = node.nextNode {

		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

//AddToEnd method adds the node with property to the end
func (linkedList *LinkedList) AddToEnd(property int) {

	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node
	lastNode = linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
	}
}

//NodeWithValue method returns Node given parameter property
func (linkedList *LinkedList) NodeWithValue(property int) *Node {

	var node *Node
	var nodeWith *Node

	for node = linkedList.headNode; node != nil; node = node.nextNode {

		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

//AddAfter method adds a node with nodeProperty after node with property
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {

	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)

	if nodeWith != nil {

		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

// RemoveNode from LinkedList
func (linkedList *LinkedList) RemoveNode(node int) bool {

	if linkedList.headNode == nil {
		return false
	}

	if linkedList.headNode.property == node {
		linkedList.headNode = linkedList.headNode.nextNode
		return true
	}

	current := linkedList.headNode
	for current.nextNode != nil {

		if current.nextNode.property == node {

			current.nextNode = current.nextNode.nextNode
			return true
		}
		current = current.nextNode
	}
	return false
}

// ListLen from LinkedList
func (linkedList *LinkedList) ListLen() {

	count := 0
	current := linkedList.headNode

	for current != nil {

		count++
		current = current.nextNode
	}
	fmt.Println(count)
}
