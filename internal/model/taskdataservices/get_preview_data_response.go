package taskdataservices

type TaskDataResponse struct {
	TaskId       uint8  `json:"taskId"`
	TaskName     string `json:"taskName"`
	TaskDeadline string `json:"taskDeadline"`
	TaskDetails  string `json:"taskDetails"`
}
