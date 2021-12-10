package day03

type binaryNode struct {
	numZeroChild int
	numOneChild int
	zeroChild *binaryNode
	oneChild *binaryNode
}

func newBinaryNode() binaryNode {
	return binaryNode{
		numZeroChild:   -1,
		numOneChild:    -1,
		zeroChild:      nil,
		oneChild:       nil,
	}
}

type binaryTree struct {
	head *binaryNode
}

func (t *binaryTree) NumChildren() int {
	if t.head == nil {
		return 0
	}
	return t.head.NumChildren()
}

func newbinaryTree(lines []string) binaryTree {
	if len(lines) == 0 {
		return binaryTree{}
	}
	// The top node has a -1 to indicate it's just a root value
	node := newBinaryNode()
	tree := binaryTree{head: &node}
	dedupedLines := map[string]bool{}
	for _, line := range lines {
		dedupedLines[line] = true
	}
	for line, _ := range dedupedLines {
		tree.head.Add(line)
	}
	return tree
}

func (b *binaryNode) HasChildren() bool {
	return b.zeroChild != nil || b.oneChild != nil
}

func (b *binaryNode) MostCommonChild() (string, *binaryNode) {
	if b.zeroChild == nil && b.oneChild == nil {
		return "", nil
	}
	if b.NumZeroChildren() > b.NumOneChildren() {
		return "0", b.zeroChild
	}
	return "1", b.oneChild
}

func (b *binaryNode) LeastCommonChild() (string, *binaryNode) {
	if b.zeroChild == nil && b.oneChild == nil {
		return "", nil
	}
	if b.zeroChild == nil {
		return "1", b.oneChild
	}
	if b.oneChild == nil {
		return "0", b.zeroChild
	}
	if b.NumZeroChildren() <= b.NumOneChildren() {
		return "0", b.zeroChild
	}
	return "1", b.oneChild
}

func (b *binaryNode) NumZeroChildren() int {
	if b.numZeroChild != -1 {
		return b.numZeroChild
	}
	if b.zeroChild == nil {
		b.numZeroChild = 0
	} else {
		b.numZeroChild = b.zeroChild.NumChildren()
	}
	return b.numZeroChild
}

func (b *binaryNode) NumOneChildren() int {
	if b.numOneChild != -1 {
		return b.numOneChild
	}
	if b.oneChild == nil {
		b.numOneChild = 0
	} else {
		b.numOneChild = b.oneChild.NumChildren()
	}
	return b.numOneChild
}

func (b *binaryNode) NumChildren() int {
	if b.zeroChild == nil && b.oneChild == nil {
		return 1
	}
	left := 0
	if b.zeroChild != nil {
		left = b.zeroChild.NumChildren()
	}
	right := 0
	if b.oneChild != nil {
		right = b.oneChild.NumChildren()
	}
	return right + left
}

func (b *binaryNode) Add(input string) {
	if len(input) == 0 {
		return
	}
	// This node will have a new descendant
	var nextChild *binaryNode
	if input[0] == '0' {
		if b.zeroChild == nil {
			node := newBinaryNode()
			b.zeroChild = &node
		}
		nextChild = b.zeroChild
	} else {
		if b.oneChild == nil {
			node := newBinaryNode()
			b.oneChild = &node
		}
		nextChild = b.oneChild
	}
	nextChild.Add(input[1:])
}
