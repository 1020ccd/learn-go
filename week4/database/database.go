package database

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


var DB *sql.DB


func init(){
	dsn := "root:123456@tcp(127.0.0.1:3306)/wordpress?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		fmt.Println("数据库打开异常")
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil{
		fmt.Println("数据库ping异常")
		fmt.Println(err)
	}
	DB = db
}
