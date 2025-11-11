package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNode(t *testing.T) {
	currentTree := Node{
		keys: []int{4, 12, 0, 0, 0},
		n:    2,
		children: []*Node{
			{
				keys:     []int{1, 2, 0, 0, 0},
				children: []*Node{nil, nil, nil, nil, nil, nil},
				leaf:     true,
				n:        2,
			},
			{
				keys:     []int{5, 6, 7, 8, 9},
				children: []*Node{nil, nil, nil, nil, nil, nil},
				leaf:     true,
				n:        5,
			},
			{
				keys:     []int{13, 14, 0, 0, 0},
				children: []*Node{nil, nil, nil, nil, nil, nil},
				leaf:     true,
				n:        2,
			},
			nil, nil, nil,
		},
	}

	expectedTree := Node{
		keys: []int{4, 7, 12, 0, 0},
		n:    3,
		children: []*Node{
			{
				keys:     []int{1, 2, 0, 0, 0},
				leaf:     true,
				children: []*Node{nil, nil, nil, nil, nil, nil},
				n:        2,
			},
			{
				keys:     []int{5, 6, 0, 0, 0},
				leaf:     true,
				children: []*Node{nil, nil, nil, nil, nil, nil},
				n:        2,
			},
			{
				keys:     []int{8, 9, 0, 0, 0},
				leaf:     true,
				children: []*Node{nil, nil, nil, nil, nil, nil},
				n:        2,
			},
			{
				keys:     []int{13, 14, 0, 0, 0},
				leaf:     true,
				children: []*Node{nil, nil, nil, nil, nil, nil},
				n:        2,
			},
			nil,
			nil,
		},
	}

	splitNode(&currentTree, 1)

	assert.Equal(t, expectedTree, currentTree)
}
