package main

import (
	"github.com/codegangsta/negroni"
	"github.com/equoia/twitchtvaudio/routers"
	"github.com/equoia/twitchtvaudio/controllers"
	"os/exec"
	"log"
)


func main() {
	currentcommit, err := exec.Command("git", "rev-parse", "--short",  "HEAD").Output()
	if err != nil {
		log.Println(err)
	}
	controllers.Version = string(currentcommit)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(controllers.HttpHeaderSet))
	n.UseHandler(routers.Router)
	n.Run(os.Getenv("OPENSHIFT_GO_IP"):os.Getenv("OPENSHIFT_GO_PORT"))
}
