// Implementation of a Binary Search Tree, and several functions including:
// - Insertion and Deletion
// - Printing Breadth-First Traversal & Depth-First Traversal

package main

import (
	"fmt"
)

func main() {
	// Always initialize and operate on a tree struct, not the nodes.
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

// First, implement two structs: one for the root, and
// one for each node within the tree. This makes it easier
// to handle operations on the BST.
//
// Operations are ONLY conducted on the tree node. There is
// always a sort of 'wrapper function' that operates on the
// root node, and delegates actual behavior to a function
// that manipulates individual nodes.
type tree struct {
	root *node
}
type node struct {
	value int
	left  *node
	right *node
}

// Insertion operation. This is the one that will be called,
// as it handles calling individual nodes properly.
//
// If the root hasn't been initialized, create it as a node,
// using the value passed in. Otherwise, call the root node's
// 'insert' function with the value.
func (t *tree) insert(val int) {
	if t.root == nil {
		t.root = &node{value: val}
	} else {
		t.root.insert(val)
	}
}

// Insert a value into a node. This implements making a new 'right'
// or 'left' node recursively.
func (n *node) insert(val int) {
	// If the current node doesn't exist, this is an invalid,
	// empty tree.
	if n == nil {
		fmt.Println("Cannot insert into an empty tree")
		return
	}

	// If this node contains the passed-in value, don't insert.
	if val == n.value {
		fmt.Println("Value", val, "already present, not inserting...")

		// If it's smaller, recursively insert on the left
	} else if val < n.value {

		if n.left == nil { // If there is no left node, initialize it
			n.left = &node{value: val} // make a new 'left' value
		} else {
			n.left.insert(val) // recursively insert on the left node
		}

		// If it's larger, recursively insert on the right
	} else if val > n.value {

		if n.right == nil { // If there is no right node, initialize it
			n.right = &node{value: val} // make a new 'right' value
		} else {
			n.right.insert(val) // recursively insert on the right node
		}
	}
}

// Deletion operation. This is the one that will be called,
// as it handles calling individual nodes properly.
//
// If the root hasn't been initialized, print an error. This function
// can also delete the root node of a tree.
func (t *tree) delete(val int) {
	if t.root == nil {
		fmt.Println("Cannot delete from empty tree")
	} else {

		// The 'delete' function has the ability to replace values in
		// the 'left' or 'right' node. However, to do this, it needs a handle
		// to the 'parent' element. For the root node, we create a 'fake parent'
		// before calling 'delete'.
		fakeParent := &node{left: t.root}

		// Call 'delete' on the root node
		t.root.delete(val, fakeParent)

		// If the root's value is nil (fakeParent.left is nil), the
		// tree was deleted, so set the 'root' to nil
		if fakeParent.left == nil {
			t.root = nil
		}
	}
}

// Delete/replace the node that contains the specified value.
// If the node to delete has only one child, replace it with that child.
// If the node has two children, replace its value with the smallest number
//    on the right subtree/child.
//
// This function requires two helper functions:
// - replaceSelf()
// - findMin()
func (n *node) delete(val int, parent *node) {
	if n == nil {
		fmt.Println("Cannot delete from empty node")
		return
	}

	// Value is smaller, so recursively insert on the left
	if val < n.value {
		n.left.delete(val, n)

		// Value is larger, so recursively insert on the right
	} else if val > n.value {
		n.right.delete(val, n)

		// This is the node we want to delete
	} else if val == n.value {

		// If both children are nil, we can simply delete this node by replacing
		// it with 'nil'
		if n.left == nil && n.right == nil {
			n.replaceSelf(nil, parent)

			// If just the right child exists, replace this node with that right child
		} else if n.left == nil {
			n.replaceSelf(n.right, parent)

			// If just the left child exists, replace this node with that left child
		} else if n.right == nil {
			n.replaceSelf(n.left, parent)

		} else {
			// Both cildren exist. From here, we replace this node's value with the
			// smallest one on the right subtree. Then, we delete that child smallest
			// child node.

			// Get the node with the smallest value on the right subtree
			min, minparent := n.right.findMin(n)

			// 'delete' this node by replacing its value with 'min''s value
			n.value = min.value

			// Since 'min' is the smallest/leftmost value, first check if it has
			// any rightmost children. If so, we can 'delete' it by replacing it
			// with that rightmost child; otherwise, 'delete' it directly by
			// replacing it with 'nil'
			if min.right != nil {
				min.replaceSelf(min.right, minparent) // replace 'min' with its child
			} else {
				min.replaceSelf(nil, minparent) // delete 'min' with nil
			}
		}
	}
}

// Replace 'n' with the node specified by 'replacement'. Additionally,
// to be able to set values for the 'left' and 'right' nodes, we need
// 'n''s parent node
func (n *node) replaceSelf(replacement, parent *node) {
	if n == nil {
		fmt.Println("Can't replace a nil node")
		return
	}

	// Find out which node to replace, by checking its value.
	// If n is the left node, use 'parent' to replace its value
	if n == parent.left {
		parent.left = replacement

		// If n is the right node, use 'parent' to replace its value
	} else if n == parent.right {
		parent.right = replacement
	}
}

// Recursively find the smallest node from an input node. This function
// returns the node found, and its parent (n, parent)
func (n *node) findMin(parent *node) (*node, *node) {
	// Since this is a Binary Search Tree, the leftmost node is always
	// the smallest
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

/**

    4
   / \
  1   10
 /   /
0   8

DFS: 4, 1, 0, 10, 8

BFS: 4, 1, 10, 0, 8

**/
