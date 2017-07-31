package main

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id        int64
	Firstname string
	Lastname  string
	Password  string
}

func (u *User) TableName() string {
	return "users"
}

func main() {

	orm, err := xorm.NewEngine("sqlite3", "test.db")
	if err != nil {
		log.Println(err)
		return
	}
	orm.ShowSQL(true)

	users := make([]User, 0)
	err = orm.Find(&users)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(users)
}
