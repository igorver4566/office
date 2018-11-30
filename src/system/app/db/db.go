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
	Db, err := sql.Open("mysql", "u839560_office:Zxczxc888@localhost:3306/u839560_office")

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
