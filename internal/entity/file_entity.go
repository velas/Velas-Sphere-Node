package entity

type File struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	DecryptionKey      string `json:"decryption_key"`
	Target             string `json:"target"`
	GetBackToken       string `json:"get_back_token"`
	RequesterPublicKey string `json:"requester_public_key"`
	Data               []byte `json:"data"`
	MerkleTreeRoot     string `json:"merkle_tree_root"`
	VerificationToken  string `json:"verification_token"`
}
