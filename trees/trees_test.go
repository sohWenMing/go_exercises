package trees

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var leafTree Tree = Tree{
	"leaf tree",
	[]*Tree{},
}

var oneBranchTree Tree = Tree{
	"one branch tree",
	[]*Tree{&leafTree},
}

var twoBranchTree Tree = Tree{
	"two branch tree",
	[]*Tree{&leafTree, &leafTree},
}

func TestDepth(t *testing.T) {
	type test struct {
		name     string
		tree     Tree
		expected int
	}
	tests := []test{
		{
			"basic 1 level depth",
			leafTree,
			1,
		},
		{
			"basic test found",
			Tree{
				"three level tree",
				[]*Tree{
					&twoBranchTree,
				},
			},
			3,
		},
	}

	for _, test := range tests {
		got := test.tree.Depth()
		if got != test.expected {
			t.Errorf("\ngot :%v\nwant: %v", got, test.expected)
		}
	}

}

func TestPrintTree(t *testing.T) {

	type test struct {
		name     string
		tree     Tree
		expected string
	}

	tests := []test{
		{
			name: "test basic tree",

			tree: Tree{
				"test tree no branches",
				[]*Tree{},
			},
			expected: "test tree no branches",
		},
		{
			name: "test one branch",
			tree: oneBranchTree,
			expected: (oneBranchTree.getTreeLabel() +
				"\n" +
				"\t" +
				leafTree.getTreeLabel()),
		},
		{
			name: "final test",
			tree: Tree{
				"final test tree",
				[]*Tree{
					{
						"first branch",
						[]*Tree{&leafTree},
					},
					{
						"second branch",
						[]*Tree{&twoBranchTree},
					},
				},
			},
			expected: ("final test tree" +
				"\n" +
				"\t" + "first branch" +
				"\n" +
				"\t\t" + leafTree.getTreeLabel() +
				"\n" +
				"\t" + "second branch" +
				"\n" +
				"\t\t" + twoBranchTree.getTreeLabel() +
				"\n" +
				"\t\t\t" + leafTree.getTreeLabel() +
				"\n" +
				"\t\t\t" + leafTree.getTreeLabel()),
		},
	}

	for _, test := range tests {
		got := test.tree.PrintTree("")
		want := test.expected
		if got != want {
			t.Errorf("got:\n %s\nwant:\n %s", got, want)
		}

	}
}

func TestCountLeaves(t *testing.T) {

	type test struct {
		name     string
		tree     Tree
		expected int
	}

	tests := []test{
		{
			name: "test basic tree",

			tree: Tree{
				"test tree no branches",
				[]*Tree{},
			},
			expected: 1,
		},
		{
			name:     "test one branch",
			tree:     oneBranchTree,
			expected: 1,
		},
		{
			name: "final test",
			tree: Tree{
				"final test tree",
				[]*Tree{
					{
						"first branch",
						[]*Tree{&leafTree},
					},
					{
						"second branch",
						[]*Tree{&twoBranchTree},
					},
				},
			},
			expected: 3,
		},
	}

	for _, test := range tests {
		got := test.tree.CountLeaves()
		want := test.expected
		if got != want {
			t.Errorf("got:\n %v\nwant:\n %v", got, want)
		}

	}
}

func TestContains(t *testing.T) {

	type test struct {
		name         string
		searchString string
		tree         Tree
		expected     bool
	}

	tests := []test{
		{
			"basic test found",
			"basic tree",
			Tree{
				"basic tree",
				[]*Tree{},
			},
			true,
		},
		{
			"basic test not found",
			"this should not be found",
			Tree{
				"basic tree",
				[]*Tree{},
			},
			false,
		},
		{
			"shoudl be found in inner branch",
			"leaf tree",
			Tree{
				"basic tree",
				[]*Tree{&twoBranchTree},
			},
			true,
		},
	}

	for _, test := range tests {

		var builder strings.Builder
		got := test.tree.Contains(test.searchString)

		if got != test.expected {
			builder.WriteString(fmt.Sprintf("test failed. test name: %s\n", test.name))
			builder.WriteString(fmt.Sprintf("got: %v\n", got))
			builder.WriteString(fmt.Sprintf("want: %v\n", test.expected))
			t.Error(builder.String())
		}
	}
}

func doubleString(input string) string {
	return input + "|" + input
}
func TestMap(t *testing.T) {
	type test struct {
		name     string
		tree     Tree
		expected Tree
		testFunc func(string) string
	}

	tests := []test{
		{
			name: "basic tree",
			tree: Tree{
				"one line",
				[]*Tree{},
			},
			expected: Tree{
				"one line|one line",
				[]*Tree{},
			},
			testFunc: doubleString,
		},
		{
			name: "complex tree",
			tree: Tree{
				"one line",
				[]*Tree{&twoBranchTree},
			},
			expected: Tree{
				"one line|one line",
				[]*Tree{
					&Tree{
						"two branch tree|two branch tree",
						[]*Tree{
							&Tree{"leaf tree|leaf tree", []*Tree{}},
							&Tree{"leaf tree|leaf tree", []*Tree{}},
						},
					},
				},
			},
			testFunc: doubleString,
		},
	}

	for _, test := range tests {
		got := *test.tree.Map(test.testFunc)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("\ngot: %v\nwant: %v\n", got, test.expected)
		}
	}
}

func TestFind(t *testing.T) {

	type result struct {
		foundTree *Tree
		isFound   bool
	}
	type test struct {
		name         string
		tree         Tree
		searchString string
		expected     result
	}

	tests := []test{
		{
			"should be found",
			oneBranchTree,
			"one branch tree",
			result{&oneBranchTree, true},
		},
		{
			"should be found",
			oneBranchTree,
			"should fail",
			result{nil, false},
		},
		{
			"should be found complex",
			Tree{"first branch", []*Tree{&twoBranchTree}},
			"leaf tree",
			result{&leafTree, true},
		},
	}

	for _, test := range tests {
		tree, isFound := test.tree.Find(test.searchString)
		got := result{
			tree, isFound,
		}
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("\ngot: %v\nwant: %v", got, test.expected)
		}
	}
}
