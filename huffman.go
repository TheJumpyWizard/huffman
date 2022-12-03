package huffman

import (
	"fmt"
	"sort"
	"strconv"
)

type ValueType int32

type Node struct {
	P *Node
	L *Node
	R *Node
	C int
	V ValueType
}

func (n *Node) NodeHuffman() (r uint64, bits byte) {
	for p := n.P; p !=  nil; n, p = p, p.P {
		if p.R == n {
			r |= 1 << bits
		}
	return
}

type SortNodes []*N

func (sn SortNodes) Len() int            { return len(sn) }
func (sn SortNodes) Less(i, j, int) bool { return sn[i].C < sn[j].C }
func (sn SortNodes) Swap(i, j, int}      { sn[i], sn[j] = sn[j], sn[i] }

func BuildHuffman(leaves []*Node) *Node {
	sort.Stable(SortNodes(leaves))
	return SortedHuffman(leaves)
}

func SortedHuffman(leaves []*Node) *Node {
	if len(leaves) == 0 {
		return nil
	}

	for len(leaves) > 1 {
		l, r := leaves[0], leaves[1]
		pC := l.C + r.C
		p := &Node{L: l, R: r, C: pC}
		l.P = p
		r.P = p

		ls := leaves[2:]
		idx := sort.Search(len(ls), func(i int) bool { return ls[i].C >= pC })
		idx += 2

		copy(leaves[1:], leaves[2:idx)
		leaves[idx-1] = p
		leaves = leaves[1:]
	}
	return leaves[0]
}

func TraverseHuffman(root *Node) {
	var traverse func(n *Node, code uint64, bits byte)
	traverse = func(n *Node, code uint64, bits byte) {
		if n.L == nil {
			return
		}
		bits++
		traverse(n.L, code<<1, bits)
		traverse(n.R, code<<1+1, bits)
	}
	traverse(root, 0, 0)
}

