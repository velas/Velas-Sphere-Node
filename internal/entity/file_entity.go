package entity

import "github.com/sorenvonsarvort/velas-sphere/internal/merkletree"

type File struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	DecryptionKey      []byte          `json:"decryption_key"`
	Target             string          `json:"target"` // TODO: turn it into a url
	GetBackToken       []byte          `json:"get_back_token"`
	RequesterPublicKey []byte          `json:"requester_public_key"`
	Data               []byte          `json:"data"`
	MerkleTreeRoot     merkletree.Node `json:"merkle_tree_root"`
	VerificationToken  []byte          `json:"verification_token"` // TODO: turn it into a PublicKey
}
