package taskdataqueries

type QueryAggParam struct {
}

func (q TaskDataQuery) DataQuery(param TaskDataParam) (*TaskDataResponse, error) {
	var data TaskDataResponse

	query := `
		SELECT
			t.task_id
			, t.task_name
			, t.task_deadline
			, t.task_details
		FROM
			tasks t
	`
	err := q.base.DbClient.Session().Debug().Raw(
		query,
	).Find(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}
