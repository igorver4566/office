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
	StatusID    int    `json:"status_id"`
}

//Task model for return to client
type TaskReturn struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Access     string `json:"access"`
	Manager    string `json:"manager"`
	TimeDev    int    `json:"time_dev"`
	TimeManage int    `json:"time_manage"`
	Owner      string `json:"owner"`
	Developer  string `json:"developer"`
	Price      int    `json:"price"`
	Tags       string `json:"tags"`
	Message    string `json:"message"`
	NameSlack  string `json:"name_slack"`
	Status     string `json:"status"`
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

func GetAllTasks() ([]TaskReturn, error) {
	var tasksArr []TaskReturn
	var task TaskReturn
	d := Init()
	defer d.Close()
	rows, err := d.Query("SELECT t.id, t.name, t.access, u.name, t.time_dev, t.time_manage, o.name, u2.name, t.price, t.tags, t.message, t.name_slack, s.name FROM task t INNER JOIN users u ON t.manager_id = u.id INNER JOIN owner o ON t.owner_id = o.id INNER JOIN users u2 ON t.developer_id = u2.id INNER JOIN status s ON t.status_id = s.id ORDER BY t.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Name, &task.Access, &task.Manager, &task.TimeDev, &task.TimeManage, &task.Owner, &task.Developer, &task.Price, &task.Tags, &task.Message, &task.NameSlack, &task.Status)
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
