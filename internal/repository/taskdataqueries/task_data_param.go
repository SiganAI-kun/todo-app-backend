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

type UpdateTaskDataParam struct {
	TaskId       int
	TaskName     string
	TaskDeadline string
	TaskDetails  string
}

type DeleteTaskDataParam struct {
	TaskId int
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

func NewUpdateTaskDataParam(
	taskId int,
	taskName string,
	taskDeadline string,
	taskDetails string,
) (*UpdateTaskDataParam, error) {
	p := &UpdateTaskDataParam{
		TaskId:       taskId,
		TaskName:     taskName,
		TaskDeadline: taskDeadline,
		TaskDetails:  taskDetails,
	}

	return p, nil
}

func NewDeleteTaskDataParam(
	taskId int,
) (*DeleteTaskDataParam, error) {
	p := &DeleteTaskDataParam{
		TaskId: taskId,
	}

	return p, nil
}
