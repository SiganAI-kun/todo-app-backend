package taskdataservices

import (
	"testing"

	db "todo-app-backend/db"
	"todo-app-backend/internal/repository/taskdataqueries"

	"github.com/stretchr/testify/assert"
)

// Mock for ITaskDataQuery
type MockTaskDataQuery struct {
}


func (m *MockTaskDataQuery) GetDataQuery(param taskdataqueries.GetTaskDataParam) ([]taskdataqueries.TaskDataResponse, error) {
    // ここではテスト用のダミーデータを返す例
    // 実際のテストでは、適切なテストデータやモックデータを返すように実装すること
    return []taskdataqueries.TaskDataResponse{
        {TaskId: 1, TaskName: "Task 1", TaskDetails: "Description 1", TaskDeadline: "2999-01-01"},
        {TaskId: 2, TaskName: "Task 2", TaskDetails: "Description 2", TaskDeadline: "2999-01-02"},
    }, nil
}

func TestGetPreviewTaskDataService_Exec(t *testing.T) {
	// モックのデータベースコンテキスト
	mockDbContext := &db.DbContext{}

	// モックのクエリ
	mockQuery := &MockTaskDataQuery{}

	// サービスの初期化
	service := NewGetPreviewDataService(mockDbContext, mockQuery)

	// テスト用のパラメータ
	param := GetTaskDataParam{
		Start:       "2000-01-01",
		End:         "3000-01-01",
		SearchQuery: "",
	}

	// サービスの実行
	result, err := service.Exec(param)

	// アサーション
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// 結果の型を確認（実際の型に合わせて修正すること）
	_, ok := result.([]taskdataqueries.TaskDataResponse)
	assert.True(t, ok)
}

// 新しく追加したメソッド
func (m *MockTaskDataQuery) CheckSameTasksDataQuery(param taskdataqueries.CreateTaskDataParam) (bool, error) {
    // 実際のコードを追加
    return true, nil
}

func (m *MockTaskDataQuery) CheckTaskExistDataQuery(param taskdataqueries.DeleteTaskDataParam) (bool, error) {
	// 実際のコードを追加
	return true, nil
}

func (m *MockTaskDataQuery) CreateDataQuery(param taskdataqueries.CreateTaskDataParam) (error, error) {
	// 実際のコードを追加
	return nil, nil
}

func (m *MockTaskDataQuery) DeleteDataQuery(param taskdataqueries.DeleteTaskDataParam) (error, error) {
	// 実際のコードを追加
	return nil, nil
}

func (m *MockTaskDataQuery) UpdateDataQuery(param taskdataqueries.UpdateTaskDataParam) (error, error) {
	// 実際のコードを追加
	return nil, nil
}