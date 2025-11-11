package btree

type Node struct {
	keys     []int
	children []*Node
	leaf     bool
	n        int
}

func splitNode(parent *Node, childIndex int) {
	childX := parent.children[childIndex]
	newN := childX.n / 2

	newChild := &Node{
		keys:     make([]int, childX.n),
		children: make([]*Node, childX.n+1),
		n:        newN,
	}

	newChild.leaf = parent.children[childIndex].leaf

	middleKeyIndex := newN

	// move keys
	for i := 0; i < newN; i++ {
		newChild.keys[i] = childX.keys[middleKeyIndex+i+1]
		childX.keys[middleKeyIndex+i+1] = 0
	}

	// move pointers
	for i := 0; i < newN; i++ {
		newChild.children[i] = childX.children[middleKeyIndex+i]
		childX.children[middleKeyIndex+i] = nil
	}
	childX.n = newN

	keyInParentIndex := -1
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
