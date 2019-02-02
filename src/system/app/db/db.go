package db

import (
	"database/sql"
	"errors"

	"../auth"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID      int    `json:"id"`
	Login   string `json:"login"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	RoleID  int    `json:"roleId"`
	Block   int8   `json:"block"`
	Address string `json:"address"`
	Img     string `json:"img"`
}

func UserCredential() UserCredentials {
	return UserCredentials{}
}

func Init() *sql.DB {
	Db, err := sql.Open("mysql", "u839560_office:Zxczxc888@tcp(185.185.69.91)/u839560_office?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
	return Db
}

func NewUser(login string, hash []byte, email string) (string, int, error) {
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("INSERT INTO users (login, password, email, role_id, block) VALUES (?, ?, ?, 3, 0)")
	if err != nil {
		return "", 0, errors.New("Ошибка при регистрации, попробуйте ещё раз")
	}
	defer insert.Close()
	lastId, err := insert.Exec(login, hash, email)
	if err != nil {
		return "", 0, errors.New("Ошибка при регистрации, попробуйте ещё раз")
	} else {
		id, _ := lastId.LastInsertId()
		return auth.Auth(int(id))
	}
}

func (user *UserCredentials) RegisterUser() (string, int, error) {
	var id int
	d := Init()
	defer d.Close()
	exists, err := d.Prepare("SELECT id FROM users WHERE (login = ? or email = ?)")
	if err != nil {
		return "", 0, errors.New("Ошибка при регистрации, попробуйте ещё раз")
	}
	defer exists.Close()
	_ = exists.QueryRow(user.Login, user.Email).Scan(&id)
	if id == 0 {
		cost := bcrypt.DefaultCost
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
		if err != nil {
			return "", 0, errors.New("Ошибка шифрования пароля")
		}
		return NewUser(user.Login, hash, user.Email)
	} else {
		return "", 0, errors.New("Пользователь с такими данными уже зарегистрирован")
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *UserCredentials) CheckUser() (string, int, error) {
	var id int
	var pass string
	d := Init()
	defer d.Close()
	exists, err := d.Prepare("SELECT id, password FROM users WHERE (login = ?)")
	if err != nil {
		return "", 0, errors.New("Ошибка проверки логина и пароля")
	}
	defer exists.Close()
	err = exists.QueryRow(user.Login).Scan(&id, &pass)
	if err == nil && CheckPasswordHash(user.Password, pass) {
		return auth.Auth(id)
	} else {
		return "", 0, errors.New("Ошибка проверки логина и пароля")
	}
}

func GetUserById(id int64, token string) (interface{}, error) {
	var user struct {
		Login string `json:"login"`
		Name  string `json:"name"`
		Token string `json:"token"`
	}
	user.Token = token
	d := Init()
	defer d.Close()
	exists, err := d.Prepare("SELECT login, name FROM users WHERE (id = ?)")
	if err != nil {
		return nil, err
	}
	defer exists.Close()
	err = exists.QueryRow(id).Scan(&user.Login, &user.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllWorkers(dt_start, dt_end string) (interface{}, error) {
	type UserSubTaskReturn struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Time     int    `json:"time"`
		TrueTime int    `json:"true_time"`
	}
	var UserSubTask UserSubTaskReturn
	var UserSubTaskArr []UserSubTaskReturn
	d := Init()
	defer d.Close()
	exists, err := d.Prepare("SELECT u.id, u.name, sum(sub.price) as sum_price, (sum(sub.time_dev) + sum(sub.time_manage)) as sum_time, sum(sub.true_time) as sum_true_time FROM users u LEFT JOIN sub_task sub ON u.id = sub.user_id WHERE sub.dt_create > ? AND sub.dt_create < ? AND sub.status_id = 4 GROUP BY u.id")
	if err != nil {
		return nil, err
	}
	defer exists.Close()
	rows, err := exists.Query(dt_start, dt_end)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&UserSubTask.ID, &UserSubTask.Name, &UserSubTask.Price, &UserSubTask.Time, &UserSubTask.TrueTime)
		if err != nil {
			return nil, err
		}
		UserSubTaskArr = append(UserSubTaskArr, UserSubTask)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return UserSubTaskArr, nil
}
