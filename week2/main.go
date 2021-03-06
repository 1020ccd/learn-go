package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func test() (string, error) {
	var email string
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/wordpress?charset=utf8")
	if err != nil {
		//	err = errors.WithStack(err)
		err = errors.Wrap(err, "message my test")
		//	fmt.Print(err)
		return "", err
	}
	rows, err := db.Query("SELECT user_email FROM wp_users where id = 9 ")
	if err != nil {
		err = errors.Wrap(err, "db message")
		//////////////// fmt.Print(err)
		return "", err
	}
	err = rows.Scan(&email)
	if err != nil {
		err = errors.Wrap(err, "未查询到")
		return "", err
	}
	return email, nil
}

func main() {

	ret, err := test()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(ret)

}
