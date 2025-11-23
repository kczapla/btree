package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNode(t *testing.T) {
	testCases := []struct {
		name         string
		keysToInsert []int
		keysToDelete []int
		t            int
		expectedTree *Tree
	}{
		{
			name:         "insert into full child at last place",
			keysToInsert: []int{1, 2, 4, 13, 14, 5, 6, 12, 7, 8, 9, 10},
			keysToDelete: []int{},
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
			keysToDelete: []int{},
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
		{
			name:         "delete from full root",
			keysToInsert: []int{1, 2, 3, 4, 5},
			keysToDelete: []int{3},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys:     []int{1, 2, 4, 5, 0},
					n:        4,
					children: []*Node{nil, nil, nil, nil, nil, nil},
					leaf:     true,
				},
			},
		},
		{
			name:         "delete from full root",
			keysToInsert: []int{1, 2, 3, 4, 5},
			keysToDelete: []int{3},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys:     []int{1, 2, 4, 5, 0},
					n:        4,
					children: []*Node{nil, nil, nil, nil, nil, nil},
					leaf:     true,
				},
			},
		},
		{
			name:         "delete from leaf node with t - 1 keys and left sibling with eq or ge than t",
			keysToInsert: []int{1, 2, 4, 5, 6, 3},
			keysToDelete: []int{5},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{3, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							leaf:     true,
							n:        2,
							keys:     []int{1, 2, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							leaf:     true,
							n:        2,
							keys:     []int{4, 6, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						nil, nil, nil, nil},
					leaf: false,
				},
			},
		},
		{
			name:         "delete from leaf node with t - 1 keys and right sibling with eq or ge than t",
			keysToInsert: []int{1, 2, 3, 4, 5, 6},
			keysToDelete: []int{2},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{4, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							leaf:     true,
							n:        2,
							keys:     []int{1, 3, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							leaf:     true,
							n:        2,
							keys:     []int{5, 6, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						nil, nil, nil, nil},
					leaf: false,
				},
			},
		},
		{
			name:         "delete from child with merging with left sibling and reducing tree height",
			keysToInsert: []int{1, 2, 3, 4, 5, 6},
			keysToDelete: []int{2, 5},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys:     []int{1, 3, 4, 6, 0},
					n:        4,
					leaf:     true,
					children: []*Node{nil, nil, nil, nil, nil, nil},
				},
			},
		},
		{
			name:         "delete from child with merging with right sibling and reducing tree height",
			keysToInsert: []int{1, 2, 3, 4, 5, 6},
			keysToDelete: []int{2, 3},
			t:            3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys:     []int{1, 4, 5, 6, 0},
					n:        4,
					leaf:     true,
					children: []*Node{nil, nil, nil, nil, nil, nil},
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
			for _, keyToDelete := range testCase.keysToDelete {
				currentTree.Delete(keyToDelete)
			}
			assert.Equal(t, testCase.expectedTree, currentTree)
		})
	}

}
