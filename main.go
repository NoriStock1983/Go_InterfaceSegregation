package main

import (
	controller "InterfaceSegregation/controller/usermst"
	"InterfaceSegregation/infrastructure/repositories"
	usecase "InterfaceSegregation/usecase/usermst"
	"fmt"
)

func main() {

	repo := repositories.NewSearchUsermstRepository()
	usermstusecase := usecase.NewSearchUsermstUsecase(repo)

	c := controller.NewUsermstController(usermstusecase)

	fmt.Println(c.SarchAllUsermst())
	fmt.Println(c.SearchByUsercd())
}
