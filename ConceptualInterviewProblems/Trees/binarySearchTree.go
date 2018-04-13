// Implementation of a Binary Search Tree, and several functions including:
// - Insertion and Deletion
// - Printing Breadth-First Traversal & Depth-First Traversal

package main

import (
	"fmt"
)

func main() {
	// Always initialize and operate on a tree struct, not the nodes.
	// This is not only an example, but also a reference usage of this
	// 'binaryTree' implemtation.
	var binaryTree tree

  fmt.Println("Inserting these values sequentially: 4, 10, 1, 0, 8...")
	binaryTree.insert(4)
	binaryTree.insert(10)
	binaryTree.insert(1)
	binaryTree.insert(0)
	binaryTree.insert(8)

  fmt.Println("\nExpected tree structure reference for BFS and DFS answers:")
  fmt.Println(`
    4
   / \
  1   10
 /   /
0   8`)

	fmt.Println("\nExpected Depth-First Search  : 4 1 0 10 8")
	fmt.Printf("Actual Depth-First Search    : ")
	binaryTree.traverseDFS()

  fmt.Println("\nExpected Breadth-First Search: 4 1 10 0 8")
	fmt.Printf("Actual Breadth-First Search  : ")
	binaryTree.traverseBFS()

  fmt.Println("\nDeleting 10 from BST...")
  binaryTree.delete(10)
  fmt.Printf("BFS of new tree: ")
  binaryTree.traverseBFS()

  fmt.Println("\nInserting 3 to BST...")
  binaryTree.insert(3)
  fmt.Printf("BFS of new tree: ")
  binaryTree.traverseBFS()

  fmt.Println("\nInserting 9, 12, and 11 to BST...")
  binaryTree.insert(9)
  binaryTree.insert(12)
  binaryTree.insert(11)
  fmt.Printf("BFS of new tree: ")
  binaryTree.traverseBFS()

  fmt.Println("\nDeleting 8 from BST...")
  binaryTree.delete(8)
  fmt.Printf("BFS of new tree: ")
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

	// Value is smaller, so recursively delete on the left
	if val < n.value {
		n.left.delete(val, n)

		// Value is larger, so recursively delete on the right
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
			// Both children exist. From here, we replace this node's value with the
			// smallest one on the right subtree. Then, we delete that rightmost,
			// smallest child node.

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

// Traverse the Binary Tree using Breadth-First Search. This simply prints
// all elements in BFS order, on the same line.
// It works by printing the root, left child, right child, and then recurses
// into each child's children, repeating the process.
func (t *tree) traverseBFS() {
	if t.root == nil {
		fmt.Println("Cannot traverse a nil tree")
	} else {
		t.root.traverseBFS() // traverse the entire tree
		fmt.Println() // print a newline after the output
	}
}
// Traverse a node with BFS. This is the actual implementation
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

	// Since 'n.left' is already printed, print its left and right children
	if n.left != nil {
    // Traverse left's left child
		if n.left.left != nil {
			n.left.left.traverseBFS()
		}
		// Traverse left's right child
		if n.left.right != nil {
			n.left.right.traverseBFS()
		}
	}

	// Since 'n.right' is already printed, print its left and right children
	if n.right != nil {
		// Traverse right's left child
		if n.right.left != nil {
			n.right.left.traverseBFS()
		}
		// Traverse right's right child
		if n.right.right != nil {
			n.right.right.traverseBFS()
		}
	}
}

// Traverse the Binary Tree using Depth-First Search. Like the BFS
// implementation, this prints all the elements out, and then prints a newline.
//
func (t *tree) traverseDFS() {
	if t.root == nil {
		fmt.Println("Cannot traverse a nil tree")
	} else {
		t.root.traverseDFS()
		fmt.Println()
	}
}
// Implementation of DFS.
func (n *node) traverseDFS() {
  // With every traversal, print the node immediately
	fmt.Printf("%d ", n.value)

  // Traverse as deep left as possible. However, after recursing as deep as
  // possible, the next right node will be traversed as deeply as possible.
	if n.left != nil {
		n.left.traverseDFS()
	}

  // Traverse rightward immediately after
	if n.right != nil {
		n.right.traverseDFS()
	}
}

//func (t *tree) isValidBST() bool {

//}
//func (n *node) isValidBST(maxVal, minVal int) bool {

//}
