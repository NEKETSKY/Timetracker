package dbrepository

type Group struct {
	GroupID    int    `json:"id"`
	GroupTitle string `json:"title"`
	Task       []Task `json:"tasks"`
}
type Task struct {
	TaskID     int         `json:"id"`
	TaskTitle  string      `json:"title"`
	GroupID    int         `json:"group_id"`
	Timeframes []Timeframe `json:"time_frames"`
}

type Timeframe struct {
	TaskID    int    `json:"task_id"`
	Time_FROM string `json:"from"`
	Time_TO   string `json:"to"`
}
