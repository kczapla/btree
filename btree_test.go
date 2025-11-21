package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNode(t *testing.T) {
	testCases := []struct {
		name         string
		keysToInsert []int
		t            int
		expectedTree *Tree
	}{
		{
			name:         "insert into full child at last place",
			keysToInsert: []int{1, 2, 4, 13, 14, 5, 6, 12, 7, 8, 9, 10},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
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
							n:        3,
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
		},
		{
			name:         "insert into full root",
			keysToInsert: []int{1, 2, 3, 4, 5, 6},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{3, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							keys:     []int{1, 2, 0, 0, 0},
							leaf:     true,
							children: []*Node{nil, nil, nil, nil, nil, nil},
							n:        2,
						},
						{
							keys:     []int{4, 5, 6, 0, 0},
							leaf:     true,
							children: []*Node{nil, nil, nil, nil, nil, nil},
							n:        3,
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			currentTree := NewTree(testCase.t)
			for _, keyToInsert := range testCase.keysToInsert {
				currentTree.Insert(keyToInsert)
			}
			assert.Equal(t, testCase.expectedTree, currentTree)
		})
	}

}
