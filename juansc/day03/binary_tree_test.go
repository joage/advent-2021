package day03

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tests := []struct{
		name string
		inputs []string
		expectedChildren int
	}{
		{
			name: "empty tree",
			inputs: []string{},
			expectedChildren: 0,
		},
		{
			name: "one child",
			inputs: []string{"0"},
			expectedChildren: 1,
		},
		{
			name: "duplicated children",
			inputs: []string{"0", "0"},
			expectedChildren: 1,
		},
		{
			name: "two children",
			inputs: []string{"0", "1"},
			expectedChildren: 2,
		},
		{
			name: "three descendants",
			inputs: []string{"01", "00", "11"},
			expectedChildren: 3,
		},
		{
			name: "four descendants",
			inputs: []string{"01", "00", "11", "10"},
			expectedChildren: 4,
		},
		{
			name: "one long descendants",
			inputs: []string{"0000000"},
			expectedChildren: 1,
		},
		{
			name: "two long descendants",
			inputs: []string{"0000000", "01"},
			expectedChildren: 2,
		},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			tree := newbinaryTree(test.inputs)
			actual := tree.NumChildren()
			if actual != test.expectedChildren {
				t.Errorf("expected %d children, found %d", test.expectedChildren, actual)
				fmt.Printf("%#v\n", tree)
			}
		})
	}
}
