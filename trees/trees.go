package trees

import "strings"

type Tree struct {
	label    string
	branches []*Tree
}

func (tree *Tree) PrintTree(tabString string) string {

	var builder strings.Builder

	currLine := tabString + tree.getTreeLabel()
	builder.WriteString(currLine)
	// base case, if no branches, just return

	for _, branch := range tree.branches {
		builder.WriteString("\n")
		builder.WriteString(branch.PrintTree(tabString + "\t"))
	}
	return builder.String()
}

func (tree *Tree) getTreeLabel() string {
	return tree.label
}

func (tree *Tree) CountLeaves() int {

	if tree.isLeaf() {
		return 1
	}
	numLeaves := 0
	for _, branch := range tree.branches {
		numLeaves += branch.CountLeaves()
	}
	return numLeaves
}

func (tree *Tree) Contains(searchString string) bool {
	// default return val to false

	if tree.label == searchString {
		return true
	}
	for _, branch := range tree.branches {
		if branch.Contains(searchString) {
			return true
		}
	}
	return false
}

func (tree *Tree) Depth() int {
	curDepth := 1
	if tree.isLeaf() {
		return curDepth
	}
	maxAdditionalDepth := 0
	for _, branch := range tree.branches {
		branchDepth := branch.Depth()
		if branchDepth > maxAdditionalDepth {
			maxAdditionalDepth = branchDepth
		}
	}
	return curDepth + maxAdditionalDepth
}

func (tree *Tree) Map(inputFunc func(string) string) *Tree {
	returnedTree := &Tree{
		inputFunc(tree.label),
		[]*Tree{},
	}
	if tree.isLeaf() {
		return returnedTree
	}
	for _, branch := range tree.branches {
		returnedFromBranch := branch.Map(inputFunc)
		returnedTree.branches = append(returnedTree.branches, returnedFromBranch)
	}
	return returnedTree
}

func (tree *Tree) Find(searchString string) (foundTree *Tree, isFound bool) {
	if tree.label == searchString {
		return tree, true
	}
	for _, branch := range tree.branches {
		currTree, currFound := branch.Find(searchString)
		if currFound {
			return currTree, true
		}
	}
	return nil, false
}

func (tree *Tree) isLeaf() bool {
	return len(tree.branches) == 0
}

// if i want to do thish recursively, for every branch, i would
// recursively move downwards until i reach a tree that has no brances
