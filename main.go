package main

import (
	"github.com/Hadis2971/go_web/util"

	"github.com/Hadis2971/go_web/layers/application"
)

// DockerFile is good to add, and required for most backends these days. Makes it way easier to share projects.
// Hadis => Can I do it at the end? I guess I can. It's just to make an image with dependecies at the end right?

func main() {

	PORT := util.GetEnvVariable("PORT")

	GoWeb := application.NewApplication(PORT)

	GoWeb.Run()
}
