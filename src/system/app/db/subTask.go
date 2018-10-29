package db

import (
	"strconv"

	"../auth"
)

//SubTask model
type SubTask struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	TimeDev    int    `json:"time_dev"`
	TimeManage int    `json:"time_manage"`
	Message    string `json:"message"`
	TaskID     int    `json:"task_id"`
	StatusID   int    `json:"status_id"`
	UserID     int    `json:"user_id"`
}

//SubTaskReturn model for return to client
type SubTaskReturn struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	TimeDev    int    `json:"time_dev"`
	TimeManage int    `json:"time_manage"`
	Message    string `json:"message"`
	TaskID     int    `json:"task_id"`
	Status     string `json:"status"`
	User       string `json:"user"`
}

//NewSubTask - make new subtask. Return json {ok: var1, data: var2}
func (subTask *SubTask) NewSubTask() []byte {
	var str []byte
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("INSERT INTO sub_task VALUES (NULL, ?, ?, ?, ?, ?, ?, 1, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	lastId, err := insert.Exec(subTask.Name,
		subTask.Price,
		subTask.TimeDev,
		subTask.TimeManage,
		subTask.Message,
		subTask.TaskID,
		subTask.UserID)
	if err != nil {
		str = auth.JsonResponseByVar("false", "Ошибка при создании задачи"+err.Error())
	} else {
		id, _ := lastId.LastInsertId()
		str = auth.JsonResponseByVar("true", "Подзадача успешно создана под номером "+strconv.FormatInt(id, 10))
	}
	return str
}

func GetAllSubTasks() ([]SubTaskReturn, error) {
	var subTaskArr []SubTaskReturn
	var subTask SubTaskReturn
	d := Init()
	defer d.Close()
	rows, err := d.Query("SELECT s.id, s.name, s.price, s.time_dev, s.time_manage, s.message, s.task_id, st.name, u.name FROM sub_task s INNER JOIN status st ON s.status_id = st.id INNER JOIN users u ON s.user_id = u.id ORDER BY s.task_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&subTask.ID, &subTask.Name, &subTask.Price, &subTask.TimeDev, &subTask.TimeManage, &subTask.Message, &subTask.TaskID, &subTask.Status, &subTask.User)
		if err != nil {
			return nil, err
		}
		subTaskArr = append(subTaskArr, subTask)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return subTaskArr, nil
}

func GetSubTasksByTask(id string) ([]SubTaskReturn, error) {
	var subTaskArr []SubTaskReturn
	var subTask SubTaskReturn
	d := Init()
	defer d.Close()
	rows, err := d.Query("SELECT s.id, s.name, s.price, s.time_dev, s.time_manage, s.message, s.task_id, st.name, u.name FROM sub_task s INNER JOIN status st ON s.status_id = st.id INNER JOIN users u ON s.user_id = u.id WHERE (s.task_id = " + id + ") ORDER BY s.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&subTask.ID, &subTask.Name, &subTask.Price, &subTask.TimeDev, &subTask.TimeManage, &subTask.Message, &subTask.TaskID, &subTask.Status, &subTask.User)
		if err != nil {
			return nil, err
		}
		subTaskArr = append(subTaskArr, subTask)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return subTaskArr, nil
}
