package dbrepository

type (
	RepositorySQL interface {
		GetGroups() (GroupsResponse, error)
		CreateGroup(group Group) (Group, error)
		UpdateGroup(group Group) (Group, error)
		DeleteGroup(id int) error
		GetTasks() (TasksResponse, error)
		CreateTask(task Task) (Task, error)
		UpdateTask(task Task) (Task, error)
		DeleteTask(id int) error
		CreateTimeframe(timeframe Timeframe) (Timeframe, error)
		DeleteTimeframe(id int) error
		getTasksByGroupID(id int) ([]Task, error)
		getTimeframesByTaskID(id int) ([]Timeframe, error)
		checkGroupByID(group Group) error
		checkTastByID(task Task) error
	}

	Group struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Task  []Task `json:"tasks,omitempty"`
	}

	Task struct {
		ID         int         `json:"id"`
		Title      string      `json:"title"`
		GroupID    int         `json:"group_id"`
		Timeframes []Timeframe `json:"time_frames,omitempty"`
	}

	Timeframe struct {
		TaskID   int    `json:"task_id"`
		TimeFrom string `json:"from"`
		TimeTo   string `json:"to"`
	}

	GroupsResponse struct {
		Groups []Group `json:"groups"`
	}

	TasksResponse struct {
		Tasks []Task `json:"tasks"`
	}
)
