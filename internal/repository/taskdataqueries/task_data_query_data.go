package taskdataqueries

type QueryAggParam struct {
}

// すべてのタスクデータの取得
func (q TaskDataQuery) GetDataQuery(param GetTaskDataParam) ([]TaskDataResponse, error) {
	var data []TaskDataResponse

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

	return data, nil
}

func (q TaskDataQuery) CreateDataQuery(param CreateTaskDataParam) (error, error) {
	query := `
	INSERT INTO 
		tasks
	(
		task_name
		, task_deadline
		, task_details
	)
	VALUES
	(
		@task_name
		, @task_deadline
		, @task_details
	)
	`

	result := q.base.DbClient.Session().Debug().Exec(
		query,
		map[string]interface{}{
			"task_name":     param.TaskName,
			"task_deadline": param.TaskDeadline,
			"task_details":  param.TaskDetails,
		},
	)

	if result.Error != nil {
		return nil, result.Error
	}

	// 何らかの成功時の処理（ここでは新規追加されたデータの ID などが必要な場合に処理を追加）

	return nil, nil
}

func (q TaskDataQuery) UpdateDataQuery(param UpdateTaskDataParam) (error, error) {
	query := `
	UPDATE tasks
	SET 
		task_name = @task_name,
		task_deadline = @task_deadline,
		task_details = @task_details,
		updated_at = NOW()
	WHERE 1  =1
		AND task_id = @task_id;
	`

	result := q.base.DbClient.Session().Debug().Exec(
		query,
		map[string]interface{}{
			"task_id":       param.TaskId,
			"task_name":     param.TaskName,
			"task_deadline": param.TaskDeadline,
			"task_details":  param.TaskDetails,
		},
	)

	if result.Error != nil {
		return nil, result.Error
	}

	// 何らかの成功時の処理（ここでは新規追加されたデータの ID などが必要な場合に処理を追加）

	return nil, nil
}

func (q TaskDataQuery) CheckSameTasksDataQuery(param CreateTaskDataParam) (bool, error) {
	var data TaskDataResponse
	query := `
		SELECT
			t.task_id
			, t.task_name
			, t.task_deadline
			, t.task_details
		FROM
			tasks t
		WHERE 1 = 1
			AND t.task_name = @task_name
			AND t.task_deadline = @task_deadline
			AND t.task_details = @task_details
		LIMIT
			1
	`

	err := q.base.DbClient.Session().Debug().Raw(
		query,
		map[string]interface{}{
			"task_name":     param.TaskName,
			"task_deadline": param.TaskDeadline,
			"task_details":  param.TaskDetails,
		},
	).Find(&data).Error

	// エラーが nil かつデータが見つかった場合
	if err == nil && data.TaskId != 0 {
		// データが見つかった場合の処理
		return true, nil
	}

	// データが見つからなかった場合の処理
	return false, nil
}

func (q TaskDataQuery) CheckTaskExistDataQuery(param UpdateTaskDataParam) (bool, error) {
	var data TaskDataResponse
	query := `
		SELECT
			t.task_id
			, t.task_name
			, t.task_deadline
			, t.task_details
		FROM
			tasks t
		WHERE 1 = 1
			AND t.task_id = @task_id
		LIMIT
			1
	`

	err := q.base.DbClient.Session().Debug().Raw(
		query,
		map[string]interface{}{
			"task_id": param.TaskId,
		},
	).Find(&data).Error

	// エラーが nil かつデータが見つかった場合
	if err == nil && data.TaskId != 0 {
		// データが見つかった場合の処理
		return true, nil
	}

	// データが見つからなかった場合の処理
	return false, nil
}
