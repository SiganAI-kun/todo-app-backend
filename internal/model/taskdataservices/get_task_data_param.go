package taskdataservices

type GetTaskDataParam struct {
	IsFilterd   bool   `json:"isFilterd"`
	Start       string `json:"start"`
	End         string `json:"end"`
	SearchQuery string `json:"searchQuery"`
}
