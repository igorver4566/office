package db

import (
	"database/sql"
	"fmt"

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
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	RoleID   int    `json:"roleId"`
	Block    int8   `json:"block"`
	Address  string `json:"address"`
	Img      string `json:"img"`
}

func UserCredential() UserCredentials {
	return UserCredentials{}
}

func Init() *sql.DB {
	Db, err := sql.Open("mysql", "u839560_office:Zxczxc888@tcp(91.217.9.197)/u839560_office")

	if err != nil {
		panic(err.Error())
	}
	return Db
}
func NewUser(login string, hash []byte, email string) []byte {
	var id int
	d := Init()
	defer d.Close()
	insert, err := d.Prepare("INSERT INTO users (login, password, email, role_id, block) VALUES (?, ?, ?, 1, 0)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	_ = insert.QueryRow(login, hash, email).Scan(&id)

	return auth.Auth(id)
}

func (user *UserCredentials) RegisterUser() []byte {
	var id int
	var strerr []byte
	d := Init()
	defer d.Close()
	exists, err := d.Prepare("SELECT id FROM users WHERE (login = ? or email = ?)")
	if err != nil {
		panic(err.Error())
	}
	defer exists.Close()
	_ = exists.QueryRow(user.Login, user.Email).Scan(&id)
	fmt.Println("id -----> ", id)
	if id == 0 {
		cost := bcrypt.DefaultCost
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
		if err != nil {
			panic(err.Error())
		}
		strerr = NewUser(user.Login, hash, user.Email)
	} else {
		var str auth.Token
		str.Ok = "false"
		str.Token = "Пользователь с такими данными уже зарегистрирован"
		strerr = auth.JsonResponse(str)
	}

	return strerr
}
