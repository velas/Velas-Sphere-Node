package merkletree

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

type Node struct {
	Value    []byte
	Children [2]*Node
}

func hash(s string) []byte {
	x := sha256.New()
	x.Write([]byte(s))
	return x.Sum(nil)
}

func VerifyNode(root Node) bool {
	if root.Children[0] != nil || root.Children[1] != nil {
		reconstructedRoot := CalculatePair(*root.Children[0], *root.Children[1])
		if bytes.Compare(reconstructedRoot.Value, root.Value) == 0 {
			for _, child := range root.Children {
				// we did already check for nil
				return VerifyNode(*child)
			}
		}
	}

	return true
}

func CalculatePair(a, b Node) Node {
	return Node{
		Value:    []byte(hash(string(a.Value) + string(b.Value))),
		Children: [2]*Node{&a, &b},
	}
}

func CalculatePairs(nodes []Node) []Node {
	pairs := []Node{}

	for i := 0; i < len(nodes); i += 2 {
		pair := Node{}

		if i+1 >= len(nodes) {
			pair = CalculatePair(
				nodes[i],
				nodes[i],
			)
		} else {
			pair = CalculatePair(
				nodes[i],
				nodes[i+1],
			)
		}

		pairs = append(pairs, pair)
	}

	return pairs
}

func CalculateRoot(items ...string) (Node, error) {
	if len(items) < 1 {
		return Node{}, errors.New("invalid items count")
	}

	nodes := []Node{}

	for _, item := range items {
		nodes = append(
			nodes,
			Node{
				Value: hash(item),
			},
		)
	}

	pairs := nodes
	for len(pairs) > 1 {
		pairs = CalculatePairs(pairs)
	}

	return pairs[0], nil
}

func PrintNode(n Node, prefix string) {
	fmt.Printf("%s[%s]\n", prefix, hex.EncodeToString([]byte(n.Value)))
	for _, child := range n.Children {
		newPrefix := prefix + "-"
		if child == nil {
			fmt.Printf("%sX\n", newPrefix)
			continue
		}

		PrintNode(*child, newPrefix)
	}
}

func findPath(treeRoot Node, nodeHash []byte, prefix []string) ([]string, error) {
	path := append(prefix, hex.EncodeToString([]byte(treeRoot.Value)))

	if bytes.Compare(treeRoot.Value, nodeHash) == 0 {
		return path, nil
	}

	for _, child := range treeRoot.Children {
		if child == nil {
			continue
		}

		newPath, err := findPath(*child, nodeHash, path)
		if err != nil {
			continue
		}

		return newPath, nil
	}

	return nil, errors.New("no such hash found")
}

func FindPath(treeRoot Node, nodeHash []byte) ([]string, error) {
	return findPath(treeRoot, nodeHash, nil)
}
