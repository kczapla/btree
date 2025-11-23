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
	keyIndex := -1
	for i, k := range t.root.keys {
		if key == k {
			keyIndex = i
			break
		}
	}

	if keyIndex == -1 {
		return
	}

	for i := keyIndex + 1; i < t.root.n; i++ {
		t.root.keys[i-1] = t.root.keys[i]
	}
	t.root.keys[t.root.n-1] = 0

	t.root.n -= 1
}
