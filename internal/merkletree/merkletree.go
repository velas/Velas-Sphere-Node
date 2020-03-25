package merkletree

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

type Node struct {
	Value    string
	Children [2]interface{}
}

func hash(s string) string {
	x := sha256.New()
	x.Write([]byte(s))
	return string(x.Sum(nil))
}

func CalculatePair(a, b Node) Node {
	return Node{
		Value:    hash(a.Value + b.Value),
		Children: [2]interface{}{a, b},
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
		childNode, ok := child.(Node)
		newPrefix := prefix + "-"
		if ok {
			PrintNode(childNode, newPrefix)
		} else {
			fmt.Printf("%sX\n", newPrefix)
		}
	}
}

func findPath(treeRoot Node, nodeHash string, prefix []string) ([]string, error) {
	path := append(prefix, hex.EncodeToString([]byte(treeRoot.Value)))

	if treeRoot.Value == nodeHash {
		return path, nil
	}

	for _, child := range treeRoot.Children {
		childNode, ok := child.(Node)
		if !ok {
			continue
		}

		newPath, err := findPath(childNode, nodeHash, path)
		if err != nil {
			continue
		}
		return newPath, nil
	}

	return nil, errors.New("no such hash found")
}

func FindPath(treeRoot Node, nodeHash string) ([]string, error) {
	return findPath(treeRoot, nodeHash, nil)
}
