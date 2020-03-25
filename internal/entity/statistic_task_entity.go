package entity

import "fmt"

type TaskStatistic struct {
	Type     string `json:"type"`
	Status   string `json:"status"`
	Finished string `json:"finished"`
	Spent    string `json:"spent"`
}

const (
	TaskStatusCreated    = "created"
	TaskStatusInProgress = "in_progress"
	TaskStatusFinished   = "finished"
)

const TaskIDPrefix = "task_id_"

func TaskIDFromString(id string) string {
	return fmt.Sprintf("%s%s", TaskIDPrefix, id)
}

func TaskIDFromInt(id int64) string {
	return fmt.Sprintf("%s%d", TaskIDPrefix, id)
}

type TaskStatisticListResponse struct {
	Data []TaskStatistic `json:"data"`
}
