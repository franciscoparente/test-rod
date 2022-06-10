package main

import (
	"fmt"
	"test-rod/service"
)

func main() {
	service := service.ServiceSt{}
	if err := service.ValidUser("asd"); err != nil {
		fmt.Println(err)
	}
}
