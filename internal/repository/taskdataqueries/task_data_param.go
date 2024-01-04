package taskdataqueries

type GetTaskDataParam struct {
	Start       string
	End         string
	SearchQuery string
}

type CreateTaskDataParam struct {
	TaskName     string
	TaskDeadline string
	TaskDetails  string
}

func NewGetTaskDataParam(
	start string,
	end string,
	searchQuery string,
) (*GetTaskDataParam, error) {
	p := &GetTaskDataParam{
		Start:       start,
		End:         end,
		SearchQuery: searchQuery,
	}

	return p, nil
}

func NewCreateTaskDataParam(
	taskName string,
	taskDeadline string,
	taskDetails string,
) (*CreateTaskDataParam, error) {
	p := &CreateTaskDataParam{
		TaskName:     taskName,
		TaskDeadline: taskDeadline,
		TaskDetails:  taskDetails,
	}

	return p, nil
}
