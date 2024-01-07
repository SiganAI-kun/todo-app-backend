package taskdataservices

type GetTaskDataParam struct {
	IsFilterd   bool   `json:"isFilterd"`
	Start       string `json:"start"`
	End         string `json:"end"`
	SearchQuery string `json:"searchQuery"`
}

type CreateTaskDataParam struct {
	TaskName     string `json:"taskName"`
	TaskDeadline string `json:"taskDeadline"`
	TaskDetails  string `json:"taskDetails"`
}

type UpdateTaskDataParam struct {
	TaskId       int    `json:"taskId"`
	TaskName     string `json:"taskName"`
	TaskDeadline string `json:"taskDeadline"`
	TaskDetails  string `json:"taskDetails"`
}

type DeleteTaskDataParam struct {
	TaskId int `json:"taskId"`
}
