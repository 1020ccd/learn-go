package model

import(
	"firstgin/database"
	"fmt"
)



type User struct{
	Name string `json:"name"`

	Email string `json:"email"`
}

type TestUser struct{

}

func(t TestUser) FindUser() User{
	
	sql := "select user_login, user_email from wp_users where ID = ?"
	var u User
	var name string
	var email string
	err := database.DB.QueryRow(sql,1).Scan(&name,&email)
	if err != nil{
		fmt.Println(err)
	}
	u = User{
		Name:name,
		Email:email,
		}

	return u
}
