package entity

type TaskExecutionRequest struct {
	Target string `json:"target"`
	ID     string `json:"id"`
	Input  string `json:"input"`
}
