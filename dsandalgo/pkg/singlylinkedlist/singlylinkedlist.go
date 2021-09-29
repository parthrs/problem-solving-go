package singlylinkedlist

type SinglyNode struct {
	Value int
	Next  *SinglyNode
}

func NewSinglyNode(val int) *SinglyNode {
	return &SinglyNode{
		Value: val,
		Next:  nil,
	}
}
