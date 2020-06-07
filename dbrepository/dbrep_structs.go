package dbrepository

type Group struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Task  []Task `json:"tasks,omitempty"`
}
type Task struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	GroupID    int         `json:"group_id"`
	Timeframes []Timeframe `json:"time_frames,omitempty"`
}

type Timeframe struct {
	TaskID   int    `json:"task_id"`
	TimeFrom string `json:"from"`
	TimeTo   string `json:"to"`
}

type GroupsResponse struct {
	Groups []Group `json:"groups"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}
