package taskdataqueries

import (
	"database/sql/driver"
	"strconv"
	"strings"
)

type TaskDataResponse struct {
	TaskId       uint   `json:"taskId"` // uintまたはuint64に変更
	TaskName     string `json:"taskName"`
	TaskDeadline string `json:"taskDeadline"`
	TaskDetails  string `json:"taskDetails"`
}

// type TaskDataResponses struct {
// 	Data []TaskDataResponse `json:"data" gorm:"foreignkey:TaskId;references:ID"`
// }

// Valuer/Scanner インターフェースの実装
func (t *TaskDataResponse) Value() (driver.Value, error) {
	// データベースに保存するための値を返す
	return []byte(strings.Join([]string{strconv.FormatUint(uint64(t.TaskId), 10), t.TaskName, t.TaskDeadline, t.TaskDetails}, ",")), nil
}

func (t *TaskDataResponse) Scan(value interface{}) error {
	// データベースからの値を構造体にスキャンする処理
	parts := strings.Split(string(value.([]byte)), ",")
	parsedTaskId, err := strconv.ParseUint(parts[0], 10, 64) // uint64に変更
	if err != nil {
		return err
	}

	t.TaskId = uint(parsedTaskId) // uintに変更
	t.TaskName = parts[1]
	t.TaskDeadline = parts[2]
	t.TaskDetails = parts[3]
	return nil
}
