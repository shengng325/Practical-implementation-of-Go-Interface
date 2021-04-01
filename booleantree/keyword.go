package boolTree

type Node interface {
	Eval(input Sentence) bool
}

// Ensure all nodes implement the Node interface
var _ Node = &ANDNode{}
var _ Node = &ORNode{}
var _ Node = &NOTNode{}
var _ Node = Word("")

// AND Implementation
func AND(n ...Node) Node {
	return &ANDNode{n}
}

type ANDNode struct {
	Nodes []Node
}

func (a *ANDNode) Eval(input Sentence) bool {
	result := true
	for i := range a.Nodes {
		result = result && a.Nodes[i].Eval(input)
	}
	return result
}

// OR Implementation
func OR(n ...Node) Node {
	return &ORNode{n}
}

type ORNode struct {
	Nodes []Node
}

func (o *ORNode) Eval(input Sentence) bool {
	result := false
	for i := range o.Nodes {
		result = result || o.Nodes[i].Eval(input)
	}
	return result
}

// NOT Implementation
func NOT(n Node) Node {
	return &NOTNode{n}
}

type NOTNode struct {
	Node Node
}

func (n *NOTNode) Eval(input Sentence) bool {
	return !n.Node.Eval(input)
}

// Keyword implementation
type Word string

func (w Word) Eval(input Sentence) bool {
	if _, ok := input[string(w)]; ok {
		return true
	}
	return false
}
