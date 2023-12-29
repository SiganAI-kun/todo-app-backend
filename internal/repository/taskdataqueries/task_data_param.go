package taskdataqueries

type TaskDataParam struct {
	Start       string
	End         string
	SearchQuery string
}

func NewGetTaskDataParam(
	start string,
	end string,
	searchQuery string,
) (*TaskDataParam, error) {
	p := &TaskDataParam{
		Start:       start,
		End:         end,
		SearchQuery: searchQuery,
	}

	return p, nil
}
