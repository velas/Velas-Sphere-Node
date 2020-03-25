package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/sorenvonsarvort/velas-sphere/internal/merkletree"
)

// TODO: return map with proofs
func store(data string) (merkletree.Node, error) {
	challenges := []string{
		"x",
		"w",
		"z",
		"j",
	}

	items := []string{}

	for _, challenge := range challenges {
		items = append(items, data+challenge)
	}

	return merkletree.CalculateRoot(items...)
}

func proove(challenge, data string, root merkletree.Node, expectedPath string) bool {
	hash := sha256.New()
	hash.Write([]byte(data + challenge))
	hashBytes := hash.Sum(nil)

	path, err := merkletree.FindPath(root, string(hashBytes))
	if err != nil {
		return false
	}

	return expectedPath == strings.Join(path, ":")
}

func simpleDemo() {
	root, err := merkletree.CalculateRoot("one", "two", "three", "four")
	if err != nil {
		log.Println("failed to calculate root:", err)
		return
	}

	merkletree.PrintNode(root, "> ")

	hash, err := hex.DecodeString("3fc4ccfe745870e2c0d99f71f30ff0656c8dedd41cc1d7d3d376b0dbe685e2f3")
	if err != nil {
		log.Println("invalid hash provided:", err)
		return
	}

	path, err := merkletree.FindPath(root, string(hash))
	if err != nil {
		log.Println("failed to find the hash:", err)
		return
	}

	log.Println("found path:", strings.Join(path, ":"))
	return
}

func storageVerificationDemo() {
	root, err := store("mydata")
	if err != nil {
		log.Println("failed to store:", err)
	}

	merkletree.PrintNode(root, "")

	fmt.Println(
		proove(
			"w",
			"mydata",
			root,
			"629708d0d6ddf0551d9e93e068824adceda09a0c994ad3630c3b7f40321a43ac:a22c668f1d45f16f8b6d830d4bd0a4ce0088b7040e3afaeba440ff93c68ac9c8:6a8a23a38d8935de2f9e6ce7b6adc4b573580ffc0a156398e3b1f961872b5071",
		),
	)
}
