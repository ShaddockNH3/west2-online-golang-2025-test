package main

import (
	user "github.com/ShaddockNH3/west2-online-golang-2025-test/task5/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
