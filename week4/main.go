package main

import(
	"firstgin/router"
)

func main(){
	r := router.InitRouter()
	r.Run(":8081")
}
