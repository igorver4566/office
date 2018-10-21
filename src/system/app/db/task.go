package db

import (
	"strconv"

	"../auth"
)

//Task model
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Access      string `json:"access"`
	ManagerID   int    `json:"manager_id"`
	TimeDev     int    `json:"time_dev"`
	TimeManage  int    `json:"time_manage"`
	OwnerID     int    `json:"owner_id"`
	DeveloperID int    `json:"developer_id"`
	Price       int    `json:"price"`
	Tags        string `json:"tags"`
	MakeSlack   int8   `json:"make_slack"`
	Message     string `json:"message"`
	NameSlack   string `json:"name_slack"`
	StatusID    string `json:"status_id"`
}

//NewTask - make new task. Return json {ok: var1, data: var2}
func (task *Task) NewTask() []byte {
	var str []byte
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("INSERT INTO task VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	lastId, err := insert.Exec(task.Name,
		task.Access,
		task.ManagerID,
		task.TimeDev,
		task.TimeManage,
		task.OwnerID,
		task.DeveloperID,
		task.Price,
		task.Tags,
		task.MakeSlack,
		task.Message,
		task.NameSlack)
	if err != nil {
		str = auth.JsonResponseByVar("false", "Ошибка при создании задачи"+err.Error())
	} else {
		id, _ := lastId.LastInsertId()
		str = auth.JsonResponseByVar("true", "Задача успешно создана под номером "+strconv.FormatInt(id, 10))
	}
	return str
}

func GetAllTasks() ([]Task, error) {
	var tasksArr []Task
	var task Task
	d := Init()
	defer d.Close()
	rows, err := d.Query("SELECT * FROM task ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Name, &task.Access, &task.ManagerID, &task.TimeDev, &task.TimeManage, &task.OwnerID, &task.DeveloperID, &task.Price, &task.Tags, &task.MakeSlack, &task.Message, &task.NameSlack, &task.StatusID)
		if err != nil {
			return nil, err
		}
		tasksArr = append(tasksArr, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasksArr, nil
}
