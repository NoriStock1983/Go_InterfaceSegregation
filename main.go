package main

import (
	controller "InterfaceSegregation/controller/usermst"
	"InterfaceSegregation/infrastructure/repositories"
	usecase "InterfaceSegregation/usecase/usermst"
	"fmt"
)

func main() {

	// DI(Repositoryコンストラクタの生成)
	repo := repositories.NewSearchUsermstRepository()

	// DI(Usecaseコンストラクタの生成⇨Repositoryを引数に渡す)
	usermstusecase := usecase.NewSearchUsermstUsecase(repo)

	// DI(Controllerコンストラクタの生成⇨Usecaseを引数に渡す)
	c := controller.NewUsermstController(usermstusecase)

	fmt.Println(c.SarchAllUsermst())
	fmt.Println(c.SearchByUsercd())
}
