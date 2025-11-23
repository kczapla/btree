package btree

type Tree struct {
	root *Node
	t    int
}

type Node struct {
	keys     []int
	children []*Node
	leaf     bool
	n        int
}

func NewTree(t int) *Tree {
	return &Tree{
		t:    t,
		root: newNode(t, true),
	}
}

func (tree *Tree) maxKeysInTree() int {
	return tree.maxChildrenInTree() - 1
}

func (tree *Tree) maxChildrenInTree() int {
	return 2 * tree.t
}

func newNode(t int, isLeaf bool) *Node {
	return &Node{
		keys:     make([]int, 2*t-1),
		children: make([]*Node, 2*t),
		n:        0,
		leaf:     isLeaf,
	}
}

func (t *Tree) splitNode(parent *Node, childIndex int) {
	childX := parent.children[childIndex]
	newN := t.t - 1

	newChild := newNode(t.t, childX.leaf)

	middleKeyIndex := newN

	// move keys
	for i := range newN {
		newChild.keys[i] = childX.keys[middleKeyIndex+i+1]
		childX.keys[middleKeyIndex+i+1] = 0
		newChild.n += 1
	}

	// move pointers
	for i := range newN {
		newChild.children[i] = childX.children[middleKeyIndex+i]
		childX.children[middleKeyIndex+i] = nil
	}
	childX.n = newN

	keyInParentIndex := parent.n
	for i, key := range parent.keys {
		if childX.keys[middleKeyIndex] < key {
			keyInParentIndex = i
			break
		}
	}

	// move keys in parent
	for i := parent.n; keyInParentIndex+1 <= i; i-- {
		parent.keys[i] = parent.keys[i-1]
	}

	parent.keys[keyInParentIndex] = childX.keys[middleKeyIndex]
	childX.keys[middleKeyIndex] = 0

	// move pointers in the parent
	for i := parent.n + 1; keyInParentIndex+2 <= i; i-- {
		parent.children[i] = parent.children[i-1]
	}

	parent.children[keyInParentIndex+1] = newChild
	parent.n += 1
}

func (t *Tree) Insert(key int) {
	if t.root.n == t.maxKeysInTree() {
		newRoot := newNode(t.t, false)

		newRoot.children[0] = t.root
		t.root = newRoot

		t.splitNode(t.root, 0)
		t.insertToNonNonEmptyNode(t.root, key)
	} else {
		t.insertToNonNonEmptyNode(t.root, key)
	}
}

func (t *Tree) insertToNonNonEmptyNode(node *Node, key int) {
	if node.leaf {
		// find first greater key
		greaterKeyIndex := node.n
		for i := 0; i < node.n; i++ {
			if key < node.keys[i] {
				greaterKeyIndex = i
				break
			}
		}

		// shift keys by 1 to right
		for i := node.n - 1; greaterKeyIndex <= i; i-- {
			node.keys[i+1] = node.keys[i]
		}
		node.keys[greaterKeyIndex] = key
		node.n += 1
	} else {
		// find first greater key
		greaterKeyIndex := node.n
		for i := 0; i < node.n; i++ {
			if key < node.keys[i] {
				greaterKeyIndex = i
				break
			}
		}

		if node.children[greaterKeyIndex].n == t.maxKeysInTree() {
			t.splitNode(node, greaterKeyIndex)
			if node.keys[greaterKeyIndex] < key {
				t.insertToNonNonEmptyNode(node.children[greaterKeyIndex+1], key)
			} else {
				t.insertToNonNonEmptyNode(node.children[greaterKeyIndex], key)
			}
		} else {
			t.insertToNonNonEmptyNode(node.children[greaterKeyIndex], key)
		}
	}
}

func (t *Tree) Delete(key int) {
	t.deleteFromNode(t.root, key)
	if t.root.n != 0 {
		return
	}

	// there is an edge case when last key is delted from root as this is the only node that can have size 1
	// this addresses that case
	t.root = t.root.children[0]
}

func (t *Tree) deleteFromNode(node *Node, key int) {
	if node.leaf {
		ok, keyIndex := findKeyInNode2(node, key)

		if !ok {
			return
		}

		for i := keyIndex + 1; i < node.n; i++ {
			node.keys[i-1] = node.keys[i]
		}
		node.keys[node.n-1] = 0

		node.n -= 1
	} else {
		keyIndex := findKeyInNode(node, key)
		childY := node.children[keyIndex]
		if childY.n < t.t {
			if 0 < keyIndex && t.t <= node.children[keyIndex-1].n {
				childX := node.children[keyIndex-1]

				indexOfLastKeyInChildX := childX.n - 1
				valueOfLastKeyInChildX := childX.keys[indexOfLastKeyInChildX]

				t.deleteFromNode(childX, valueOfLastKeyInChildX)

				// shift to right by one in childY
				for i := childY.n; 0 < i; i-- {
					childY.keys[i] = childY.keys[i-1]
				}

				childY.keys[0] = node.keys[keyIndex-1]
				node.keys[keyIndex-1] = valueOfLastKeyInChildX
				childY.n += 1

				t.deleteFromNode(childY, key)
			} else if keyIndex < node.n && t.t <= node.children[keyIndex+1].n {
				childX := node.children[keyIndex+1]

				indexOfFirstKeyInChildX := 0
				valueOfFirstKeyInChildX := childX.keys[indexOfFirstKeyInChildX]

				t.deleteFromNode(childX, valueOfFirstKeyInChildX)

				childY.keys[childY.n] = node.keys[keyIndex]
				node.keys[keyIndex] = valueOfFirstKeyInChildX
				childY.n += 1

				t.deleteFromNode(childY, key)
			} else {
				if 0 < keyIndex {
					childX := node.children[keyIndex-1]
					childX.keys[childX.n] = node.keys[keyIndex-1]
					childX.n += 1

					for i := 0; i < childY.n; i++ {
						childX.keys[childX.n+i] = childY.keys[i]
					}
					childX.n += childY.n

					// move children. Use fact that t.t - 1 is index of middle elem and t.t is child right to that middle elem.
					// the ones that I want to populate
					for i := 0; i < childY.n+1; i++ {
						childX.children[t.t+i] = childY.children[i]
					}

					// shift keys in node to the left by 1
					for i := keyIndex + 1; i < node.n; i++ {
						node.keys[i-1] = node.keys[i]
					}
					node.keys[node.n-1] = 0

					// shift children in node to the left by 1
					for i := keyIndex + 1; i < node.n+1; i++ {
						node.children[i-1] = node.children[i]
					}
					node.children[node.n] = nil
					node.n -= 1

					t.deleteFromNode(childX, key)
				}
			}
		}
	}
}

func findKeyInNode(node *Node, key int) int {
	for i, k := range node.keys {
		if key < k {
			return i
		}
	}
	return node.n
}

func findKeyInNode2(node *Node, key int) (bool, int) {
	for i, k := range node.keys {
		if key == k {
			return true, i
		} else if key < k {
			return false, i
		}
	}
	return false, node.n
}
