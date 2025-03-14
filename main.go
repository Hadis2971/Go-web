package main

import (
	"github.com/Hadis2971/go_web/util"

	"github.com/Hadis2971/go_web/layers/application"
)


func main () {

	PORT := util.GetEnvVariable("PORT")

	application := application.NewApplication(PORT);

	application.Run();
}