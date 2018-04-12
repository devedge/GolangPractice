package main

import (
	"fmt"
)

type tree struct {
	root *node
}

type node struct {
	value int
	left  *node
	right *node
}

func (t *tree) insert(val int) {
	if t.root == nil {
		t.root = &node{value: val}
	} else {
		t.root.insert(val)
	}
}

func (t *tree) delete(val int) {
	if t.root == nil {
		fmt.Println("Cannot delete from empty tree")
	} else {
		// create a fake parent node to handle the node.delete()
		fakeParent := &node{left: t.root}

		t.root.delete(val, fakeParent)

		// if the values in 'left' are nil, set the tree root to nil
		if fakeParent.left == nil {
			t.root = nil
		}
	}
}

// insert a value, avoiding duplicates
func (n *node) insert(val int) {
	if n == nil {
		fmt.Println("Cannot insert into an empty tree")
		return
	}

	if val == n.value {
		fmt.Println("Value", val, "already present, not inserting...")

	} else if val < n.value {
		if n.left == nil {
			n.left = &node{value: val} // make a new 'left' value
		} else {
			n.left.insert(val)
		}

	} else if val > n.value {
		if n.right == nil {
			n.right = &node{value: val} // make a new 'right' value
		} else {
			n.right.insert(val)
		}
	}
}

// deletes can only work using a parent node
func (n *node) delete(val int, parent *node) {
	if n == nil {
		fmt.Println("Cannot delete from empty node")
		return
	}

	if val < n.value {
		n.left.delete(val, n)

	} else if val > n.value {
		n.right.delete(val, n)

	} else if val == n.value {
		if n.left == nil && n.right == nil {
			// if both are nil
			n.replaceNode(nil, parent)

		} else if n.left == nil {
			// if just the left is nil, replace the right one with its child
			n.replaceNode(n.right, parent)

		} else if n.right == nil {
			// if just the right is nil, replace the left one with its child
			n.replaceNode(n.left, parent)

		} else {
			// they're both non-nil

			// get min from smallest value on right subtree
			min, minparent := n.right.findMin(n)

			// 'delete' this node by replacing its value with that of 'min'
			n.value = min.value

			if min.right != nil {
				min.replaceNode(min.right, minparent) // replace 'min' with its child
			} else {
				min.replaceNode(nil, minparent) // delete 'min' with nil
			}
		}
	}
}

// replace 'n' with 'replacement' using 'parent'.
// 'replacement' can be nil
func (n *node) replaceNode(replacement, parent *node) {
	if n == nil {
		fmt.Println("Can't replace a nil node")
		return
	}

	if n == parent.left {
		parent.left = replacement

	} else if n == parent.right {
		parent.right = replacement
	}
}

// get the smallest value (& parent) of the passed-in node
func (n *node) findMin(parent *node) (*node, *node) {
	if n.left == nil {
		return n, parent
	} else {
		return n.left.findMin(n)
	}
}

// traverse & print a tree with Breadth-First Search
func (t *tree) traverseBFS() {
	if t.root == nil {
		fmt.Println("Cannot traverse a nil tree")
	} else {
		t.root.traverseBFS()
		fmt.Println()
	}
}
func (n *node) traverseBFS() {
	// print root
	fmt.Printf("%d ", n.value)

	// print left
	if n.left != nil {
		fmt.Printf("%d ", n.left.value)
	}

	// print right
	if n.right != nil {
		fmt.Printf("%d ", n.right.value)
	}

	// if left is not nil
	if n.left != nil {
		// if its left child is non-nil, traverse
		if n.left.left != nil {
			n.left.left.traverseBFS()
		}
		// if its right child is non-nil, traverse
		if n.left.right != nil {
			n.left.right.traverseBFS()
		}
	}

	// if right is non-nil
	if n.right != nil {
		// if its left child is non-nil, traverse
		if n.right.left != nil {
			n.right.left.traverseBFS()
		}
		// if its right child is non-nil, traverse
		if n.right.right != nil {
			n.right.right.traverseBFS()
		}
	}
}

// traverse & print a tree with Depth-First Search
func (t *tree) traverseDFS() {
	if t.root == nil {
		fmt.Println("Cannot traverse a nil tree")
	} else {
		t.root.traverseDFS()
		fmt.Println()
	}
}
func (n *node) traverseDFS() {
	fmt.Printf("%d ", n.value)

	if n.left != nil {
		n.left.traverseDFS()
	}

	if n.right != nil {
		n.right.traverseDFS()
	}
}

//func (t *tree) isValidBST() bool {

//}
//func (n *node) isValidBST(maxVal, minVal int) bool {

//}

func main() {
	var binaryTree tree

	binaryTree.insert(4)
	binaryTree.insert(10)
	binaryTree.insert(1)
	binaryTree.insert(0)
	binaryTree.insert(8)

	fmt.Printf("DFS: ")
	binaryTree.traverseDFS()

	fmt.Printf("BFS: ")
	binaryTree.traverseBFS()
}

/**

    4
   / \
  1   10
 /   /
0   8

DFS: 4, 1, 0, 10, 8

BFS: 4, 1, 10, 0, 8

**/
