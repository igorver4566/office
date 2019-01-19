package db

import (
	"fmt"
	"strconv"
	"time"

	"../auth"
)

//SubTask model
type SubTask struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	TimeDev    int       `json:"time_dev"`
	TimeManage int       `json:"time_manage"`
	Message    string    `json:"message"`
	TaskID     int       `json:"task_id"`
	StatusID   int       `json:"status_id"`
	UserID     int       `json:"user_id"`
	DTCreate   time.Time `json:"dt_create"`
	Priority   int       `json:"priority"`
	TrueTime   int       `json:"true_time"`
}

//SubTaskReturn model for return to client
type SubTaskReturn struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	TimeDev    int       `json:"time_dev"`
	TimeManage int       `json:"time_manage"`
	Message    string    `json:"message"`
	TaskID     int       `json:"task_id"`
	Status     string    `json:"status"`
	User       string    `json:"user"`
	DTCreate   time.Time `json:"dt_create"`
	Priority   int       `json:"priority"`
	TrueTime   int       `json:"true_time"`
}

//NewSubTask - make new subtask. Return json {ok: var1, data: var2}
func (subTask *SubTask) NewSubTask() []byte {
	var str []byte
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("INSERT INTO sub_task VALUES (NULL, ?, ?, ?, ?, ?, ?, 1, ?, NOW(), ?, NULL)")
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
		subTask.UserID,
		subTask.Priority)
	if err != nil {
		str = auth.JsonResponseByVar("false", "Ошибка при создании задачи"+err.Error())
	} else {
		id, _ := lastId.LastInsertId()
		str = auth.JsonResponseByVar("true", "Подзадача успешно создана под номером "+strconv.FormatInt(id, 10))
	}
	return str
}

func (subTask *SubTask) EditSubTasksByID() []byte {
	var str []byte
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("UPDATE sub_task SET name = ?, price = ?, time_dev = ?, time_manage = ?, message = ?, user_id = ?, priority = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	_, err = insert.Exec(subTask.Name,
		subTask.Price,
		subTask.TimeDev,
		subTask.TimeManage,
		subTask.Message,
		subTask.UserID,
		subTask.Priority,
		subTask.ID)
	if err != nil {
		str = auth.JsonResponseByVar("false", "Ошибка при редактировании задачи"+err.Error())
	} else {
		str = auth.JsonResponseByVar("true", "Подзадача успешно отредактирована")
	}
	return str
}

func (subTask *SubTask) EditStatusSubTaskByID() []byte {
	var str []byte
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("UPDATE sub_task SET status_id = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	_, err = insert.Exec(
		subTask.StatusID,
		subTask.ID)
	if err != nil {
		str = auth.JsonResponseByVar("false", "Ошибка при редактировании статуса задачи"+err.Error())
	} else {
		str = auth.JsonResponseByVar("true", "Статус успешно отредактирован")
	}
	return str
}

func (subTask *SubTask) EditTimeSubTaskByID() []byte {
	var str []byte
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("UPDATE sub_task SET true_time = (SELECT s.true_time FROM sub_task s WHERE s.id = ?)+? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(subTask.TrueTime)
	defer insert.Close()
	_, err = insert.Exec(
		subTask.ID,
		subTask.TrueTime,
		subTask.ID)
	if err != nil {
		str = auth.JsonResponseByVar("false", "Ошибка при редактировании времени задачи"+err.Error())
	} else {
		str = auth.JsonResponseByVar("true", "Время успешно добавлено")
	}
	return str
}

func GetAllSubTasks() ([]SubTaskReturn, error) {
	var subTaskArr []SubTaskReturn
	var subTask SubTaskReturn
	d := Init()
	defer d.Close()
	rows, err := d.Query("SELECT s.id, s.name, s.price, s.time_dev, s.time_manage, s.message, s.task_id, st.name, u.name, s.dt_create, s.priority FROM sub_task s INNER JOIN status st ON s.status_id = st.id INNER JOIN users u ON s.user_id = u.id ORDER BY s.task_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&subTask.ID, &subTask.Name, &subTask.Price, &subTask.TimeDev, &subTask.TimeManage, &subTask.Message, &subTask.TaskID, &subTask.Status, &subTask.User, &subTask.DTCreate, &subTask.Priority)
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
	rows, err := d.Query("SELECT s.id, s.name, s.price, s.time_dev, s.time_manage, s.message, s.task_id, st.name, u.name, s.dt_create, s.priority, s.true_time FROM sub_task s INNER JOIN status st ON s.status_id = st.id INNER JOIN users u ON s.user_id = u.id WHERE (s.task_id = " + id + ") ORDER BY s.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&subTask.ID, &subTask.Name, &subTask.Price, &subTask.TimeDev, &subTask.TimeManage, &subTask.Message, &subTask.TaskID, &subTask.Status, &subTask.User, &subTask.DTCreate, &subTask.Priority, &subTask.TrueTime)
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
