package dbrepository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var dbase *sql.DB

// DBInit - создает соединение с базой данных
func DBInit() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbName := os.Getenv("db_name")
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	log.Println(dbUri)

	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err.Error())
	}
	dbase = db
	return db
}

/*
// возвращает дескриптор объекта DB
func GetDB() *sql.DB {
	return dbase
}
*/

// GetGroups - получает из БД все группы и возвращает сформированный срез групп
func GetGroups() (GroupsResponse, error) {

	var groups GroupsResponse

	groupsTable, err := dbase.Query("SELECT * FROM groups ORDER BY group_id")
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
		group.Task, err = getTasksByGroupID(group.ID)
		if err != nil {
			return groups, err
		}
		groups.Groups = append(groups.Groups, group)
	}
	return groups, nil
}

// CreateGroup - создает в БД новую группу с заданными значениями и возвращает эту запись
func CreateGroup(group Group) (Group, error) {
	var createdGroup Group
	err := dbase.QueryRow("INSERT INTO groups(title) VALUES($1) RETURNING group_id, title", group.Title).Scan(&createdGroup.ID, &createdGroup.Title)
	if err != nil {
		return createdGroup, err
	}
	return createdGroup, nil
}

// UpdateGroup - обновляет группу в БД по ID полученными значениями и возвращает обновленную запись
func UpdateGroup(group Group) (Group, error) {
	var updatedGroup Group
	err := dbase.QueryRow("UPDATE groups SET title=$1 WHERE group_id=$2 RETURNING group_id, title", group.Title, group.ID).Scan(&updatedGroup.ID, &updatedGroup.Title)
	if err != nil {
		return updatedGroup, err
	}
	return updatedGroup, nil
}

// DeleteGroup - удаляет из БД группу по полученному ID
func DeleteGroup(id int) error {
	_, err := dbase.Exec("DELETE FROM groups WHERE group_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// GetTasks - получает из БД все таски и возвращает сформированный срез тасков
func GetTasks() (TasksResponse, error) {

	var tasks TasksResponse

	tasksTable, err := dbase.Query("SELECT * FROM tasks ORDER BY task_id")
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
		task.Timeframes, err = getTimeframesByTaskID(task.ID)
		if err != nil {
			return tasks, err
		}
		tasks.Tasks = append(tasks.Tasks, task)
	}
	return tasks, nil
}

// CreateTask - создает в БД новую таску с заданными значениями и возвращает эту запись
func CreateTask(task Task) (Task, error) {
	var createdTask Task
	err := dbase.QueryRow("INSERT INTO tasks(title, group_id) VALUES($1,$2) RETURNING task_id, title, group_id", task.Title, task.GroupID).Scan(&createdTask.ID, &createdTask.Title, &createdTask.GroupID)
	if err != nil {
		return createdTask, err
	}
	return createdTask, nil
}

// UpdateTask - обновляет таску в БД по ID полученными значениями и возвращает обновленную запись
func UpdateTask(task Task) (Task, error) {
	var updatedTask Task
	err := dbase.QueryRow("UPDATE tasks SET title=$1, group_id=$2 WHERE task_id=$3 RETURNING task_id, title, group_id;", task.Title, task.GroupID, task.ID).Scan(&updatedTask.ID, &updatedTask.Title, &updatedTask.GroupID)
	if err != nil {
		return updatedTask, err
	}
	return updatedTask, nil
}

// DeleteTask - удаляет из БД таску по полученному ID
func DeleteTask(id int) error {
	_, err := dbase.Exec("DELETE FROM tasks WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// CreateTimeframe - создает в БД новый таймфрейм с заданными значениями и возвращает эту запись
func CreateTimeframe(timeframe Timeframe) (Timeframe, error) {
	var createdTimeframe Timeframe
	err := dbase.QueryRow("INSERT INTO timeframes(task_id, time_start, time_end) VALUES ($1,$2, $3) RETURNING task_id, time_start, time_end", timeframe.TaskID, timeframe.TimeFrom, timeframe.TimeTo).Scan(&createdTimeframe.TaskID, &createdTimeframe.TimeFrom, &createdTimeframe.TimeTo)
	if err != nil {
		return createdTimeframe, err
	}
	return createdTimeframe, nil
}

// DeleteTimeframe - удаляет из БД таймфрейм по полученному ID
func DeleteTimeframe(id int) error {
	_, err := dbase.Exec("DELETE FROM timeframes WHERE task_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// getTasksByGroupID - получает из БД таски по ID и возвращает сформированный срез тасков
func getTasksByGroupID(id int) ([]Task, error) {
	var tasks []Task
	taskTable, err := dbase.Query("SELECT * FROM tasks WHERE group_id=$1 ORDER BY task_id", id)
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
		task.Timeframes, err = getTimeframesByTaskID(task.ID)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// getTimeframesByTaskID - получает из БД таймфреймы по ID и возвращает сформированный срез таймфреймов
func getTimeframesByTaskID(id int) ([]Timeframe, error) {
	var timeframes []Timeframe
	timeframesTable, err := dbase.Query("SELECT * FROM timeframes WHERE task_id=$1", id)
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
