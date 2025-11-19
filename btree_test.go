package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNode(t *testing.T) {
	testCases := []struct {
		name         string
		currentTree  *Node
		expectedTree *Node
	}{
		{
			name: "insert into full child",
			currentTree: &Node{
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
			},
			expectedTree: &Node{
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
						keys:     []int{8, 9, 10, 0, 0},
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
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			Insert(testCase.currentTree, 10)
			assert.Equal(t, testCase.expectedTree, testCase.currentTree)
		})
	}

}
