package dbrepository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type TaskRepositorySQL struct {
	DB *sql.DB
}

// GetGroups - allows you to get all groups from the database
func (repo TaskRepositorySQL) GetGroups() (GroupsResponse, error) {
	var groups GroupsResponse
	groupsTable, err := repo.DB.Query("SELECT * FROM groups ORDER BY group_id")
	if err != nil {
		return groups, err
	}
	defer groupsTable.Close()

	for groupsTable.Next() {
		var group Group
		err := groupsTable.Scan(&group.ID, &group.Title)
		if err != nil {
			return groups, err
		}
		group.Task, err = repo.getTasksByGroupID(group.ID)
		if err != nil {
			return groups, err
		}
		groups.Groups = append(groups.Groups, group)
	}
	return groups, nil
}

// CreateGroup - allows you to create a new group in database
func (repo TaskRepositorySQL) CreateGroup(group Group) (Group, error) {
	var createdGroup Group
	err := repo.DB.QueryRow("INSERT INTO groups(title) VALUES($1) RETURNING group_id, title", group.Title).Scan(&createdGroup.ID, &createdGroup.Title)
	if err != nil {
		return createdGroup, err
	}
	return createdGroup, nil
}

// UpdateGroup - allows you to update an existing group in the database by ID
func (repo TaskRepositorySQL) UpdateGroup(group Group) (Group, error) {
	var updatedGroup Group
	err := repo.DB.QueryRow("UPDATE groups SET title=$1 WHERE group_id=$2 RETURNING group_id, title", group.Title, group.ID).Scan(&updatedGroup.ID, &updatedGroup.Title)
	if err != nil {
		return updatedGroup, err
	}
	return updatedGroup, nil
}

// DeleteGroup - allows to delete an existing group from database by ID
func (repo TaskRepositorySQL) DeleteGroup(id int) error {
	_, err := repo.DB.Exec("DELETE FROM groups WHERE group_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// GetTasks - allows you to get all tasks from the database
func (repo TaskRepositorySQL) GetTasks() (TasksResponse, error) {
	var tasks TasksResponse
	tasksTable, err := repo.DB.Query("SELECT * FROM tasks ORDER BY task_id")
	if err != nil {
		return tasks, err
	}
	defer tasksTable.Close()

	for tasksTable.Next() {
		var task Task
		err := tasksTable.Scan(&task.ID, &task.Title, &task.GroupID)
		if err != nil {
			return tasks, err
		}
		task.Timeframes, err = repo.getTimeframesByTaskID(task.ID)
		if err != nil {
			return tasks, err
		}
		tasks.Tasks = append(tasks.Tasks, task)
	}
	return tasks, nil
}

// CreateTask - allows you to create a new task in the database
func (repo TaskRepositorySQL) CreateTask(task Task) (Task, error) {
	var createdTask Task
	err := repo.DB.QueryRow("INSERT INTO tasks(title, group_id) VALUES($1,$2) RETURNING task_id, title, group_id", task.Title, task.GroupID).Scan(&createdTask.ID, &createdTask.Title, &createdTask.GroupID)
	if err != nil {
		return createdTask, err
	}
	return createdTask, nil
}

// UpdateTask - allows you to update an existing task in the database by ID
func (repo TaskRepositorySQL) UpdateTask(task Task) (Task, error) {
	var updatedTask Task
	err := repo.DB.QueryRow("UPDATE tasks SET title=$1, group_id=$2 WHERE task_id=$3 RETURNING task_id, title, group_id;", task.Title, task.GroupID, task.ID).Scan(&updatedTask.ID, &updatedTask.Title, &updatedTask.GroupID)
	if err != nil {
		return updatedTask, err
	}
	return updatedTask, nil
}

// DeleteTask - allows to delete an existing task from database by ID
func (repo TaskRepositorySQL) DeleteTask(id int) error {
	_, err := repo.DB.Exec("DELETE FROM tasks WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// CreateTimeframe - allows you to create a new timeframe in database
func (repo TaskRepositorySQL) CreateTimeframe(timeframe Timeframe) (Timeframe, error) {
	var createdTimeframe Timeframe
	err := repo.DB.QueryRow("INSERT INTO timeframes(task_id, time_start, time_end) VALUES ($1,$2, $3) RETURNING task_id, time_start, time_end", timeframe.TaskID, timeframe.TimeFrom, timeframe.TimeTo).Scan(&createdTimeframe.TaskID, &createdTimeframe.TimeFrom, &createdTimeframe.TimeTo)
	if err != nil {
		return createdTimeframe, err
	}
	return createdTimeframe, nil
}

// DeleteTimeframe - allows to delete all timeframes from database by ID
func (repo TaskRepositorySQL) DeleteTimeframe(id int) error {
	_, err := repo.DB.Exec("DELETE FROM timeframes WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// getTasksByGroupID - allows to get all tasks from database by group ID
func (repo TaskRepositorySQL) getTasksByGroupID(id int) ([]Task, error) {
	var tasks []Task
	taskTable, err := repo.DB.Query("SELECT * FROM tasks WHERE group_id=$1 ORDER BY task_id", id)
	if err != nil {
		return tasks, err
	}

	defer taskTable.Close()

	for taskTable.Next() {
		var task Task
		err := taskTable.Scan(&task.ID, &task.Title, &task.GroupID)
		if err != nil {
			return tasks, err
		}
		task.Timeframes, err = repo.getTimeframesByTaskID(task.ID)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// getTimeframesByTaskID - allows to get all timeframes from database by task ID
func (repo TaskRepositorySQL) getTimeframesByTaskID(id int) ([]Timeframe, error) {
	var timeframes []Timeframe
	timeframesTable, err := repo.DB.Query("SELECT * FROM timeframes WHERE task_id=$1", id)
	if err != nil {
		return timeframes, err
	}
	defer timeframesTable.Close()

	for timeframesTable.Next() {
		var timeframe Timeframe
		err := timeframesTable.Scan(&timeframe.TaskID, &timeframe.TimeFrom, &timeframe.TimeTo)
		if err != nil {
			return timeframes, err
		}
		timeframes = append(timeframes, timeframe)
	}
	return timeframes, err
}


//CheckGroupByID - allows you to check the existence of a group by ID
func (repo TaskRepositorySQL) CheckGroupByID(group Group) error{
	var id int
	 err := repo.DB.QueryRow("SELECT group_id FROM groups WHERE group_id=$1", group.ID).Scan(&id)
	 if err!= nil {
		return err
	}
	return nil
}

//CheckGroupByID -  allows you to check the presence of a task by ID and the validity of the group ID
func (repo TaskRepositorySQL) CheckTaskByID(task Task) error{
	var stub int
	err := repo.DB.QueryRow("SELECT task_id FROM tasks WHERE task_id=$1", task.ID).Scan(&stub)
	if err!= nil {
		return err
	}
	err = repo.DB.QueryRow("SELECT group_id FROM tasks WHERE group_id=$1", task.GroupID).Scan(&stub)
	if err!= nil {
		return err
	}
	return nil
}
