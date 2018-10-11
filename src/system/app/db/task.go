package db

import (
	"../auth"
)

//Task model
type Task struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Access     string `json:"access"`
	ManagerID  int    `json:"manager_id"`
	TimeDev    int    `json:"time_dev"`
	TimeManage int    `json:"time_manage"`
	OwnerID    int    `json:"owner_id"`
	Price      int    `json:"price"`
	Tags       string `json:"tags"`
	MakeSlack  int8   `json:"make_slack"`
	Message    string `json:"message"`
	NameSlack  string `json:"name_slack"`
	StatusID   string `json:"status_id"`
}

//NewTask - make new task. Return json {ok: var1, data: var2}
func (task *Task) NewTask() []byte {
	var str []byte
	var id int
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("INSERT INTO task VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	err = insert.QueryRow(task.Name,
		task.Access,
		task.ManagerID,
		task.TimeDev,
		task.TimeManage,
		task.OwnerID,
		task.Price,
		task.Tags,
		task.MakeSlack,
		task.Message,
		task.NameSlack,
		task.StatusID).Scan(&id)
	if err != nil {
		str = auth.JsonResponseByVar("false", "Ошибка при создании задачи")
		panic(err.Error())
	} else {
		str = auth.JsonResponseByVar("true", "Задача успешно создана под номером "+string(id))
	}
	return str
}
