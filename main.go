package main

import (
	"github.com/Hadis2971/go_web/util"

	"github.com/Hadis2971/go_web/layers/application"
)

// DockerFile is good to add, and required for most backends these days. Makes it way easier to share projects.

func main() { // Use `gofmt w .` golang has a linter/formatter built in to the compiler

	PORT := util.GetEnvVariable("PORT")

	application := application.NewApplication(PORT) // application := application ? This would overwrite the library, use a different name

	application.Run()
}
