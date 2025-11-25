package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNode(t *testing.T) {
	testCases := []struct {
		name                   string
		keysToInsertPreDelete  []int
		keysToDelete           []int
		keysToInsertPostDelete []int
		t                      int
		expectedTree           *Tree
	}{
		{
			name:                  "insert into full child at last place",
			keysToInsertPreDelete: []int{1, 2, 4, 13, 14, 5, 6, 12, 7, 8, 9, 10},
			keysToDelete:          []int{},
			t:                     3,
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
			name:                  "insert into full root",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6},
			keysToDelete:          []int{},
			t:                     3,
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
			name:                  "delete from full root",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5},
			keysToDelete:          []int{3},
			t:                     3,
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
			name:                  "delete from full root",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5},
			keysToDelete:          []int{3},
			t:                     3,
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
			name:                  "delete from leaf node with t - 1 keys and left sibling with eq or ge than t",
			keysToInsertPreDelete: []int{1, 2, 4, 5, 6, 3},
			keysToDelete:          []int{5},
			t:                     3,
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
			name:                  "delete from leaf node with t - 1 keys and right sibling with eq or ge than t",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6},
			keysToDelete:          []int{2},
			t:                     3,
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
			name:                  "delete from child with merging with left sibling and reducing tree height",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6},
			keysToDelete:          []int{2, 5},
			t:                     3,
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
			name:                  "delete from child with merging with right sibling and reducing tree height",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6},
			keysToDelete:          []int{2, 3},
			t:                     3,
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
		{
			name:                  "delete from second level child 1",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			keysToDelete:          []int{13},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{9, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							n:    2,
							keys: []int{3, 6, 0, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{1, 2, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{4, 5, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{7, 8, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						{
							n:    2,
							keys: []int{15, 18, 0, 0, 0},
							children: []*Node{
								{
									n:        4,
									leaf:     true,
									keys:     []int{10, 11, 12, 14, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{16, 17, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        3,
									leaf:     true,
									keys:     []int{19, 20, 21, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
		{
			// key 12 have to children - {10, 11}, {13, 14}
			// in this test case I want to delete 10 so that I force merging of left child into its right sibling testing that branch
			// this test is tests implementation detail - its fine because I tests particular implementation of btree
			name:                  "delete from second level child 2",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			keysToDelete:          []int{10},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{9, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							n:    2,
							keys: []int{3, 6, 0, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{1, 2, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{4, 5, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{7, 8, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						{
							n:    2,
							keys: []int{15, 18, 0, 0, 0},
							children: []*Node{
								{
									n:        4,
									leaf:     true,
									keys:     []int{11, 12, 13, 14, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{16, 17, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        3,
									leaf:     true,
									keys:     []int{19, 20, 21, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
		{
			// this varaint tests if child borrows key from sibling
			name:                  "delete from second level child 3",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			keysToDelete:          []int{17},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{9, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							n:    2,
							keys: []int{3, 6, 0, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{1, 2, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{4, 5, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{7, 8, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						{
							n:    3,
							keys: []int{12, 15, 19, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{10, 11, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{13, 14, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{16, 18, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{20, 21, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil,
							},
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
		{
			// this varaint tests if child borrows key from sibling
			name:                  "delete from second level child 4",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			keysToDelete:          []int{20},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{9, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							n:    2,
							keys: []int{3, 6, 0, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{1, 2, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{4, 5, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{7, 8, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						{
							n:    3,
							keys: []int{12, 15, 18, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{10, 11, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{13, 14, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{16, 17, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{19, 21, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil,
							},
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
		{
			name:                   "delete all keys and insert again",
			keysToInsertPreDelete:  []int{1, 2, 3, 4, 5},
			keysToDelete:           []int{1, 2, 3, 4, 5},
			keysToInsertPostDelete: []int{1},
			t:                      3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys:     []int{1, 0, 0, 0, 0},
					n:        1,
					leaf:     true,
					children: []*Node{nil, nil, nil, nil, nil, nil},
				},
			},
		},
		{
			// this varaint tests if child borrows key from sibling
			name:                  "delete from inter node",
			keysToInsertPreDelete: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			keysToDelete:          []int{15},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{9, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							n:    2,
							keys: []int{3, 6, 0, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{1, 2, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{4, 5, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        2,
									leaf:     true,
									keys:     []int{7, 8, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						{
							n:    2,
							keys: []int{12, 18, 0, 0, 0},
							children: []*Node{
								{
									n:        2,
									leaf:     true,
									keys:     []int{10, 11, 0, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        4,
									leaf:     true,
									keys:     []int{13, 14, 16, 17, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								{
									n:        3,
									leaf:     true,
									keys:     []int{19, 20, 21, 0, 0},
									children: []*Node{nil, nil, nil, nil, nil, nil},
								},
								nil, nil, nil,
							},
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
		{
			name:                  "delete from single key root and its right child with t or more elems",
			keysToInsertPreDelete: []int{1, 10, 20, 30, 40, 35},
			keysToDelete:          []int{20},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{30, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							n:        2,
							leaf:     true,
							keys:     []int{1, 10, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        2,
							leaf:     true,
							keys:     []int{35, 40, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
		{
			name:                  "delete from single key root and its left child with t or more elems",
			keysToInsertPreDelete: []int{1, 10, 20, 30, 40, 5},
			keysToDelete:          []int{20},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{10, 0, 0, 0, 0},
					n:    1,
					children: []*Node{
						{
							n:        2,
							leaf:     true,
							keys:     []int{1, 5, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        2,
							leaf:     true,
							keys:     []int{30, 40, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						nil, nil, nil, nil,
					},
				},
			},
		},
		{
			name:                  "delete from single key root and its both children with less than t keys",
			keysToInsertPreDelete: []int{1, 10, 20, 30, 40, 5},
			keysToDelete:          []int{5, 20},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					n:        4,
					leaf:     true,
					keys:     []int{1, 10, 30, 40, 0},
					children: []*Node{nil, nil, nil, nil, nil, nil},
				},
			},
		},
		{
			name:                  "delete mid key from inter node and its right child with less than t keys",
			keysToInsertPreDelete: []int{1, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 401},
			keysToDelete:          []int{500},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					n:    3,
					leaf: false,
					keys: []int{200, 401, 800, 0, 0},
					children: []*Node{
						{
							n:        2,
							leaf:     true,
							keys:     []int{1, 100, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        2,
							leaf:     true,
							keys:     []int{300, 400, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        2,
							leaf:     true,
							keys:     []int{600, 700, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        3,
							leaf:     true,
							keys:     []int{900, 1000, 1100, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						nil, nil,
					},
				},
			},
		},
		{
			name:                  "delete mid key from inter node and its left child with less than t keys",
			keysToInsertPreDelete: []int{1, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 501},
			keysToDelete:          []int{500},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					n:    3,
					leaf: false,
					keys: []int{200, 501, 800, 0, 0},
					children: []*Node{
						{
							n:        2,
							leaf:     true,
							keys:     []int{1, 100, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        2,
							leaf:     true,
							keys:     []int{300, 400, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        2,
							leaf:     true,
							keys:     []int{600, 700, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        3,
							leaf:     true,
							keys:     []int{900, 1000, 1100, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						nil, nil,
					},
				},
			},
		},
		{
			name:                  "delete mid key from inter node and its both children with less than t keys",
			keysToInsertPreDelete: []int{1, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100},
			keysToDelete:          []int{500},
			t:                     3,
			expectedTree: &Tree{
				t: 3,
				root: &Node{
					n:    2,
					leaf: false,
					keys: []int{200, 800, 0, 0, 0},
					children: []*Node{
						{
							n:        2,
							leaf:     true,
							keys:     []int{1, 100, 0, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        4,
							leaf:     true,
							keys:     []int{300, 400, 600, 700, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						{
							n:        3,
							leaf:     true,
							keys:     []int{900, 1000, 1100, 0, 0},
							children: []*Node{nil, nil, nil, nil, nil, nil},
						},
						nil, nil, nil,
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			currentTree := NewTree(testCase.t)
			for _, keyToInsert := range testCase.keysToInsertPreDelete {
				currentTree.Insert(keyToInsert)
			}
			for _, keyToDelete := range testCase.keysToDelete {
				currentTree.Delete(keyToDelete)
			}
			for _, keyToInsert := range testCase.keysToInsertPostDelete {
				currentTree.Insert(keyToInsert)
			}
			assert.Equal(t, testCase.expectedTree, currentTree)
		})
	}

}
