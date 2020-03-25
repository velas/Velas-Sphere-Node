package entity

type TaskExecutionRequest struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Input string `json:"input"`
}
