package btree

type Tree struct {
	root *Node
}

type Node struct {
	keys     []int
	children []*Node
	leaf     bool
	n        int
}

func (t *Tree) splitNode(parent *Node, childIndex int) {
	childX := parent.children[childIndex]
	newN := childX.n / 2

	newChild := &Node{
		keys:     make([]int, childX.n),
		children: make([]*Node, childX.n+1),
		n:        newN,
		leaf:     childX.leaf,
	}

	middleKeyIndex := newN

	// move keys
	for i := range newN {
		newChild.keys[i] = childX.keys[middleKeyIndex+i+1]
		childX.keys[middleKeyIndex+i+1] = 0
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
	if t.root.n == len(t.root.keys) {
		size := len(t.root.children)
		newRoot := &Node{
			n:        0,
			leaf:     false,
			keys:     make([]int, size-1),
			children: make([]*Node, size),
		}
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
	} else {
		// find first greater key
		greaterKeyIndex := node.n
		for i := 0; i < node.n; i++ {
			if key < node.keys[i] {
				greaterKeyIndex = i
				break
			}
		}

		if node.children[greaterKeyIndex].n == len(node.keys) {
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
