package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	merkle "github.com/keybase/go-merkle-tree"
)

type sha256Hasher struct{}

func (sha256Hasher) Hash(in []byte) merkle.Hash {
	return sha256.New().Sum(in)
}

type valueConstructor struct{}

func (valueConstructor) Construct() interface{} {
	return nil
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func generateData(count int, t *merkle.Tree, original string) (interface{}, *merkle.SortedMap, error) {
	list := []merkle.KeyValuePair{}
	var item interface{}
	for i := 0; i < count; i++ {
		entry, err := randomHex(8)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to generate a random hex: %w", err)
		}

		entry += original

		if i == 0 {
			fmt.Println("original entry:", []byte(entry))
			item = sha256Hasher{}.Hash([]byte(entry))
		}

		list = append(
			list,
			merkle.KeyValuePair{
				Key:   sha256Hasher{}.Hash([]byte(entry)),
				Value: entry,
			},
		)

		if err != nil {
			return nil, nil, fmt.Errorf("failed to upsert an item: %w", err)
		}
	}

	return item, merkle.NewSortedMapFromSortedList(list), nil
}

func merkleTreeDemo() error {
	primaryTree := merkle.NewTree(
		merkle.NewMemEngine(),
		merkle.NewConfig(
			sha256Hasher{},
			256,
			1,
			valueConstructor{},
		),
	)

	entry, sortedDataMap, err := generateData(128, primaryTree, "test")
	if err != nil {
		return fmt.Errorf("failed to generate data: %w", err)
	}

	_ = entry

	err = primaryTree.Build(context.TODO(), sortedDataMap, nil)
	if err != nil {
		return fmt.Errorf("failed to build tree %w", err)
	}

	ret, root, err := primaryTree.Find(context.TODO(), entry.(merkle.Hash))
	if err != nil {
		return fmt.Errorf("failed to find 0 item: %w", err)
	}

	fmt.Println("found item:", ret)

	fmt.Println("found root:", root)

	return nil
}
