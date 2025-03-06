package main

import "github.com/Hadis2971/go_web/layers/application"


func main () {
	application := application.NewApplication(":3000");

	application.Run();
}